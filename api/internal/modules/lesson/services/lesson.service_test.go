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
				r.On("FindAll", mock.Anything, mock.Anything).Return(
					&response.PaginatedResult[*models.Lesson]{
						Data: []*models.Lesson{{ID: 1, Title: "Lesson 1"}},
						Meta: response.NewMeta(1, 10, 1),
					}, nil)
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

func TestLessonService_Create(t *testing.T) {
	tests := []struct {
		name     string
		body     req.CreateLessonReq
		setup    func(*repositories.MockLessonRepo)
		wantID   uint
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			body: req.CreateLessonReq{Title: "New Lesson", VideoURL: "http://example.com"},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("Create", mock.Anything, mock.MatchedBy(func(l *models.Lesson) bool {
					return l.Title == "New Lesson"
				})).Run(func(args mock.Arguments) {
					args.Get(1).(*models.Lesson).ID = 1
				}).Return(nil)
			},
			wantID: 1,
		},
		{
			name: "db error returns 500",
			body: req.CreateLessonReq{Title: "Fail", VideoURL: "http://example.com"},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
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

			result, err := svc.Create(context.Background(), tt.body)

			if tt.wantErr {
				assert.Nil(t, result)
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

func TestLessonService_Update(t *testing.T) {
	existing := &models.Lesson{ID: 1, Title: "Original"}
	tests := []struct {
		name     string
		id       uint
		body     req.UpdateLessonReq
		setup    func(*repositories.MockLessonRepo)
		wantID   uint
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			body: req.UpdateLessonReq{Title: "Updated"},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.MatchedBy(func(l *models.Lesson) bool {
					return l.Title == "Updated"
				})).Return(nil)
			},
			wantID: 1,
		},
		{
			name: "not found returns 404",
			id:   999,
			body: req.UpdateLessonReq{Title: "Nope"},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "update error returns 500",
			id:   1,
			body: req.UpdateLessonReq{Title: "Updated"},
			setup: func(r *repositories.MockLessonRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.Anything).Return(errors.New("db error"))
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

			result, err := svc.Update(context.Background(), tt.id, tt.body)

			if tt.wantErr {
				assert.Nil(t, result)
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

func TestLessonService_Delete(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockLessonRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockLessonRepo) {
				r.On("Delete", mock.Anything, uint(1)).Return(nil)
			},
		},
		{
			name: "db error returns 500",
			id:   2,
			setup: func(r *repositories.MockLessonRepo) {
				r.On("Delete", mock.Anything, uint(2)).Return(errors.New("db error"))
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

			err := svc.Delete(context.Background(), tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
