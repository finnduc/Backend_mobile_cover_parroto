package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/category/dtos/req"
	"go-cover-parroto/internal/modules/category/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListCategories_Success(t *testing.T) {
	repo := new(repositories.MockCategoryRepo)
	svc := NewCategoryService(repo)

	categories := []*models.Category{
		{ID: 1, Name: "Science"},
		{ID: 2, Name: "History"},
	}
	paginatedResult := &response.PaginatedResult[*models.Category]{
		Data: categories,
		Meta: response.NewMeta(1, 10, 2),
	}
	query := req.ListCategoryQuery{Page: 1, Limit: 10}
	dbQuery := query.ToQuery()
	repo.On("FindAll", mock.Anything, dbQuery).Return(paginatedResult, nil)

	result, appErr := svc.List(context.Background(), query)

	assert.Nil(t, appErr)
	assert.Len(t, result.Data, 2)
	assert.Equal(t, "Science", result.Data[0].Name)
	repo.AssertExpectations(t)
}

func TestListCategories_Error(t *testing.T) {
	repo := new(repositories.MockCategoryRepo)
	svc := NewCategoryService(repo)

	query := req.ListCategoryQuery{Page: 1, Limit: 10}
	dbQuery := query.ToQuery()
	repo.On("FindAll", mock.Anything, dbQuery).Return(nil, errors.New("db error"))

	result, appErr := svc.List(context.Background(), query)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}
