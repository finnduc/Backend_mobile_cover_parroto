package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/learning_history/dtos/req"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockLearningHistoryRepo struct {
	mock.Mock
}

func (m *mockLearningHistoryRepo) Create(ctx context.Context, history *models.LearningHistory) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *mockLearningHistoryRepo) Update(ctx context.Context, history *models.LearningHistory) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *mockLearningHistoryRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error) {
	args := m.Called(ctx, userID, lessonID)
	if history, ok := args.Get(0).(*models.LearningHistory); ok {
		return history, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockLearningHistoryRepo) FindAllByUser(ctx context.Context, userID string, filter string) ([]*models.LearningHistory, error) {
	args := m.Called(ctx, userID, filter)
	if histories, ok := args.Get(0).([]*models.LearningHistory); ok {
		return histories, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockLearningHistoryRepo) CountSummary(ctx context.Context, userID string) (int64, int64, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(int64), args.Get(1).(int64), args.Error(2)
}

func (m *mockLearningHistoryRepo) CountCompletedLessonTranscripts(ctx context.Context, userID string, lessonID uint) (int64, int64, error) {
	args := m.Called(ctx, userID, lessonID)
	return args.Get(0).(int64), args.Get(1).(int64), args.Error(2)
}

func boolPtr(v bool) *bool {
	return &v
}

func TestLearningHistoryService_CreateInsertsNewHistory(t *testing.T) {
	repo := new(mockLearningHistoryRepo)
	repo.On("FindByUserAndLesson", mock.Anything, "user1", uint(15)).Return(nil, coreError.ErrNotFound)
	repo.On("Create", mock.Anything, mock.MatchedBy(func(history *models.LearningHistory) bool {
		return history.UserID == "user1" && history.LessonID == 15 && history.CompletedDictation && !history.CompletedPronunciation
	})).Return(nil)

	result, err := NewLearningHistoryService(repo).Create(context.Background(), "user1", req.CreateLearningHistoryReq{
		LessonID:           15,
		CompletedDictation: boolPtr(true),
	})

	assert.Nil(t, err)
	assert.Equal(t, uint(15), result.LessonID)
	assert.True(t, result.CompletedDictation)
	assert.False(t, result.CompletedPronunciation)
	repo.AssertExpectations(t)
}

func TestLearningHistoryService_CreateUpdatesOnlyProvidedFields(t *testing.T) {
	repo := new(mockLearningHistoryRepo)
	existing := &models.LearningHistory{
		UserID:                 "user1",
		LessonID:               15,
		CompletedDictation:     true,
		CompletedPronunciation: false,
	}
	repo.On("FindByUserAndLesson", mock.Anything, "user1", uint(15)).Return(existing, nil)
	repo.On("Update", mock.Anything, mock.MatchedBy(func(history *models.LearningHistory) bool {
		return history.CompletedDictation && history.CompletedPronunciation
	})).Return(nil)

	result, err := NewLearningHistoryService(repo).Create(context.Background(), "user1", req.CreateLearningHistoryReq{
		LessonID:               15,
		CompletedPronunciation: boolPtr(true),
	})

	assert.Nil(t, err)
	assert.True(t, result.CompletedDictation)
	assert.True(t, result.CompletedPronunciation)
	repo.AssertExpectations(t)
}

func TestLearningHistoryService_ListFinished(t *testing.T) {
	repo := new(mockLearningHistoryRepo)
	repo.On("FindAllByUser", mock.Anything, "user1", "finished").Return([]*models.LearningHistory{
		{UserID: "user1", LessonID: 1, CompletedDictation: true, CompletedPronunciation: true},
	}, nil)

	result, err := NewLearningHistoryService(repo).List(context.Background(), "user1", "finished")

	assert.Nil(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(1), result[0].LessonID)
	repo.AssertExpectations(t)
}

func TestLearningHistoryService_Summary(t *testing.T) {
	repo := new(mockLearningHistoryRepo)
	repo.On("CountSummary", mock.Anything, "user1").Return(int64(3), int64(2), nil)

	result, err := NewLearningHistoryService(repo).Summary(context.Background(), "user1")

	assert.Nil(t, err)
	assert.Equal(t, int64(3), result.CompletedCount)
	assert.Equal(t, int64(2), result.UnfinishedCount)
	repo.AssertExpectations(t)
}

func TestLearningHistoryService_LessonSummaryClampsUncompleted(t *testing.T) {
	repo := new(mockLearningHistoryRepo)
	repo.On("CountCompletedLessonTranscripts", mock.Anything, "user1", uint(15)).Return(int64(5), int64(3), nil)

	result, err := NewLearningHistoryService(repo).LessonSummary(context.Background(), "user1", 15)

	assert.Nil(t, err)
	assert.Equal(t, int64(5), result.Completed)
	assert.Equal(t, int64(0), result.Uncompleted)
	repo.AssertExpectations(t)
}

func TestLearningHistoryService_GetByLessonNotFound(t *testing.T) {
	repo := new(mockLearningHistoryRepo)
	repo.On("FindByUserAndLesson", mock.Anything, "user1", uint(15)).Return(nil, coreError.ErrNotFound)

	result, err := NewLearningHistoryService(repo).GetByLesson(context.Background(), "user1", 15)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Code)
	repo.AssertExpectations(t)
}

func TestLearningHistoryService_CreateRepoError(t *testing.T) {
	repo := new(mockLearningHistoryRepo)
	repo.On("FindByUserAndLesson", mock.Anything, "user1", uint(15)).Return(nil, errors.New("db error"))

	result, err := NewLearningHistoryService(repo).Create(context.Background(), "user1", req.CreateLearningHistoryReq{LessonID: 15})

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Code)
	repo.AssertExpectations(t)
}
