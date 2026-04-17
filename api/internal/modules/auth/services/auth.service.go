package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	authres "go-cover-parroto/internal/modules/auth/dtos/res"
	"go-cover-parroto/internal/modules/auth/repositories"

	fb "go-cover-parroto/internal/firebase"
)

type IAuthService interface {
	SyncUser(ctx context.Context, firebaseToken string) (*authres.SyncRes, *response.AppError)
}

type authService struct {
	repo repositories.IAuthRepo
}

func NewAuthService(repo repositories.IAuthRepo) IAuthService {
	return &authService{repo: repo}
}

func (s *authService) SyncUser(ctx context.Context, firebaseToken string) (*authres.SyncRes, *response.AppError) {
	decoded, err := fb.AuthClient.VerifyIDToken(ctx, firebaseToken)
	if err != nil {
		return nil, response.Unauthorized("invalid firebase token")
	}

	email, _ := decoded.Claims["email"].(string)
	name, _ := decoded.Claims["name"].(string)
	picture, _ := decoded.Claims["picture"].(string)

	if email == "" {
		return nil, response.BadRequest("firebase token missing email")
	}

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, coreError.ErrNotFound) {
			return nil, response.Internal("failed to find user")
		}

		user = &models.User{
			Email:     email,
			Name:      name,
			AvatarURL: picture,
			Password:  "",
		}

		if createErr := s.repo.Create(ctx, user); createErr != nil {
			return nil, response.Internal("failed to create user")
		}
	}

	return &authres.SyncRes{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}, nil
}
