package services

import (
	"context"
	"errors"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/lesson/dtos/req"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockLessonRepo struct{ mock.Mock }

func (m *mockLessonRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Lesson], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.Lesson]), args.Error(1)
}

func (m *mockLessonRepo) FindByID(ctx context.Context, id uint) (*models.Lesson, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Lesson), args.Error(1)
}

func (m *mockLessonRepo) Create(ctx context.Context, lesson *models.Lesson) error {
	args := m.Called(ctx, lesson)
	return args.Error(0)
}

func (m *mockLessonRepo) Update(ctx context.Context, lesson *models.Lesson) error {
	args := m.Called(ctx, lesson)
	return args.Error(0)
}

func (m *mockLessonRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestListLessons_Success(t *testing.T) {
	mockRepo := new(mockLessonRepo)
	svc := NewLessonService(mockRepo)
	q := req.ListLessonQuery{Page: 1, Limit: 10}

	paginated := &response.PaginatedResult[*models.Lesson]{
		Data: []*models.Lesson{
			{ID: 1, Title: "Lesson 1"},
		},
		Meta: response.NewMeta(1, 10, 1),
	}

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)

	result, err := svc.List(context.Background(), q)
	assert.Nil(t, err)
	assert.Len(t, result.Data, 1)
	mockRepo.AssertExpectations(t)
}

func TestListLessons_Error(t *testing.T) {
	mockRepo := new(mockLessonRepo)
	svc := NewLessonService(mockRepo)
	q := req.ListLessonQuery{Page: 1, Limit: 10}

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	_, err := svc.List(context.Background(), q)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetLesson_Success(t *testing.T) {
	mockRepo := new(mockLessonRepo)
	svc := NewLessonService(mockRepo)
	body := req.GetLessonReq{ID: 1}

	lesson := &models.Lesson{ID: 1, Title: "Lesson 1"}
	mockRepo.On("FindByID", mock.Anything, uint(1)).Return(lesson, nil)

	result, err := svc.Get(context.Background(), body)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetLesson_NotFound(t *testing.T) {
	mockRepo := new(mockLessonRepo)
	svc := NewLessonService(mockRepo)
	body := req.GetLessonReq{ID: 999}

	mockRepo.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)

	_, err := svc.Get(context.Background(), body)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Code)
	mockRepo.AssertExpectations(t)
}
