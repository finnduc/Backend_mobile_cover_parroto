package policy

import (
	"context"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
)

type contextKey string

const (
	ContextKeyUserID   contextKey = "userID"
	ContextKeyUserRole contextKey = "userRole"
)

func Allow(ctx context.Context, resourceOwnerID uint) *response.AppError {
	userID, ok := ctx.Value(ContextKeyUserID).(uint)
	if !ok {
		return response.Unauthorized("user not authenticated")
	}

	role, ok := ctx.Value(ContextKeyUserRole).(enums.UserRole)
	if ok && role == enums.UserRoleAdmin {
		return nil
	}

	if userID == resourceOwnerID {
		return nil
	}

	return response.Forbidden("access denied")
}

func GetUserID(ctx context.Context) (uint, *response.AppError) {
	userID, ok := ctx.Value(ContextKeyUserID).(uint)
	if !ok {
		return 0, response.Unauthorized("user not authenticated")
	}
	return userID, nil
}
