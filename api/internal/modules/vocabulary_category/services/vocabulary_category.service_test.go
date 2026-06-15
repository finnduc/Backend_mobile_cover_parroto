package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/vocabulary_category/dtos/req"
	"go-cover-parroto/internal/modules/vocabulary_category/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVocabularyCategoryService_List(t *testing.T) {
	tests := []struct {
		name     string
		query    req.ListVocabularyCategoryQuery
		setup    func(*repositories.MockVocabularyCategoryRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:  "success — returns all categories",
			query: req.ListVocabularyCategoryQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(
					&response.PaginatedResult[*models.VocabularyCategory]{
						Data: []*models.VocabularyCategory{
							{ID: 1, Name: "Grammar"},
							{ID: 2, Name: "Daily Life"},
						},
						Meta: response.NewMeta(1, 10, 2),
					}, nil)
			},
			wantLen: 2,
		},
		{
			name:  "success — empty list",
			query: req.ListVocabularyCategoryQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(
					&response.PaginatedResult[*models.VocabularyCategory]{
						Data: []*models.VocabularyCategory{},
						Meta: response.NewMeta(1, 10, 0),
					}, nil)
			},
			wantLen: 0,
		},
		{
			name:  "db error returns 500",
			query: req.ListVocabularyCategoryQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyCategoryRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyCategoryService(mockRepo)

			result, err := svc.List(context.Background(), tt.query)

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

func TestVocabularyCategoryService_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockVocabularyCategoryRepo)
		wantID   uint
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(
					&models.VocabularyCategory{ID: 1, Name: "Grammar", Description: "Grammar rules"}, nil)
			},
			wantID: 1,
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "db error returns 500",
			id:   1,
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyCategoryRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyCategoryService(mockRepo)

			result, err := svc.GetByID(context.Background(), tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantID, result.ID)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyCategoryService_Create(t *testing.T) {
	tests := []struct {
		name     string
		body     req.CreateVocabularyCategoryReq
		setup    func(*repositories.MockVocabularyCategoryRepo)
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			body: req.CreateVocabularyCategoryReq{Name: "Science", Description: "Science vocabulary"},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("Create", mock.Anything, mock.MatchedBy(func(c *models.VocabularyCategory) bool {
					return c.Name == "Science"
				})).Run(func(args mock.Arguments) {
					args.Get(1).(*models.VocabularyCategory).ID = 5
				}).Return(nil)
			},
			wantName: "Science",
		},
		{
			name: "db error returns 500",
			body: req.CreateVocabularyCategoryReq{Name: "Fail"},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyCategoryRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyCategoryService(mockRepo)

			result, err := svc.Create(context.Background(), tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantName, result.Name)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyCategoryService_Update(t *testing.T) {
	existing := &models.VocabularyCategory{ID: 1, Name: "Old Name", Description: "Old desc"}
	strPtr := func(s string) *string { return &s }

	tests := []struct {
		name     string
		id       uint
		body     req.UpdateVocabularyCategoryReq
		setup    func(*repositories.MockVocabularyCategoryRepo)
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			body: req.UpdateVocabularyCategoryReq{Name: strPtr("New Name"), Description: strPtr("New desc")},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.MatchedBy(func(c *models.VocabularyCategory) bool {
					return c.Name == "New Name"
				})).Return(nil)
			},
			wantName: "New Name",
		},
		{
			name: "not found returns 404",
			id:   999,
			body: req.UpdateVocabularyCategoryReq{Name: strPtr("X")},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "update db error returns 500",
			id:   1,
			body: req.UpdateVocabularyCategoryReq{Name: strPtr("New Name")},
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyCategoryRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyCategoryService(mockRepo)

			result, err := svc.Update(context.Background(), tt.id, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.wantName, result.Name)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestVocabularyCategoryService_Delete(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockVocabularyCategoryRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(&models.VocabularyCategory{ID: 1}, nil)
				r.On("Delete", mock.Anything, uint(1)).Return(nil)
			},
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "db error returns 500",
			id:   2,
			setup: func(r *repositories.MockVocabularyCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(2)).Return(&models.VocabularyCategory{ID: 2}, nil)
				r.On("Delete", mock.Anything, uint(2)).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockVocabularyCategoryRepo)
			tt.setup(mockRepo)
			svc := NewVocabularyCategoryService(mockRepo)

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
