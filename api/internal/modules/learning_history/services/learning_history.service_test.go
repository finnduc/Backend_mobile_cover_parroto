package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/enums"
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

func (m *mockLHRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.LearningHistory], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.LearningHistory]), args.Error(1)
}

func (m *mockLHRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.LearningHistory), args.Error(1)
}

func testCtx(userID string, role enums.UserRole) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)
	return ctx
}

func TestRecord_Success(t *testing.T) {
	mockRepo := new(mockLHRepo)
	svc := NewLearningHistoryService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := lhreq.RecordHistoryReq{LessonID: 3, DurationWatched: 120.5, Completed: true}

	mockRepo.On("Upsert", mock.Anything, mock.MatchedBy(func(h *models.LearningHistory) bool {
		return h.UserID == "1" && h.LessonID == 3
	})).Return(nil)

	result, err := svc.Record(ctx, body)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestRecord_UpsertError(t *testing.T) {
	mockRepo := new(mockLHRepo)
	svc := NewLearningHistoryService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := lhreq.RecordHistoryReq{LessonID: 3, DurationWatched: 10.0}

	mockRepo.On("Upsert", mock.Anything, mock.Anything).Return(errors.New("db error"))

	_, err := svc.Record(ctx, body)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Code)
	mockRepo.AssertExpectations(t)
}

func TestList_Success(t *testing.T) {
	mockRepo := new(mockLHRepo)
	svc := NewLearningHistoryService(mockRepo)
	q := lhreq.ListHistoryQuery{Page: 1, Limit: 10}

	paginated := &response.PaginatedResult[*models.LearningHistory]{
		Data: []*models.LearningHistory{
			{ID: 1, UserID: "1", LessonID: 2, DurationWatched: 60.0, Completed: false},
			{ID: 2, UserID: "1", LessonID: 3, DurationWatched: 90.0, Completed: true},
		},
		Meta: response.NewMeta(1, 10, 2),
	}
	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)

	result, err := svc.List(context.Background(), q)
	assert.Nil(t, err)
	assert.Len(t, result.Data, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetByLesson_Success(t *testing.T) {
	mockRepo := new(mockLHRepo)
	svc := NewLearningHistoryService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := lhreq.GetHistoryReq{LessonID: 5}

	history := &models.LearningHistory{ID: 1, UserID: "1", LessonID: 5, DurationWatched: 45.0, Completed: false}
	mockRepo.On("FindByUserAndLesson", mock.Anything, "1", uint(5)).Return(history, nil)

	result, err := svc.GetByLesson(ctx, body)
	assert.Nil(t, err)
	assert.Equal(t, uint(5), result.LessonID)
	mockRepo.AssertExpectations(t)
}

func TestGetByLesson_NotFound(t *testing.T) {
	mockRepo := new(mockLHRepo)
	svc := NewLearningHistoryService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := lhreq.GetHistoryReq{LessonID: 99}

	mockRepo.On("FindByUserAndLesson", mock.Anything, "1", uint(99)).Return(nil, coreError.ErrNotFound)

	_, err := svc.GetByLesson(ctx, body)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Code)
	mockRepo.AssertExpectations(t)
}
