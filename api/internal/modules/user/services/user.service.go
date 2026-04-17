package services

import (
	"context"
	"errors"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/user/dtos/res"
	"go-cover-parroto/internal/modules/user/repositories"
)

type IUserService interface {
	GetProfile(ctx context.Context, userID uint) (*res.UserRes, *response.AppError)
}

type userService struct {
	repo repositories.IUserRepo
}

func NewUserService(repo repositories.IUserRepo) IUserService {
	return &userService{repo: repo}
}

func (s *userService) GetProfile(ctx context.Context, userID uint) (*res.UserRes, *response.AppError) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, response.NotFound("user not found")
		}
		return nil, response.Internal("failed to get profile")
	}
	return &res.UserRes{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}, nil
}
