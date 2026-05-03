package services

import (
	"context"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/firebase"
	"go-cover-parroto/internal/modules/auth/repositories"

	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "auth")
}

type IAuthService interface {
	CompleteSignUp(ctx context.Context) *response.AppError
}

type authService struct {
	repo   repositories.IAuthRepo
	fbAuth firebase.IFirebaseAuth
}

func NewAuthService(repo repositories.IAuthRepo, fbAuth firebase.IFirebaseAuth) IAuthService {
	return &authService{repo: repo, fbAuth: fbAuth}
}

func (s *authService) CompleteSignUp(ctx context.Context) *response.AppError {
	log := sLog()

	userID, appErr := policy.GetUserID(ctx)
	if appErr != nil {
		log.Errorw("failed to get user id from context", "error", appErr)
		return appErr
	}

	log.With("userId", userID)

	err := s.fbAuth.SetCustomUserClaims(ctx, userID, map[string]interface{}{"role": enums.UserRoleUser})
	if err != nil {
		log.Errorw("failed to set custom user claims", "error", err)
		return response.Internal("failed to set custom user claims")
	}

	return nil
}
