package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/category/dtos/req"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCategoryRepo struct{ mock.Mock }

func (m *mockCategoryRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Category], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.Category]), args.Error(1)
}

func (m *mockCategoryRepo) FindByID(ctx context.Context, id uint) (*models.Category, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func TestListCategories_Success(t *testing.T) {
	repo := new(mockCategoryRepo)
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
	repo := new(mockCategoryRepo)
	svc := NewCategoryService(repo)

	query := req.ListCategoryQuery{Page: 1, Limit: 10}
	dbQuery := query.ToQuery()
	repo.On("FindAll", mock.Anything, dbQuery).Return(nil, errors.New("db error"))

	result, appErr := svc.List(context.Background(), query)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}
