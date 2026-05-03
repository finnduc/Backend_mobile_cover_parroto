package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/user/dtos/req"
	"go-cover-parroto/internal/modules/user/dtos/res"
	"go-cover-parroto/internal/modules/user/repositories"
	"go-cover-parroto/internal/utils"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "user")
}

type IUserService interface {
	GetProfile(ctx context.Context) (*res.UserRes, *response.AppError)
	List(ctx context.Context, query req.ListUserQuery) (*response.PaginatedResponse[res.UserRes], *response.AppError)
	GetByID(ctx context.Context, id string) (*res.UserRes, *response.AppError)
	Update(ctx context.Context, id string, body req.UpdateUserReq) (*res.UserRes, *response.AppError)
	Delete(ctx context.Context, id string) *response.AppError
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

	log := sLog().With("userId", userID)
	log.Infow("getting user profile")
	user, findErr := s.repo.FindByID(ctx, userID)
	if findErr != nil {
		if errors.Is(findErr, coreError.ErrNotFound) {
			return nil, response.NotFound("user not found")
		}
		log.Errorw("failed to get profile", "error", findErr)
		return nil, response.Internal("failed to get profile")
	}
	var result res.UserRes
	if err := utils.MapToDTO(user, &result); err != nil {
		return nil, response.Internal("failed to map user")
	}
	return &result, nil
}

func (s *userService) List(ctx context.Context, query req.ListUserQuery) (*response.PaginatedResponse[res.UserRes], *response.AppError) {
	log := sLog()
	log.Infow("listing users")
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list users", "error", err)
		return nil, response.Internal("failed to list users")
	}
	var users []res.UserRes
	if err := utils.MapToDTOs(result.Data, &users); err != nil {
		return nil, response.Internal("failed to map users")
	}
	return &response.PaginatedResponse[res.UserRes]{Data: users, Meta: result.Meta}, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*res.UserRes, *response.AppError) {
	log := sLog().With("userId", id)
	log.Infow("getting user by id")
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("user not found")
		}
		log.Errorw("failed to get user", "error", err)
		return nil, response.Internal("failed to get user")
	}
	var result res.UserRes
	if err := utils.MapToDTO(user, &result); err != nil {
		return nil, response.Internal("failed to map user")
	}
	return &result, nil
}

func (s *userService) Update(ctx context.Context, id string, body req.UpdateUserReq) (*res.UserRes, *response.AppError) {
	log := sLog().With("userId", id)
	log.Infow("updating user")
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("user not found")
		}
		log.Errorw("failed to get user for update", "error", err)
		return nil, response.Internal("failed to get user")
	}
	if body.Name != nil {
		user.Name = *body.Name
	}
	if body.AvatarURL != nil {
		user.AvatarURL = *body.AvatarURL
	}
	if updateErr := s.repo.Update(ctx, user); updateErr != nil {
		log.Errorw("failed to update user", "error", updateErr)
		return nil, response.Internal("failed to update user")
	}
	log.Infow("user updated")
	var result res.UserRes
	if err := utils.MapToDTO(user, &result); err != nil {
		return nil, response.Internal("failed to map user")
	}
	return &result, nil
}

func (s *userService) Delete(ctx context.Context, id string) *response.AppError {
	log := sLog().With("userId", id)
	log.Infow("deleting user")
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Errorw("failed to delete user", "error", err)
		return response.Internal("failed to delete user")
	}
	log.Infow("user deleted")
	return nil
}
