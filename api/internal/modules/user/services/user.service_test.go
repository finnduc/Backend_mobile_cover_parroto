package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/enums"
	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/database/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct{ mock.Mock }

func (m *mockUserRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.User], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.User]), args.Error(1)
}

func (m *mockUserRepo) Update(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockUserRepo) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func testCtx(userID string, role enums.UserRole) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)
	return ctx
}

func TestGetProfile_Success(t *testing.T) {
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)

	user := &models.User{ID: "1", Email: "test@example.com", Name: "Test"}
	mockRepo.On("FindByID", mock.Anything, "1").Return(user, nil)

	result, err := svc.GetProfile(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "test@example.com", result.Email)
	mockRepo.AssertExpectations(t)
}

func TestGetProfile_NotFound(t *testing.T) {
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := testCtx("999", enums.UserRoleUser)

	mockRepo.On("FindByID", mock.Anything, "999").Return(nil, coreError.ErrNotFound)

	_, err := svc.GetProfile(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetProfile_InternalError(t *testing.T) {
	mockRepo := new(mockUserRepo)
	svc := NewUserService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)

	mockRepo.On("FindByID", mock.Anything, "1").Return(nil, errors.New("db error"))

	_, err := svc.GetProfile(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Code)
	mockRepo.AssertExpectations(t)
}
