package services

import (
	"context"
	"encoding/json"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	db_repos "go-cover-parroto/internal/database/repositories"

	clerkUser "github.com/clerk/clerk-sdk-go/v2/user"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "auth")
}

type IAuthService interface {
	CompleteSignUp(ctx context.Context) *response.AppError
}

type authService struct {
	repo db_repos.IAuthRepo
}

func NewAuthService(repo db_repos.IAuthRepo) IAuthService {
	return &authService{repo: repo}
}

func (s *authService) CompleteSignUp(ctx context.Context) *response.AppError {
	log := sLog()

	userID, appErr := policy.GetUserID(ctx)
	if appErr != nil {
		log.Errorw("failed to get user id from context", "error", appErr)
		return appErr
	}

	log.With("userId", userID)
	roleMeta := map[string]interface{}{
		"role": string(enums.UserRoleUser),
	}
	metaJSON, _ := json.Marshal(roleMeta)
	meta := json.RawMessage(metaJSON)

	_, err := clerkUser.Update(ctx, userID, &clerkUser.UpdateParams{
		PublicMetadata: &meta,
	})

	if err != nil {
		log.Error("Failed to sync role to Clerk", zap.Error(err))
		return response.Internal("failed to sync role to Clerk")
	}

	return nil
}
