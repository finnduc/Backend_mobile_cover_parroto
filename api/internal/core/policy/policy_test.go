package policy

import (
	"context"
	"testing"

	"go-cover-parroto/internal/core/enums"
)

func TestAllow_AdminPass(t *testing.T) {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, "1")
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, enums.UserRoleAdmin)

	err := Allow(ctx, "999")
	if err != nil {
		t.Errorf("expected admin to pass, got %v", err)
	}
}

func TestAllow_OwnerPass(t *testing.T) {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, "42")
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, enums.UserRoleUser)

	err := Allow(ctx, "42")
	if err != nil {
		t.Errorf("expected owner to pass, got %v", err)
	}
}

func TestAllow_NotOwnerForbidden(t *testing.T) {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, "1")
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, enums.UserRoleUser)

	err := Allow(ctx, "2")
	if err == nil {
		t.Error("expected forbidden error")
	} else if err.Code != 403 {
		t.Errorf("expected 403, got %d", err.Code)
	}
}

func TestAllow_NoUserID(t *testing.T) {
	ctx := context.Background()

	err := Allow(ctx, "1")
	if err == nil || err.Code != 401 {
		t.Errorf("expected 401 unauthorized, got %v", err)
	}
}

func TestGetUserID_Success(t *testing.T) {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, "7")

	uid, err := GetUserID(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if uid != "7" {
		t.Errorf("expected 7, got %s", uid)
	}
}

func TestGetUserID_Missing(t *testing.T) {
	ctx := context.Background()

	_, err := GetUserID(ctx)
	if err == nil || err.Code != 401 {
		t.Errorf("expected 401 unauthorized, got %v", err)
	}
}
