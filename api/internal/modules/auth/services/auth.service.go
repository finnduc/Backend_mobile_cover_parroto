package services

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/auth/dtos/res"

	clerkusersdk "github.com/clerk/clerk-sdk-go/v2/user"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "auth")
}

type IAuthService interface {
	CompleteSignUp(ctx context.Context) *response.AppError
	SyncUser(ctx context.Context) (*res.AuthUserRes, *response.AppError)
	GetUserProfile(ctx context.Context) (*res.AuthUserRes, *response.AppError)
}

type authService struct{}

func NewAuthService() IAuthService {
	return &authService{}
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
	metaJSON, err := json.Marshal(roleMeta)
	if err != nil {
		log.Errorw("failed to marshal role metadata", "error", err)
		return response.Internal("failed to process role metadata")
	}
	meta := json.RawMessage(metaJSON)

	_, err = clerkusersdk.Update(ctx, userID, &clerkusersdk.UpdateParams{
		PublicMetadata: &meta,
	})
	if err != nil {
		log.Error("Failed to sync role to Clerk", zap.Error(err))
		return response.Internal("failed to sync role to Clerk")
	}

	return nil
}

func (s *authService) getUserFromClerk(ctx context.Context) (*res.AuthUserRes, *response.AppError) {
	userID, appErr := policy.GetUserID(ctx)
	if appErr != nil {
		return nil, appErr
	}

	clerkUser, err := clerkusersdk.Get(ctx, userID)
	if err != nil {
		return nil, response.Unauthorized("invalid user")
	}

	email := ""
	if len(clerkUser.EmailAddresses) > 0 {
		email = clerkUser.EmailAddresses[0].EmailAddress
	}

	firstName := ""
	lastName := ""
	if clerkUser.FirstName != nil {
		firstName = *clerkUser.FirstName
	}
	if clerkUser.LastName != nil {
		lastName = *clerkUser.LastName
	}
	name := strings.TrimSpace(firstName + " " + lastName)
	if name == "" {
		name = "User"
	}

	avatarURL := ""
	if clerkUser.HasImage && clerkUser.ImageURL != nil {
		avatarURL = *clerkUser.ImageURL
	}

	role := enums.UserRoleUser
	if customClaims, ok := ctx.Value(enums.ContextKeyUserRole).(enums.UserRole); ok && customClaims != "" {
		role = customClaims
	}

	return &res.AuthUserRes{
		ID:        userID,
		Email:     email,
		Name:      name,
		UserRole:  role,
		AvatarURL: avatarURL,
		CreatedAt: time.UnixMilli(clerkUser.CreatedAt).Format(time.RFC3339),
	}, nil
}

func (s *authService) SyncUser(ctx context.Context) (*res.AuthUserRes, *response.AppError) {
	log := sLog()
	log.Infow("syncing user")
	return s.getUserFromClerk(ctx)
}

func (s *authService) GetUserProfile(ctx context.Context) (*res.AuthUserRes, *response.AppError) {
	log := sLog()
	log.Infow("getting user profile")
	return s.getUserFromClerk(ctx)
}
