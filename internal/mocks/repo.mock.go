package mocks

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/pkg/utils"

	"github.com/stretchr/testify/mock"
)

// MockAuthRepo
type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockAuthRepo) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockAuthRepo) GetUserRoles(ctx context.Context, userID uint) ([]models.Role, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]models.Role), args.Error(1)
}

// MockAttemptRepo
type MockAttemptRepo struct {
	mock.Mock
}

func (m *MockAttemptRepo) Create(ctx context.Context, attempt *models.Attempt) error {
	args := m.Called(ctx, attempt)
	return args.Error(0)
}

func (m *MockAttemptRepo) FindByID(ctx context.Context, id uint) (*models.Attempt, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Attempt), args.Error(1)
}

func (m *MockAttemptRepo) CountAnswers(ctx context.Context, attemptID uint) (int64, error) {
	args := m.Called(ctx, attemptID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAttemptRepo) MarkCompleted(ctx context.Context, attemptID uint, totalScore float64) error {
	args := m.Called(ctx, attemptID, totalScore)
	return args.Error(0)
}

// MockAnswerRepo
type MockAnswerRepo struct {
	mock.Mock
}

func (m *MockAnswerRepo) Upsert(ctx context.Context, answer *models.UserAnswer) error {
	args := m.Called(ctx, answer)
	return args.Error(0)
}

func (m *MockAnswerRepo) BulkUpsert(ctx context.Context, answers []models.UserAnswer) error {
	args := m.Called(ctx, answers)
	return args.Error(0)
}

func (m *MockAnswerRepo) FindByAttemptAndTranscript(ctx context.Context, attemptID, transcriptID uint) (*models.UserAnswer, error) {
	args := m.Called(ctx, attemptID, transcriptID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserAnswer), args.Error(1)
}

// MockProgressRepo
type MockProgressRepo struct {
	mock.Mock
}

func (m *MockProgressRepo) FindOrCreate(ctx context.Context, userID, lessonID uint) (*models.UserProgress, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserProgress), args.Error(1)
}

func (m *MockProgressRepo) Save(ctx context.Context, p *models.UserProgress) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

// MockBookmarkRepo
type MockBookmarkRepo struct {
	mock.Mock
}

func (m *MockBookmarkRepo) Toggle(ctx context.Context, userID, lessonID uint) (bool, error) {
	args := m.Called(ctx, userID, lessonID)
	return args.Bool(0), args.Error(1)
}

func (m *MockBookmarkRepo) FindByUser(ctx context.Context, userID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	args := m.Called(ctx, userID, q)
	return args.Get(0).([]models.Lesson), args.Get(1).(int64), args.Error(2)
}

// MockCategoryRepo
type MockCategoryRepo struct {
	mock.Mock
}

func (m *MockCategoryRepo) FindAll(ctx context.Context) ([]models.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryRepo) FindLessonsByCategory(ctx context.Context, categoryID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	args := m.Called(ctx, categoryID, q)
	return args.Get(0).([]models.Lesson), args.Get(1).(int64), args.Error(2)
}
