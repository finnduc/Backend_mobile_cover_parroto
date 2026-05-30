package services

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go-cover-parroto/internal/modules/chat/dtos/res"
	"go-cover-parroto/internal/modules/chat/hub"
	userrepo "go-cover-parroto/internal/modules/user/repositories"

	clerkuser "github.com/clerk/clerk-sdk-go/v2/user"
	"go.uber.org/zap"
)

const (
	defaultHistoryLimit = 20
	maxHistoryLimit     = 50
	maxContentLength    = 1000
	sendCooldown        = time.Second
)

var ErrRateLimited = errors.New("you are sending messages too fast")
var ErrEmptyContent = errors.New("message content cannot be empty")
var ErrContentTooLong = errors.New("message content exceeds maximum length")

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "chat")
}

type IChatService interface {
	GetHistory(ctx context.Context, beforeID uint64, limit int) (*res.HistoryRes, *response.AppError)
	SendMessage(ctx context.Context, userID, content string) error
	SyncUser(ctx context.Context, userID string)
}

type chatService struct {
	repo     db_repos.IChatRepo
	userRepo userrepo.IUserRepo
	hub      *hub.Hub

	mu       sync.Mutex
	lastSend map[string]time.Time

	syncedMu sync.Mutex
	synced   map[string]time.Time
}

func NewChatService(repo db_repos.IChatRepo, userRepo userrepo.IUserRepo, h *hub.Hub) IChatService {
	return &chatService{
		repo:     repo,
		userRepo: userRepo,
		hub:      h,
		lastSend: make(map[string]time.Time),
		synced:   make(map[string]time.Time),
	}
}

// SyncUser ensures the local users table has up-to-date name/avatar for the
// given Clerk user id. Called when a client connects to chat so message
// metadata renders correctly. Cached for 1h to avoid spamming Clerk.
func (s *chatService) SyncUser(ctx context.Context, userID string) {
	const ttl = time.Hour

	s.syncedMu.Lock()
	if last, ok := s.synced[userID]; ok && time.Since(last) < ttl {
		s.syncedMu.Unlock()
		return
	}
	s.syncedMu.Unlock()

	u, err := clerkuser.Get(ctx, userID)
	if err != nil {
		sLog().Warnw("failed to fetch user from clerk", "error", err, "userId", userID)
		return
	}

	name := ""
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

	email := ""
	for _, e := range u.EmailAddresses {
		if u.PrimaryEmailAddressID != nil && e.ID == *u.PrimaryEmailAddressID {
			email = e.EmailAddress
			break
		}
	}
	if email == "" && len(u.EmailAddresses) > 0 {
		email = u.EmailAddresses[0].EmailAddress
	}
	if name == "" {
		name = email
	}

	avatar := ""
	if u.ImageURL != nil {
		avatar = *u.ImageURL
	}

	row := &models.User{
		ID:        userID,
		Email:     email,
		Name:     name,
		AvatarURL: avatar,
		UserRole:  enums.UserRoleUser,
	}
	if err := s.userRepo.Upsert(ctx, row); err != nil {
		sLog().Errorw("failed to upsert user", "error", err, "userId", userID)
		return
	}

	s.syncedMu.Lock()
	s.synced[userID] = time.Now()
	s.syncedMu.Unlock()
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
		messages = append(messages, toMessageRes(m))
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

func (s *chatService) SendMessage(ctx context.Context, userID, content string) error {
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
		return errors.New("failed to save message")
	}

	// reload with user so broadcast contains user info
	full, err := s.repo.FindHistory(ctx, msg.ID+1, 1)
	var enriched *models.GlobalChatMessage
	if err == nil && len(full) > 0 && full[0].ID == msg.ID {
		enriched = full[0]
	} else {
		enriched = msg
	}

	payload := res.WSEvent{
		Type: "message",
		Data: toMessageRes(enriched),
	}
	raw, mErr := json.Marshal(payload)
	if mErr != nil {
		sLog().Errorw("failed to marshal broadcast payload", "error", mErr)
		return nil
	}
	s.hub.Broadcast(raw)
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

func toMessageRes(m *models.GlobalChatMessage) res.MessageRes {
	out := res.MessageRes{
		ID:        m.ID,
		UserID:    m.UserID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}
	if m.User != nil {
		out.UserName = m.User.Name
		out.AvatarURL = m.User.AvatarURL
	}
	return out
}
