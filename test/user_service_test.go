package test

import (
	"context"
	"testing"

	"go-familytree/internal/repo"
	"go-familytree/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGetFamilyService(t *testing.T) {
	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	family := userService.GetFamilyService()
	
	expected := []string{"Father", "mother", "brother"}
	assert.ElementsMatch(t, expected, family, "The family members should match")
}

func TestRegisterService(t *testing.T) {
	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	
	ctx := context.Background()
	user, err := userService.RegisterService(ctx, "test@example.com", "study")
	
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, 1, user.ID)
}
