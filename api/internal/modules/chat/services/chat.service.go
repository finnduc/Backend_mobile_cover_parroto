package services

import (
	"context"
	"strings"
	"sync"
	"time"

	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go-cover-parroto/internal/modules/chat/dtos/res"
	"go-cover-parroto/internal/modules/chat/hub"

	clerkuser "github.com/clerk/clerk-sdk-go/v2/user"
	"go.uber.org/zap"
)

const (
	defaultHistoryLimit = 20
	maxHistoryLimit     = 50
	maxContentLength    = 1000
	sendCooldown        = time.Second
)

var ErrRateLimited = response.TooManyRequests("you are sending messages too fast")
var ErrEmptyContent = response.BadRequest("message content cannot be empty")
var ErrContentTooLong = response.BadRequest("message content exceeds maximum length")

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "chat")
}

type IChatService interface {
	GetHistory(ctx context.Context, beforeID uint64, limit int) (*res.HistoryRes, *response.AppError)
	SendMessage(ctx context.Context, userID, content string) *response.AppError
}

type userCacheEntry struct {
	Name      string
	AvatarURL string
	ExpiresAt time.Time
}

type chatService struct {
	repo   db_repos.IChatRepo
	sseHub *hub.SSEHub

	mu       sync.Mutex
	lastSend map[string]time.Time

	userCache   map[string]*userCacheEntry
	userCacheMu sync.RWMutex
}

func NewChatService(repo db_repos.IChatRepo, sse *hub.SSEHub) IChatService {
	return &chatService{
		repo:      repo,
		sseHub:    sse,
		lastSend:  make(map[string]time.Time),
		userCache: make(map[string]*userCacheEntry),
	}
}

func (s *chatService) getCachedUser(ctx context.Context, userID string) (name, avatarURL string, ok bool) {
	s.userCacheMu.RLock()
	entry, found := s.userCache[userID]
	s.userCacheMu.RUnlock()

	if found && time.Now().Before(entry.ExpiresAt) {
		return entry.Name, entry.AvatarURL, true
	}

	u, err := clerkuser.Get(ctx, userID)
	if err != nil {
		sLog().Warnw("failed to fetch user from clerk", "error", err, "userId", userID)
		return "", "", false
	}

	if u.FirstName != nil {
		name = *u.FirstName
	}
	if u.LastName != nil {
		if name != "" {
			name += " "
		}
		name += *u.LastName
	}
	if name == "" && u.Username != nil {
		name = *u.Username
	}
	if name == "" {
		name = "User"
	}

	if u.ImageURL != nil {
		avatarURL = *u.ImageURL
	}

	s.userCacheMu.Lock()
	s.userCache[userID] = &userCacheEntry{
		Name:      name,
		AvatarURL: avatarURL,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	s.userCacheMu.Unlock()

	return name, avatarURL, true
}

func (s *chatService) toMessageRes(ctx context.Context, m *models.GlobalChatMessage) res.MessageRes {
	out := res.MessageRes{
		ID:        m.ID,
		UserID:    m.UserID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}
	if name, avatar, ok := s.getCachedUser(ctx, m.UserID); ok {
		out.UserName = name
		out.AvatarURL = avatar
	}
	return out
}

func (s *chatService) GetHistory(ctx context.Context, beforeID uint64, limit int) (*res.HistoryRes, *response.AppError) {
	if limit <= 0 {
		limit = defaultHistoryLimit
	}
	if limit > maxHistoryLimit {
		limit = maxHistoryLimit
	}

	// fetch one extra to know if there is more
	rows, err := s.repo.FindHistory(ctx, beforeID, limit+1)
	if err != nil {
		sLog().Errorw("failed to fetch history", "error", err)
		return nil, response.Internal("failed to fetch chat history")
	}

	hasMore := len(rows) > limit
	if hasMore {
		rows = rows[:limit]
	}

	messages := make([]res.MessageRes, 0, len(rows))
	for _, m := range rows {
		messages = append(messages, s.toMessageRes(ctx, m))
	}

	var nextID *uint64
	if hasMore && len(rows) > 0 {
		oldest := rows[len(rows)-1].ID
		nextID = &oldest
	}

	return &res.HistoryRes{
		Messages: messages,
		HasMore:  hasMore,
		NextID:   nextID,
	}, nil
}

func (s *chatService) SendMessage(ctx context.Context, userID, content string) *response.AppError {
	content = strings.TrimSpace(content)
	if content == "" {
		return ErrEmptyContent
	}
	if len(content) > maxContentLength {
		return ErrContentTooLong
	}

	if !s.allowSend(userID) {
		return ErrRateLimited
	}

	msg := &models.GlobalChatMessage{
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	if err := s.repo.Create(ctx, msg); err != nil {
		sLog().Errorw("failed to save message", "error", err, "userId", userID)
		return response.Internal("failed to save message")
	}

	s.sseHub.Publish(hub.MessagesStream, "chat.message.created", s.toMessageRes(ctx, msg))
	return nil
}

func (s *chatService) allowSend(userID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	if last, ok := s.lastSend[userID]; ok && now.Sub(last) < sendCooldown {
		return false
	}
	s.lastSend[userID] = now
	return true
}
