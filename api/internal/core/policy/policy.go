package policy

import (
	"context"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/utils"
)

func ActorFromContext(ctx context.Context) (*Actor, *response.AppError) {
	userID, err := utils.GetFromContext[string](ctx, enums.ContextKeyUserID)
	if err != nil {
		return nil, err
	}
	role, _ := utils.GetFromContext[enums.UserRole](ctx, enums.ContextKeyUserRole)
	return &Actor{UserID: userID, Role: role}, nil
}

func CanMutate(actor *Actor, resourceOwnerID string) *response.AppError {
	if actor == nil {
		return response.Unauthorized("login required")
	}
	if actor.Role == enums.UserRoleAdmin {
		return nil
	}
	if actor.UserID == resourceOwnerID {
		return nil
	}
	return response.Forbidden("access denied")
}

func CanRead(actor *Actor, resourceOwnerID string) *response.AppError {
	if actor == nil {
		return response.Unauthorized("login required")
	}
	if actor.Role == enums.UserRoleAdmin {
		return nil
	}
	if actor.UserID == resourceOwnerID {
		return nil
	}
	return response.Forbidden("access denied")
}
