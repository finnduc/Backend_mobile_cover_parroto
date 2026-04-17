package services

import (
	"context"
	"errors"
	"time"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	authreq "go-cover-parroto/internal/modules/auth/dtos/req"
	authres "go-cover-parroto/internal/modules/auth/dtos/res"
	"go-cover-parroto/internal/modules/auth/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(ctx context.Context, body authreq.RegisterReq) (*authres.RegisterRes, *response.AppError)
	Login(ctx context.Context, body authreq.LoginReq) (*authres.LoginRes, *response.AppError)
	RefreshToken(ctx context.Context, body authreq.RefreshReq) (*authres.RefreshRes, *response.AppError)
}

type authService struct {
	repo repositories.IAuthRepo
}

func NewAuthService(repo repositories.IAuthRepo) IAuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, body authreq.RegisterReq) (*authres.RegisterRes, *response.AppError) {
	existing, err := s.repo.FindByEmail(ctx, body.Email)
	if err != nil && !errors.Is(err, coreError.ErrNotFound) {
		return nil, response.Internal("failed to check existing user")
	}
	if existing != nil {
		return nil, response.Conflict("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, response.Internal("failed to hash password")
	}

	user := &models.User{
		Email:    body.Email,
		Password: string(hashedPassword),
		Name:     body.Name,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, response.Internal("failed to create user")
	}

	return &authres.RegisterRes{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s *authService) Login(ctx context.Context, body authreq.LoginReq) (*authres.LoginRes, *response.AppError) {
	user, err := s.repo.FindByEmail(ctx, body.Email)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.Unauthorized("invalid credentials")
		}
		return nil, response.Internal("failed to find user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, response.Unauthorized("invalid credentials")
	}

	accessToken, _ := GenerateToken(user.ID, time.Hour)
	refreshToken, _ := GenerateToken(user.ID, time.Hour*24*7)

	return &authres.LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: authres.UserInfo{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, body authreq.RefreshReq) (*authres.RefreshRes, *response.AppError) {
	claims, err := ValidateToken(body.RefreshToken)
	if err != nil {
		return nil, response.Unauthorized("invalid refresh token")
	}

	accessToken, _ := GenerateToken(claims.UserID, time.Hour)
	return &authres.RefreshRes{AccessToken: accessToken}, nil
}
