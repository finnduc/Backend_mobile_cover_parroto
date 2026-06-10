package policy

import (
	"testing"

	"go-cover-parroto/internal/core/enums"
)

func TestCanMutate_AdminPasses(t *testing.T) {
	actor := &Actor{UserID: "1", Role: enums.UserRoleAdmin}
	err := CanMutate(actor, "999")
	if err != nil {
		t.Errorf("expected admin to pass, got %v", err)
	}
}

func TestCanMutate_OwnerPasses(t *testing.T) {
	actor := &Actor{UserID: "42", Role: enums.UserRoleUser}
	err := CanMutate(actor, "42")
	if err != nil {
		t.Errorf("expected owner to pass, got %v", err)
	}
}

func TestCanMutate_NotOwnerForbidden(t *testing.T) {
	actor := &Actor{UserID: "1", Role: enums.UserRoleUser}
	err := CanMutate(actor, "2")
	if err == nil {
		t.Error("expected forbidden error")
	} else if err.Code != 403 {
		t.Errorf("expected 403, got %d", err.Code)
	}
}

func TestCanMutate_NilActor(t *testing.T) {
	err := CanMutate(nil, "1")
	if err == nil || err.Code != 401 {
		t.Errorf("expected 401 unauthorized, got %v", err)
	}
}

func TestCanRead_AdminPasses(t *testing.T) {
	actor := &Actor{UserID: "1", Role: enums.UserRoleAdmin}
	err := CanRead(actor, "999")
	if err != nil {
		t.Errorf("expected admin to pass, got %v", err)
	}
}

func TestCanRead_OwnerPasses(t *testing.T) {
	actor := &Actor{UserID: "42", Role: enums.UserRoleUser}
	err := CanRead(actor, "42")
	if err != nil {
		t.Errorf("expected owner to pass, got %v", err)
	}
}

func TestCanRead_NotOwnerForbidden(t *testing.T) {
	actor := &Actor{UserID: "1", Role: enums.UserRoleUser}
	err := CanRead(actor, "2")
	if err == nil {
		t.Error("expected forbidden error")
	} else if err.Code != 403 {
		t.Errorf("expected 403, got %d", err.Code)
	}
}

func TestCanRead_NilActor(t *testing.T) {
	err := CanRead(nil, "1")
	if err == nil || err.Code != 401 {
		t.Errorf("expected 401 unauthorized, got %v", err)
	}
}
