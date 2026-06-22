package services

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	authreq "go-cover-parroto/internal/modules/auth/dtos/req"
	"go-cover-parroto/internal/modules/auth/dtos/res"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkusersdk "github.com/clerk/clerk-sdk-go/v2/user"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "auth")
}

type IAuthService interface {
	CompleteSignUp(ctx context.Context, userID string) *response.AppError
	SyncUser(ctx context.Context, userID string) (*res.AuthUserRes, *response.AppError)
	GetUserProfile(ctx context.Context, userID string) (*res.AuthUserRes, *response.AppError)
	UpdateUserProfile(ctx context.Context, userID string, body authreq.UpdateProfileReq) (*res.AuthUserRes, *response.AppError)
}

type authService struct{}

func NewAuthService() IAuthService {
	return &authService{}
}

func (s *authService) CompleteSignUp(ctx context.Context, userID string) *response.AppError {
	log := sLog().With("userId", userID)

	roleMeta := map[string]interface{}{}
	if current, err := clerkusersdk.Get(ctx, userID); err == nil && len(current.PublicMetadata) > 0 {
		_ = json.Unmarshal(current.PublicMetadata, &roleMeta)
	}
	roleMeta["role"] = string(enums.UserRoleUser)
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

func (s *authService) getUserFromClerk(ctx context.Context, userID string) (*res.AuthUserRes, *response.AppError) {
	clerkUser, err := clerkusersdk.Get(ctx, userID)
	if err != nil {
		return nil, response.Unauthorized("invalid user")
	}

	return s.mapClerkUser(ctx, clerkUser), nil
}

func (s *authService) mapClerkUser(ctx context.Context, clerkUser *clerk.User) *res.AuthUserRes {
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

	metadata := map[string]interface{}{}
	if len(clerkUser.PublicMetadata) > 0 {
		_ = json.Unmarshal(clerkUser.PublicMetadata, &metadata)
	}

	role := enums.UserRoleUser
	if customClaims, ok := ctx.Value(enums.ContextKeyUserRole).(enums.UserRole); ok && customClaims != "" {
		role = customClaims
	} else if roleVal, ok := metadata["role"].(string); ok && roleVal != "" {
		role = enums.UserRole(roleVal)
	}

	phone := ""
	if phoneVal, ok := metadata["phone"].(string); ok {
		phone = phoneVal
	}

	return &res.AuthUserRes{
		ID:        clerkUser.ID,
		Email:     email,
		Name:      name,
		UserRole:  role,
		AvatarURL: avatarURL,
		Phone:     phone,
		CreatedAt: time.UnixMilli(clerkUser.CreatedAt).Format(time.RFC3339),
	}
}

func (s *authService) SyncUser(ctx context.Context, userID string) (*res.AuthUserRes, *response.AppError) {
	log := sLog()
	log.Infow("syncing user")
	return s.getUserFromClerk(ctx, userID)
}

func (s *authService) GetUserProfile(ctx context.Context, userID string) (*res.AuthUserRes, *response.AppError) {
	log := sLog()
	log.Infow("getting user profile")
	return s.getUserFromClerk(ctx, userID)
}

func (s *authService) UpdateUserProfile(ctx context.Context, userID string, body authreq.UpdateProfileReq) (*res.AuthUserRes, *response.AppError) {
	log := sLog().With("userId", userID)
	log.Infow("updating user profile")

	current, err := clerkusersdk.Get(ctx, userID)
	if err != nil {
		return nil, response.Unauthorized("invalid user")
	}

	metadata := map[string]interface{}{}
	if len(current.PublicMetadata) > 0 {
		_ = json.Unmarshal(current.PublicMetadata, &metadata)
	}
	if _, ok := metadata["role"]; !ok {
		metadata["role"] = string(enums.UserRoleUser)
	}
	metadata["phone"] = body.Phone
	metaJSON, err := json.Marshal(metadata)
	if err != nil {
		return nil, response.Internal("failed to process profile metadata")
	}
	meta := json.RawMessage(metaJSON)

	firstName, lastName := splitName(body.Name)
	updated, err := clerkusersdk.Update(ctx, userID, &clerkusersdk.UpdateParams{
		FirstName:      &firstName,
		LastName:       &lastName,
		PublicMetadata: &meta,
	})
	if err != nil {
		log.Error("failed to update Clerk user", zap.Error(err))
		return nil, response.Internal("failed to update profile")
	}

	role, _ := metadata["role"].(string)
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, enums.UserRole(role))
	return s.mapClerkUser(ctx, updated), nil
}

func splitName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], strings.Join(parts[1:], " ")
}
