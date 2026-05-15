package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"go-cover-parroto/internal/core/enums"
	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	lhreq "go-cover-parroto/internal/modules/learning_history/dtos/req"
	"go-cover-parroto/internal/modules/learning_history/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func testCtxLH(userID string, role enums.UserRole) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)
	return ctx
}

func TestLearningHistoryService_Record(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		body     lhreq.RecordHistoryReq
		setup    func(*repositories.MockLearningHistoryRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			ctx:  testCtxLH("1", enums.UserRoleUser),
			body: lhreq.RecordHistoryReq{LessonID: 3, DurationWatched: 120.5, Completed: true},
			setup: func(r *repositories.MockLearningHistoryRepo) {
				r.On("Upsert", mock.Anything, mock.MatchedBy(func(h *models.LearningHistory) bool {
					return h.UserID == "1" && h.LessonID == 3
				})).Return(nil)
			},
		},
		{
			name: "db error returns 500",
			ctx:  testCtxLH("1", enums.UserRoleUser),
			body: lhreq.RecordHistoryReq{LessonID: 3, DurationWatched: 10.0},
			setup: func(r *repositories.MockLearningHistoryRepo) {
				r.On("Upsert", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockLearningHistoryRepo)
			tt.setup(mockRepo)
			svc := NewLearningHistoryService(mockRepo)

			result, err := svc.Record(tt.ctx, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, result)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestLearningHistoryService_List(t *testing.T) {
	tests := []struct {
		name    string
		query   lhreq.ListHistoryQuery
		setup   func(*repositories.MockLearningHistoryRepo)
		wantLen int
	}{
		{
			name:  "success",
			query: lhreq.ListHistoryQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockLearningHistoryRepo) {
				paginated := &response.PaginatedResult[*models.LearningHistory]{
					Data: []*models.LearningHistory{
						{ID: 1, UserID: "1", LessonID: 2, DurationWatched: 60.0, Completed: false},
						{ID: 2, UserID: "1", LessonID: 3, DurationWatched: 90.0, Completed: true},
					},
					Meta: response.NewMeta(1, 10, 2),
				}
				r.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)
			},
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockLearningHistoryRepo)
			tt.setup(mockRepo)
			svc := NewLearningHistoryService(mockRepo)

			result, err := svc.List(context.Background(), tt.query)

			assert.Nil(t, err)
			assert.Len(t, result.Data, tt.wantLen)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestLearningHistoryService_GetByLesson(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		body     lhreq.GetHistoryReq
		setup    func(*repositories.MockLearningHistoryRepo)
		wantID   uint
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			ctx:  testCtxLH("1", enums.UserRoleUser),
			body: lhreq.GetHistoryReq{LessonID: 5},
			setup: func(r *repositories.MockLearningHistoryRepo) {
				r.On("FindByUserAndLesson", mock.Anything, "1", uint(5)).Return(
					&models.LearningHistory{ID: 1, UserID: "1", LessonID: 5, DurationWatched: 45.0, Completed: false}, nil)
			},
			wantID: 5,
		},
		{
			name: "not found returns 404",
			ctx:  testCtxLH("1", enums.UserRoleUser),
			body: lhreq.GetHistoryReq{LessonID: 99},
			setup: func(r *repositories.MockLearningHistoryRepo) {
				r.On("FindByUserAndLesson", mock.Anything, "1", uint(99)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockLearningHistoryRepo)
			tt.setup(mockRepo)
			svc := NewLearningHistoryService(mockRepo)

			result, err := svc.GetByLesson(tt.ctx, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantID, result.LessonID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
