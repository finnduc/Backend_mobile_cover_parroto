package services

import (
	"context"
	"errors"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct{ mock.Mock }

func (m *mockUserRepo) FindByID(ctx context.Context, id uint) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func TestGetProfile_Success(t *testing.T) {
	repo := new(mockUserRepo)
	svc := NewUserService(repo)

	user := &models.User{ID: 1, Email: "test@example.com", Name: "Test", AvatarURL: "https://avatar.jpg"}
	repo.On("FindByID", mock.Anything, uint(1)).Return(user, nil)

	result, appErr := svc.GetProfile(context.Background(), 1)

	assert.Nil(t, appErr)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "test@example.com", result.Email)
	assert.Equal(t, "Test", result.Name)
	repo.AssertExpectations(t)
}

func TestGetProfile_NotFound(t *testing.T) {
	repo := new(mockUserRepo)
	svc := NewUserService(repo)

	repo.On("FindByID", mock.Anything, uint(99)).Return(nil, coreError.ErrNotFound)

	result, appErr := svc.GetProfile(context.Background(), 99)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 404, appErr.Code)
	repo.AssertExpectations(t)
}

func TestGetProfile_InternalError(t *testing.T) {
	repo := new(mockUserRepo)
	svc := NewUserService(repo)

	repo.On("FindByID", mock.Anything, uint(1)).Return(nil, errors.New("db connection error"))

	result, appErr := svc.GetProfile(context.Background(), 1)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
	repo.AssertExpectations(t)
}
