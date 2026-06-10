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
	"go-cover-parroto/internal/modules/shadowing_status/dtos/req"
	"go-cover-parroto/internal/modules/shadowing_status/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestShadowingStatusService_Create(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		body     req.CreateShadowingStatusReq
		setup    func(*repositories.MockShadowingStatusRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name:   "success",
			userID: "user1",
			body:   req.CreateShadowingStatusReq{TranscriptID: 5},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				r.On("FindByUserAndTranscript", mock.Anything, "user1", uint(5)).
					Return(nil, gorm.ErrRecordNotFound)
				r.On("Create", mock.Anything, mock.MatchedBy(func(s *models.ShadowingStatus) bool {
					return s.UserID == "user1" && s.TranscriptID == 5
				})).Return(nil)
			},
		},
		{
			name:   "duplicate transcript returns 409",
			userID: "user1",
			body:   req.CreateShadowingStatusReq{TranscriptID: 5},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				existing := &models.ShadowingStatus{UserID: "user1", TranscriptID: 5}
				r.On("FindByUserAndTranscript", mock.Anything, "user1", uint(5)).
					Return(existing, nil)
			},
			wantErr:  true,
			wantCode: http.StatusConflict,
		},
		{
			name:   "find db error returns 500",
			userID: "user1",
			body:   req.CreateShadowingStatusReq{TranscriptID: 5},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				r.On("FindByUserAndTranscript", mock.Anything, "user1", uint(5)).
					Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
		{
			name:   "create db error returns 500",
			userID: "user1",
			body:   req.CreateShadowingStatusReq{TranscriptID: 5},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				r.On("FindByUserAndTranscript", mock.Anything, "user1", uint(5)).
					Return(nil, gorm.ErrRecordNotFound)
				r.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockShadowingStatusRepo)
			tt.setup(mockRepo)
			svc := NewShadowingStatusService(mockRepo)

			result, err := svc.Create(testCtx("user1", enums.UserRoleUser), tt.userID, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "user1", result.UserID)
				assert.Equal(t, uint(5), result.TranscriptID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestShadowingStatusService_List(t *testing.T) {
	lessonID := uint(10)

	tests := []struct {
		name     string
		ctx      context.Context
		query    req.ListShadowingStatusQuery
		setup    func(*repositories.MockShadowingStatusRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:  "success without filter",
			ctx:   testCtx("user1", enums.UserRoleUser),
			query: req.ListShadowingStatusQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				paginated := &response.PaginatedResult[*models.ShadowingStatus]{
					Data: []*models.ShadowingStatus{
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
			query: req.ListShadowingStatusQuery{LessonID: &lessonID, Page: 1, Limit: 5},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				paginated := &response.PaginatedResult[*models.ShadowingStatus]{
					Data: []*models.ShadowingStatus{
						{UserID: "user1", TranscriptID: 3, LessonID: 10},
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
			query: req.ListShadowingStatusQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				paginated := &response.PaginatedResult[*models.ShadowingStatus]{
					Data: []*models.ShadowingStatus{},
					Meta: response.NewMeta(1, 10, 0),
				}
				r.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)
			},
			wantLen: 0,
		},
		{
			name:  "db error returns 500",
			ctx:   testCtx("user1", enums.UserRoleUser),
			query: req.ListShadowingStatusQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockShadowingStatusRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockShadowingStatusRepo)
			tt.setup(mockRepo)
			svc := NewShadowingStatusService(mockRepo)

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
