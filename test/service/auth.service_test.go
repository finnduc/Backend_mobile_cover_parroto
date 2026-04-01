package service

import (
	"context"
	"go-familytree/internal/mocks"
	"go-familytree/internal/models"
	"go-familytree/internal/service"
	"go-familytree/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockRepo := new(mocks.MockAuthRepo)
	s := service.NewAuthService(mockRepo, nil)

	input := service.RegisterInput{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	// Mock FindUserByEmail to return not found (empty user)
	mockRepo.On("FindUserByEmail", mock.Anything, input.Email).Return(&models.User{}, nil)

	// Mock CreateUser
	mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(u *models.User) bool {
		return u.Email == input.Email && u.Name == input.Name
	})).Return(nil)

	user, err := s.Register(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.Email, user.Email)
	assert.Equal(t, input.Name, user.Name)
	mockRepo.AssertExpectations(t)
}

func TestRegisterDuplicateEmail(t *testing.T) {
	mockRepo := new(mocks.MockAuthRepo)
	s := service.NewAuthService(mockRepo, nil)

	input := service.RegisterInput{
		Email:    "duplicate@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	// Mock FindUserByEmail to return an existing user
	existingUser := &models.User{Base: models.Base{ID: 1}, Email: input.Email}
	mockRepo.On("FindUserByEmail", mock.Anything, input.Email).Return(existingUser, nil)

	user, err := s.Register(context.Background(), input)

	assert.Error(t, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockRepo := new(mocks.MockAuthRepo)
	s := service.NewAuthService(mockRepo, nil)

	password := "password123"
	email := "test@example.com"
	hashedPassword, _ := utils.HashPassword(password)

	mockUser := &models.User{
		Base:         models.Base{ID: 1},
		Email:        email,
		PasswordHash: hashedPassword,
	}

	mockRepo.On("FindUserByEmail", mock.Anything, email).Return(mockUser, nil)
	mockRepo.On("GetUserRoles", mock.Anything, mockUser.ID).Return([]models.Role{}, nil)

	input := service.LoginInput{
		Email:    email,
		Password: password,
	}

	tokenPair, err := s.Login(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, tokenPair)
	assert.NotEmpty(t, tokenPair.AccessToken)
	assert.NotEmpty(t, tokenPair.RefreshToken)
	mockRepo.AssertExpectations(t)
}
