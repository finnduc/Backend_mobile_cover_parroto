package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/firebase"
	authres "go-cover-parroto/internal/modules/auth/dtos/res"
	"go-cover-parroto/internal/modules/auth/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "auth")
}

type IAuthService interface {
	SyncUser(ctx context.Context, firebaseToken string, reqName string) (*authres.SyncRes, *response.AppError)
}

type authService struct {
	repo   repositories.IAuthRepo
	fbAuth firebase.IFirebaseAuth
}

func NewAuthService(repo repositories.IAuthRepo, fbAuth firebase.IFirebaseAuth) IAuthService {
	return &authService{repo: repo, fbAuth: fbAuth}
}

func (s *authService) SyncUser(ctx context.Context, firebaseToken string, reqName string) (*authres.SyncRes, *response.AppError) {
	decoded, err := s.fbAuth.VerifyIDToken(ctx, firebaseToken)
	if err != nil {
		return nil, response.Unauthorized("invalid firebase token")
	}

	email, _ := decoded.Claims["email"].(string)
	name, _ := decoded.Claims["name"].(string)
	picture, _ := decoded.Claims["picture"].(string)

	if name == "" && reqName != "" {
		name = reqName
	}

	if email == "" {
		return nil, response.BadRequest("firebase token missing email")
	}

	log := sLog().With("email", email)

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, coreError.ErrNotFound) {
			log.Errorw("failed to find user by email", "error", err)
			return nil, response.Internal("failed to find user")
		}

		user = &models.User{
			Email:     email,
			Name:      name,
			AvatarURL: picture,
			Password:  "",
		}

		if createErr := s.repo.Create(ctx, user); createErr != nil {
			log.Errorw("failed to create user", "error", createErr)
			return nil, response.Internal("failed to create user")
		}
		log.Infow("new user created via sync", "userId", user.ID)
	}

	log.Infow("user sync completed", "userId", user.ID)
	return &authres.SyncRes{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}, nil
}
