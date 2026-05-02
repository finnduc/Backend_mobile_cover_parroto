package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/user/dtos/res"
	"go-cover-parroto/internal/modules/user/repositories"
	"go-cover-parroto/internal/utils"
)

type IUserService interface {
	GetProfile(ctx context.Context) (*res.UserRes, *response.AppError)
}

type userService struct {
	repo repositories.IUserRepo
}

func NewUserService(repo repositories.IUserRepo) IUserService {
	return &userService{repo: repo}
}

func (s *userService) GetProfile(ctx context.Context) (*res.UserRes, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, findErr := s.repo.FindByID(ctx, userID)
	if findErr != nil {
		if errors.Is(findErr, coreError.ErrNotFound) {
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
