package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/database"
	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	lhreq "go-cover-parroto/internal/modules/learning_history/dtos/req"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockLHRepo struct{ mock.Mock }

func (m *mockLHRepo) Upsert(ctx context.Context, history *models.LearningHistory) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *mockLHRepo) FindAllByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResult[*models.LearningHistory], error) {
	args := m.Called(ctx, userID, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.LearningHistory]), args.Error(1)
}

func (m *mockLHRepo) FindByUserAndLesson(ctx context.Context, userID, lessonID uint) (*models.LearningHistory, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.LearningHistory), args.Error(1)
}

func TestRecord_Success(t *testing.T) {
	repo := new(mockLHRepo)
	svc := NewLearningHistoryService(repo)

	body := lhreq.RecordHistoryReq{LessonID: 3, DurationWatched: 120.5, Completed: true}
	repo.On("Upsert", mock.Anything, mock.AnythingOfType("*models.LearningHistory")).Return(nil)

	result, appErr := svc.Record(context.Background(), 1, body)

	assert.Nil(t, appErr)
	assert.Equal(t, uint(1), result.UserID)
	assert.Equal(t, uint(3), result.LessonID)
	assert.Equal(t, 120.5, result.DurationWatched)
	assert.True(t, result.Completed)
	repo.AssertExpectations(t)
}

func TestRecord_UpsertError(t *testing.T) {
	repo := new(mockLHRepo)
	svc := NewLearningHistoryService(repo)

	body := lhreq.RecordHistoryReq{LessonID: 3, DurationWatched: 10.0}
	repo.On("Upsert", mock.Anything, mock.AnythingOfType("*models.LearningHistory")).Return(errors.New("db error"))

	result, appErr := svc.Record(context.Background(), 1, body)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}

func TestListByUser_Success(t *testing.T) {
	repo := new(mockLHRepo)
	svc := NewLearningHistoryService(repo)

	histories := []*models.LearningHistory{
		{ID: 1, UserID: 1, LessonID: 2, DurationWatched: 60.0, Completed: false},
		{ID: 2, UserID: 1, LessonID: 3, DurationWatched: 90.0, Completed: true},
	}
	query := database.NewQuery().SetPage(1).SetLimit(10)
	paginatedResult := &response.PaginatedResult[*models.LearningHistory]{
		Data: histories,
		Meta: response.NewMeta(1, 10, 2),
	}
	repo.On("FindAllByUser", mock.Anything, uint(1), query).Return(paginatedResult, nil)

	result, appErr := svc.ListByUser(context.Background(), 1, query)

	assert.Nil(t, appErr)
	assert.Len(t, result.Data, 2)
	assert.Equal(t, int64(2), result.Meta.Total)
}

func TestGetByLesson_Success(t *testing.T) {
	repo := new(mockLHRepo)
	svc := NewLearningHistoryService(repo)

	history := &models.LearningHistory{ID: 1, UserID: 1, LessonID: 5, DurationWatched: 45.0, Completed: false}
	repo.On("FindByUserAndLesson", mock.Anything, uint(1), uint(5)).Return(history, nil)

	result, appErr := svc.GetByLesson(context.Background(), 1, 5)

	assert.Nil(t, appErr)
	assert.Equal(t, uint(5), result.LessonID)
	assert.Equal(t, 45.0, result.DurationWatched)
}

func TestGetByLesson_NotFound(t *testing.T) {
	repo := new(mockLHRepo)
	svc := NewLearningHistoryService(repo)

	repo.On("FindByUserAndLesson", mock.Anything, uint(1), uint(99)).Return(nil, coreError.ErrNotFound)

	result, appErr := svc.GetByLesson(context.Background(), 1, 99)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 404, appErr.Code)
}
