package services

import (
	"context"
	"errors"
	"testing"

	firebaseauth "firebase.google.com/go/v4/auth"
	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- mocks ---

type mockAuthRepo struct{ mock.Mock }

func (m *mockAuthRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockAuthRepo) Create(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

type mockFirebaseAuth struct{ mock.Mock }

func (m *mockFirebaseAuth) VerifyIDToken(ctx context.Context, idToken string) (*firebaseauth.Token, error) {
	args := m.Called(ctx, idToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*firebaseauth.Token), args.Error(1)
}

// --- helpers ---

func makeToken(email, name, picture string) *firebaseauth.Token {
	return &firebaseauth.Token{
		Claims: map[string]interface{}{
			"email":   email,
			"name":    name,
			"picture": picture,
		},
	}
}

// --- tests ---

func TestSyncUser_NewUser(t *testing.T) {
	repo := new(mockAuthRepo)
	fb := new(mockFirebaseAuth)
	svc := NewAuthService(repo, fb)

	token := makeToken("new@example.com", "New User", "https://pic.jpg")
	fb.On("VerifyIDToken", mock.Anything, "valid-token").Return(token, nil)
	repo.On("FindByEmail", mock.Anything, "new@example.com").Return(nil, coreError.ErrNotFound)
	repo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	result, appErr := svc.SyncUser(context.Background(), "valid-token")

	assert.Nil(t, appErr)
	assert.Equal(t, "new@example.com", result.Email)
	assert.Equal(t, "New User", result.Name)
	repo.AssertExpectations(t)
	fb.AssertExpectations(t)
}

func TestSyncUser_ExistingUser(t *testing.T) {
	repo := new(mockAuthRepo)
	fb := new(mockFirebaseAuth)
	svc := NewAuthService(repo, fb)

	existing := &models.User{ID: 1, Email: "existing@example.com", Name: "Existing"}
	token := makeToken("existing@example.com", "Existing", "")
	fb.On("VerifyIDToken", mock.Anything, "valid-token").Return(token, nil)
	repo.On("FindByEmail", mock.Anything, "existing@example.com").Return(existing, nil)

	result, appErr := svc.SyncUser(context.Background(), "valid-token")

	assert.Nil(t, appErr)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "existing@example.com", result.Email)
	repo.AssertExpectations(t)
}

func TestSyncUser_InvalidToken(t *testing.T) {
	repo := new(mockAuthRepo)
	fb := new(mockFirebaseAuth)
	svc := NewAuthService(repo, fb)

	fb.On("VerifyIDToken", mock.Anything, "bad-token").Return(nil, errors.New("invalid token"))

	result, appErr := svc.SyncUser(context.Background(), "bad-token")

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 401, appErr.Code)
}

func TestSyncUser_MissingEmail(t *testing.T) {
	repo := new(mockAuthRepo)
	fb := new(mockFirebaseAuth)
	svc := NewAuthService(repo, fb)

	token := &firebaseauth.Token{Claims: map[string]interface{}{}}
	fb.On("VerifyIDToken", mock.Anything, "no-email-token").Return(token, nil)

	result, appErr := svc.SyncUser(context.Background(), "no-email-token")

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 400, appErr.Code)
}

func TestSyncUser_CreateError(t *testing.T) {
	repo := new(mockAuthRepo)
	fb := new(mockFirebaseAuth)
	svc := NewAuthService(repo, fb)

	token := makeToken("err@example.com", "Err", "")
	fb.On("VerifyIDToken", mock.Anything, "valid-token").Return(token, nil)
	repo.On("FindByEmail", mock.Anything, "err@example.com").Return(nil, coreError.ErrNotFound)
	repo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(errors.New("db error"))

	result, appErr := svc.SyncUser(context.Background(), "valid-token")

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}
