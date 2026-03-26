package test

import (
	"testing"

	"go-familytree/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGetInfoService(t *testing.T) {
	userService := service.NewUserService()
	info := userService.GetInfoService()
	
	assert.Equal(t, "Finn", info, "The info should be 'Finn'")
}

func TestGetFamilyService(t *testing.T) {
	userService := service.NewUserService()
	family := userService.GetFamilyService()
	
	expected := []string{"Father", "mother", "brother"}
	assert.ElementsMatch(t, expected, family, "The family members should match")
}
