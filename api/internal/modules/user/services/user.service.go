package services

import (
	"context"

	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/user/repositories"
)

type IUserService interface {
	GetProfile(ctx context.Context) (*models.User, *response.AppError)
}

type userService struct {
	repo repositories.IUserRepo
}

func NewUserService(repo repositories.IUserRepo) IUserService {
	return &userService{repo: repo}
}

func (s *userService) GetProfile(ctx context.Context) (*models.User, *response.AppError) {
	userID, appErr := policy.GetUserID(ctx)
	if appErr != nil {
		return nil, appErr
	}

	user, findErr := s.repo.FindByID(ctx, userID)
	if findErr != nil {
		return nil, response.NotFound("user not found")
	}

	return user, nil
}
