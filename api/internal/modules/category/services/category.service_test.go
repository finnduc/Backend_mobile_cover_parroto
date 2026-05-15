package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/category/dtos/req"
	"go-cover-parroto/internal/modules/category/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCategoryService_List(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*repositories.MockCategoryRepo)
		wantLen  int
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(
					&response.PaginatedResult[*models.Category]{
						Data: []*models.Category{{ID: 1, Name: "Science"}, {ID: 2, Name: "History"}},
						Meta: response.NewMeta(1, 10, 2),
					}, nil)
			},
			wantLen:  2,
			wantName: "Science",
		},
		{
			name: "db error returns 500",
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(repositories.MockCategoryRepo)
			tt.setup(repo)
			svc := NewCategoryService(repo)

			result, appErr := svc.List(context.Background(), req.ListCategoryQuery{Page: 1, Limit: 10})

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Len(t, result.Data, tt.wantLen)
				assert.Equal(t, tt.wantName, result.Data[0].Name)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestCategoryService_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockCategoryRepo)
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(&models.Category{ID: 1, Name: "Science"}, nil)
			},
			wantName: "Science",
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(repositories.MockCategoryRepo)
			tt.setup(repo)
			svc := NewCategoryService(repo)

			result, appErr := svc.GetByID(context.Background(), tt.id)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Equal(t, tt.wantName, result.Name)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestCategoryService_Create(t *testing.T) {
	tests := []struct {
		name     string
		body     req.CreateCategoryReq
		setup    func(*repositories.MockCategoryRepo)
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			body: req.CreateCategoryReq{Name: "New Category"},
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("Create", mock.Anything, mock.MatchedBy(func(c *models.Category) bool {
					return c.Name == "New Category"
				})).Return(nil)
			},
			wantName: "New Category",
		},
		{
			name: "db error returns 500",
			body: req.CreateCategoryReq{Name: "Fail"},
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(repositories.MockCategoryRepo)
			tt.setup(repo)
			svc := NewCategoryService(repo)

			result, appErr := svc.Create(context.Background(), tt.body)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Equal(t, tt.wantName, result.Name)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestCategoryService_Update(t *testing.T) {
	existing := &models.Category{ID: 1, Name: "Original"}
	tests := []struct {
		name     string
		id       uint
		body     req.UpdateCategoryReq
		setup    func(*repositories.MockCategoryRepo)
		wantName string
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			body: req.UpdateCategoryReq{Name: "Updated"},
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.MatchedBy(func(c *models.Category) bool {
					return c.Name == "Updated"
				})).Return(nil)
			},
			wantName: "Updated",
		},
		{
			name: "not found returns 404",
			id:   999,
			body: req.UpdateCategoryReq{Name: "Nope"},
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "update error returns 500",
			id:   1,
			body: req.UpdateCategoryReq{Name: "Updated"},
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.Anything).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(repositories.MockCategoryRepo)
			tt.setup(repo)
			svc := NewCategoryService(repo)

			result, appErr := svc.Update(context.Background(), tt.id, tt.body)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Equal(t, tt.wantName, result.Name)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestCategoryService_Delete(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockCategoryRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("Delete", mock.Anything, uint(1)).Return(nil)
			},
		},
		{
			name: "db error returns 500",
			id:   2,
			setup: func(r *repositories.MockCategoryRepo) {
				r.On("Delete", mock.Anything, uint(2)).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(repositories.MockCategoryRepo)
			tt.setup(repo)
			svc := NewCategoryService(repo)

			appErr := svc.Delete(context.Background(), tt.id)

			if tt.wantErr {
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
			}
			repo.AssertExpectations(t)
		})
	}
}
