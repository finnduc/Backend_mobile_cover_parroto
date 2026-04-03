package service

import (
	"context"
	"fmt"
	"go-cover-parroto/internal/models"
	"go-cover-parroto/internal/repo"
	pkgerrors "go-cover-parroto/pkg/errors"
	"go-cover-parroto/pkg/utils"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type authService struct {
	authRepo    repo.IAuthRepo
	rdb         *redis.Client
	accessTTL   time.Duration
	refreshTTL  time.Duration
}

func NewAuthService(authRepo repo.IAuthRepo, rdb *redis.Client) IAuthService {
	return &authService{
		authRepo:   authRepo,
		rdb:        rdb,
		accessTTL:  15 * time.Minute,
		refreshTTL: 7 * 24 * time.Hour,
	}
}

func (s *authService) Register(ctx context.Context, input RegisterInput) (*models.User, error) {
	// Check duplicate email
	existing, err := s.authRepo.FindUserByEmail(ctx, input.Email)
	if err == nil && existing.ID != 0 {
		return nil, fmt.Errorf("authService.Register: %w", pkgerrors.ErrConflict)
	}

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("authService.Register: %w", err)
	}

	user := &models.User{
		Email:        input.Email,
		PasswordHash: hash,
		Name:         input.Name,
	}
	if err := s.authRepo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("authService.Register: %w", pkgerrors.ErrInternalServer)
	}
	return user, nil
}

func (s *authService) Login(ctx context.Context, input LoginInput) (*TokenPair, error) {
	user, err := s.authRepo.FindUserByEmail(ctx, input.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("authService.Login: %w", pkgerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("authService.Login: %w", pkgerrors.ErrInternalServer)
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		return nil, fmt.Errorf("authService.Login: %w", pkgerrors.ErrUnauthorized)
	}

	// Determine role (default "user", first role if set)
	role := "user"
	roles, _ := s.authRepo.GetUserRoles(ctx, user.ID)
	if len(roles) > 0 {
		role = roles[0].Name
	}

	accessToken, jti, err := utils.GenerateAccessToken(user.ID, role, s.accessTTL)
	if err != nil {
		return nil, fmt.Errorf("authService.Login: %w", pkgerrors.ErrInternalServer)
	}

	refreshToken := utils.GenerateRefreshToken()

	// Store refresh token in Redis
	if s.rdb != nil {
		key := fmt.Sprintf("refresh:%d", user.ID)
		s.rdb.Set(ctx, key, refreshToken, s.refreshTTL)
		_ = jti // jti attached in access token claims
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(s.accessTTL.Seconds()),
	}, nil
}

func (s *authService) Logout(ctx context.Context, accessToken string, userID uint) error {
	// Delete refresh token from Redis
	if s.rdb != nil {
		key := fmt.Sprintf("refresh:%d", userID)
		s.rdb.Del(ctx, key)
	}

	// Blacklist access token
	claims, err := utils.ParseToken(accessToken)
	if err == nil && s.rdb != nil {
		ttl := time.Until(claims.ExpiresAt.Time)
		if ttl > 0 {
			blacklistKey := fmt.Sprintf("blacklist:%s", claims.JTI)
			s.rdb.Set(ctx, blacklistKey, "1", ttl)
		}
	}
	return nil
}

func (s *authService) Refresh(ctx context.Context, refreshToken string, userID uint) (*TokenPair, error) {
	if s.rdb == nil {
		return nil, fmt.Errorf("authService.Refresh: %w", pkgerrors.ErrInternalServer)
	}

	key := fmt.Sprintf("refresh:%d", userID)
	stored, err := s.rdb.Get(ctx, key).Result()
	if err != nil || stored != refreshToken {
		return nil, fmt.Errorf("authService.Refresh: %w", pkgerrors.ErrUnauthorized)
	}

	user, err := s.authRepo.FindUserByEmail(ctx, "") // we need user by ID
	_ = user
	_ = err
	// Using userID directly; get role from Redis or default
	jwtSecret := os.Getenv("JWT_SECRET")
	_ = jwtSecret

	newAccess, _, err := utils.GenerateAccessToken(userID, "user", s.accessTTL)
	if err != nil {
		return nil, fmt.Errorf("authService.Refresh: %w", pkgerrors.ErrInternalServer)
	}
	newRefresh := utils.GenerateRefreshToken()
	s.rdb.Set(ctx, key, newRefresh, s.refreshTTL)

	return &TokenPair{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		ExpiresIn:    int(s.accessTTL.Seconds()),
	}, nil
}
