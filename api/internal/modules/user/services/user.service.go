package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/user/dtos/res"
	"go-cover-parroto/internal/modules/user/repositories"
	"go-cover-parroto/internal/utils"
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
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("user not found")
		}
		return nil, response.Internal("failed to get profile")
	}
	var result res.UserRes
	if err := utils.MapToDTO(user, &result); err != nil {
		return nil, response.Internal("failed to map user")
	}
	return &result, nil
}
