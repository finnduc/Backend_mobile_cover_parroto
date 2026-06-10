package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/dictation_status/dtos/req"
	"go-cover-parroto/internal/modules/dictation_status/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDictationStatusService_Create(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		body     req.CreateDictationStatusReq
		setup    func(*repositories.MockDictationStatusRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name:   "success",
			userID: "user1",
			body:   req.CreateDictationStatusReq{TranscriptID: 7},
			setup: func(r *repositories.MockDictationStatusRepo) {
				r.On("CreateOrIgnore", mock.Anything, mock.MatchedBy(func(s *models.DictationStatus) bool {
					return s.UserID == "user1" && s.TranscriptID == 7
				})).Return(nil)
			},
		},
		{
			name:   "duplicate transcript is silently ignored",
			userID: "user1",
			body:   req.CreateDictationStatusReq{TranscriptID: 7},
			setup: func(r *repositories.MockDictationStatusRepo) {
				// CreateOrIgnore handles conflicts at DB level — no error returned
				r.On("CreateOrIgnore", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name:   "db error returns 500",
			userID: "user1",
			body:   req.CreateDictationStatusReq{TranscriptID: 7},
			setup: func(r *repositories.MockDictationStatusRepo) {
				r.On("CreateOrIgnore", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockDictationStatusRepo)
			tt.setup(mockRepo)
			svc := NewDictationStatusService(mockRepo)

			result, err := svc.Create(context.Background(), tt.userID, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "user1", result.UserID)
				assert.Equal(t, uint(7), result.TranscriptID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDictationStatusService_List(t *testing.T) {
	lessonID := uint(10)

	tests := []struct {
		name     string
		ctx      context.Context
		query    req.ListDictationStatusQuery
		setup    func(*repositories.MockDictationStatusRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:  "success without filter",
			ctx:   testCtx("user1", enums.UserRoleUser),
			query: req.ListDictationStatusQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockDictationStatusRepo) {
				paginated := &response.PaginatedResult[*models.DictationStatus]{
					Data: []*models.DictationStatus{
						{UserID: "user1", TranscriptID: 1, LessonID: 10},
						{UserID: "user1", TranscriptID: 2, LessonID: 10},
					},
					Meta: response.NewMeta(1, 10, 2),
				}
				r.On("FindAll", mock.Anything, mock.MatchedBy(func(q *database.Query) bool {
					return true
				})).Return(paginated, nil)
			},
			wantLen: 2,
		},
		{
			name:  "success with lesson_id filter",
			ctx:   testCtx("user1", enums.UserRoleUser),
			query: req.ListDictationStatusQuery{LessonID: &lessonID, Page: 1, Limit: 5},
			setup: func(r *repositories.MockDictationStatusRepo) {
				paginated := &response.PaginatedResult[*models.DictationStatus]{
					Data: []*models.DictationStatus{
						{UserID: "user1", TranscriptID: 4, LessonID: 10},
					},
					Meta: response.NewMeta(1, 5, 1),
				}
				r.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)
			},
			wantLen: 1,
		},
		{
			name:  "empty result",
			ctx:   testCtx("user1", enums.UserRoleUser),
			query: req.ListDictationStatusQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockDictationStatusRepo) {
				paginated := &response.PaginatedResult[*models.DictationStatus]{
					Data: []*models.DictationStatus{},
					Meta: response.NewMeta(1, 10, 0),
				}
				r.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)
			},
			wantLen: 0,
		},
		{
			name:  "db error returns 500",
			ctx:   testCtx("user1", enums.UserRoleUser),
			query: req.ListDictationStatusQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockDictationStatusRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockDictationStatusRepo)
			tt.setup(mockRepo)
			svc := NewDictationStatusService(mockRepo)

			result, err := svc.List(tt.ctx, tt.query)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.Len(t, result.Data, tt.wantLen)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
