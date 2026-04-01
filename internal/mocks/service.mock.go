package mocks

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/internal/service"
	"go-familytree/pkg/utils"

	"github.com/stretchr/testify/mock"
)

// MockAuthService
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(ctx context.Context, input service.RegisterInput) (*models.User, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockAuthService) Login(ctx context.Context, input service.LoginInput) (*service.TokenPair, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.TokenPair), args.Error(1)
}

func (m *MockAuthService) Logout(ctx context.Context, accessToken string, userID uint) error {
	args := m.Called(ctx, accessToken, userID)
	return args.Error(0)
}

func (m *MockAuthService) Refresh(ctx context.Context, refreshToken string, userID uint) (*service.TokenPair, error) {
	args := m.Called(ctx, refreshToken, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.TokenPair), args.Error(1)
}

// MockLessonService
type MockLessonService struct {
	mock.Mock
}

func (m *MockLessonService) ListLessons(ctx context.Context, level string, categoryID *uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	args := m.Called(ctx, level, categoryID, q)
	return args.Get(0).([]models.Lesson), args.Get(1).(int64), args.Error(2)
}

func (m *MockLessonService) CreateLesson(ctx context.Context, input service.CreateLessonInput) (*models.Lesson, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Lesson), args.Error(1)
}

func (m *MockLessonService) GetLesson(ctx context.Context, id uint) (*models.Lesson, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Lesson), args.Error(1)
}

func (m *MockLessonService) GetTranscripts(ctx context.Context, lessonID uint) ([]models.Transcript, error) {
	args := m.Called(ctx, lessonID)
	return args.Get(0).([]models.Transcript), args.Error(1)
}

// MockAttemptService
type MockAttemptService struct {
	mock.Mock
}

func (m *MockAttemptService) CreateAttempt(ctx context.Context, userID, lessonID uint) (*models.Attempt, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Attempt), args.Error(1)
}

func (m *MockAttemptService) GetAttempt(ctx context.Context, id uint) (*models.Attempt, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Attempt), args.Error(1)
}

// MockAnswerService
type MockAnswerService struct {
	mock.Mock
}

func (m *MockAnswerService) SubmitAnswer(ctx context.Context, userID uint, input service.SubmitAnswerInput) (*service.AnswerResult, error) {
	args := m.Called(ctx, userID, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.AnswerResult), args.Error(1)
}

func (m *MockAnswerService) BulkSubmit(ctx context.Context, userID uint, input service.BulkSubmitInput) (*service.BulkAnswerResult, error) {
	args := m.Called(ctx, userID, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.BulkAnswerResult), args.Error(1)
}

// MockProgressService
type MockProgressService struct {
	mock.Mock
}

func (m *MockProgressService) GetProgress(ctx context.Context, userID, lessonID uint) (*models.UserProgress, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserProgress), args.Error(1)
}

// MockBookmarkService
type MockBookmarkService struct {
	mock.Mock
}

func (m *MockBookmarkService) Toggle(ctx context.Context, userID, lessonID uint) (bool, error) {
	args := m.Called(ctx, userID, lessonID)
	return args.Bool(0), args.Error(1)
}

func (m *MockBookmarkService) ListBookmarks(ctx context.Context, userID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	args := m.Called(ctx, userID, q)
	return args.Get(0).([]models.Lesson), args.Get(1).(int64), args.Error(2)
}

// MockCategoryService
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) ListCategories(ctx context.Context) ([]models.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryService) GetLessonsByCategory(ctx context.Context, categoryID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	args := m.Called(ctx, categoryID, q)
	return args.Get(0).([]models.Lesson), args.Get(1).(int64), args.Error(2)
}
