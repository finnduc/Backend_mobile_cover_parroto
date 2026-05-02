package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/enums"
	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/policy"
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

func testCtx(userID uint, role enums.UserRole) context.Context {
	ctx := context.WithValue(context.Background(), policy.ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, policy.ContextKeyUserRole, role)
	return ctx
}

func TestGetProfile_Success(t *testing.T) {
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := testCtx(1, enums.UserRoleUser)

	user := &models.User{ID: 1, Email: "test@example.com", Name: "Test"}
	mockRepo.On("FindByID", mock.Anything, uint(1)).Return(user, nil)

	result, err := svc.GetProfile(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "test@example.com", result.Email)
	mockRepo.AssertExpectations(t)
}

func TestGetProfile_NotFound(t *testing.T) {
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := testCtx(999, enums.UserRoleUser)

	mockRepo.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)

	_, err := svc.GetProfile(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetProfile_InternalError(t *testing.T) {
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := testCtx(1, enums.UserRoleUser)

	mockRepo.On("FindByID", mock.Anything, uint(1)).Return(nil, errors.New("db error"))

	_, err := svc.GetProfile(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Code)
	mockRepo.AssertExpectations(t)
}
