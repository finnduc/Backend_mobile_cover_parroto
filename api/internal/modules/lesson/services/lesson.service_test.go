package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/lesson/dtos/req"
	"go-cover-parroto/internal/modules/lesson/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLessonService_List(t *testing.T) {
	tests := []struct {
		name     string
		query    req.ListLessonQuery
		setup    func(*repositories.MockLessonRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:  "success",
			query: req.ListLessonQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockLessonRepo) {
				paginated := &response.PaginatedResult[*models.Lesson]{
					Data: []*models.Lesson{{ID: 1, Title: "Lesson 1"}},
					Meta: response.NewMeta(1, 10, 1),
				}
				r.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)
			},
			wantLen: 1,
		},
		{
			name:  "db error returns 500",
			query: req.ListLessonQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockLessonRepo)
			tt.setup(mockRepo)
			svc := NewLessonService(mockRepo)

			result, err := svc.List(context.Background(), tt.query)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Len(t, result.Data, tt.wantLen)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestLessonService_Get(t *testing.T) {
	tests := []struct {
		name     string
		body     req.GetLessonReq
		setup    func(*repositories.MockLessonRepo)
		wantID   uint
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			body: req.GetLessonReq{ID: 1},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(&models.Lesson{ID: 1, Title: "Lesson 1"}, nil)
			},
			wantID: 1,
		},
		{
			name: "not found returns 404",
			body: req.GetLessonReq{ID: 999},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockLessonRepo)
			tt.setup(mockRepo)
			svc := NewLessonService(mockRepo)

			result, err := svc.Get(context.Background(), tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantID, result.ID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
