package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/database"
	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"

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

func TestListLessons_Success(t *testing.T) {
	repo := new(mockLessonRepo)
	svc := NewLessonService(repo)

	lessons := []*models.Lesson{
		{ID: 1, Title: "Lesson 1", VideoURL: "https://vid1.mp4"},
		{ID: 2, Title: "Lesson 2", VideoURL: "https://vid2.mp4"},
	}
	paginatedResult := &response.PaginatedResult[*models.Lesson]{
		Data: lessons,
		Meta: response.NewMeta(1, 10, 2),
	}
	query := database.NewQuery().SetPage(1).SetLimit(10)
	repo.On("FindAll", mock.Anything, query).Return(paginatedResult, nil)

	result, appErr := svc.ListLessons(context.Background(), query)

	assert.Nil(t, appErr)
	assert.Len(t, result.Data, 2)
	assert.Equal(t, int64(2), result.Meta.Total)
	repo.AssertExpectations(t)
}

func TestListLessons_Error(t *testing.T) {
	repo := new(mockLessonRepo)
	svc := NewLessonService(repo)

	query := database.NewQuery()
	repo.On("FindAll", mock.Anything, query).Return(nil, errors.New("db error"))

	result, appErr := svc.ListLessons(context.Background(), query)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}

func TestGetLesson_Success(t *testing.T) {
	repo := new(mockLessonRepo)
	svc := NewLessonService(repo)

	lesson := &models.Lesson{ID: 1, Title: "Lesson 1", VideoURL: "https://vid.mp4", Level: "beginner"}
	repo.On("FindByID", mock.Anything, uint(1)).Return(lesson, nil)

	result, appErr := svc.GetLesson(context.Background(), 1)

	assert.Nil(t, appErr)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "Lesson 1", result.Title)
	assert.Equal(t, "beginner", result.Level)
	repo.AssertExpectations(t)
}

func TestGetLesson_NotFound(t *testing.T) {
	repo := new(mockLessonRepo)
	svc := NewLessonService(repo)

	repo.On("FindByID", mock.Anything, uint(99)).Return(nil, coreError.ErrNotFound)

	result, appErr := svc.GetLesson(context.Background(), 99)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 404, appErr.Code)
}
