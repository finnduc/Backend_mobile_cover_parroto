package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/user/dtos/req"
	"go-cover-parroto/internal/modules/user/dtos/res"
	"go-cover-parroto/internal/modules/user/repositories"
	"go-cover-parroto/internal/utils"
)

type IUserService interface {
	GetProfile(ctx context.Context) (*res.UserRes, *response.AppError)
	List(ctx context.Context, query req.ListUserQuery) (*response.PaginatedResponse[res.UserRes], *response.AppError)
	GetByID(ctx context.Context, id uint) (*res.UserRes, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateUserReq) (*res.UserRes, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
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

func (s *userService) List(ctx context.Context, query req.ListUserQuery) (*response.PaginatedResponse[res.UserRes], *response.AppError) {
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		return nil, response.Internal("failed to list users")
	}
	var users []res.UserRes
	if err := utils.MapToDTOs(result.Data, &users); err != nil {
		return nil, response.Internal("failed to map users")
	}
	return &response.PaginatedResponse[res.UserRes]{Data: users, Meta: result.Meta}, nil
}

func (s *userService) GetByID(ctx context.Context, id uint) (*res.UserRes, *response.AppError) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("user not found")
		}
		return nil, response.Internal("failed to get user")
	}
	var result res.UserRes
	if err := utils.MapToDTO(user, &result); err != nil {
		return nil, response.Internal("failed to map user")
	}
	return &result, nil
}

func (s *userService) Update(ctx context.Context, id uint, body req.UpdateUserReq) (*res.UserRes, *response.AppError) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("user not found")
		}
		return nil, response.Internal("failed to get user")
	}
	if body.Name != nil {
		user.Name = *body.Name
	}
	if body.AvatarURL != nil {
		user.AvatarURL = *body.AvatarURL
	}
	if updateErr := s.repo.Update(ctx, user); updateErr != nil {
		return nil, response.Internal("failed to update user")
	}
	var result res.UserRes
	if err := utils.MapToDTO(user, &result); err != nil {
		return nil, response.Internal("failed to map user")
	}
	return &result, nil
}

func (s *userService) Delete(ctx context.Context, id uint) *response.AppError {
	if err := s.repo.Delete(ctx, id); err != nil {
		return response.Internal("failed to delete user")
	}
	return nil
}
