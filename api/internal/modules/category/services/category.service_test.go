package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

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
				paginated := &response.PaginatedResult[*models.Category]{
					Data: []*models.Category{
						{ID: 1, Name: "Science"},
						{ID: 2, Name: "History"},
					},
					Meta: response.NewMeta(1, 10, 2),
				}
				r.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)
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
