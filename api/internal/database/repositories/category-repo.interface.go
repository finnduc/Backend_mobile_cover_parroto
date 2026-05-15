package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type ICategoryRepo interface {
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Category], error)
	FindByID(ctx context.Context, id uint) (*models.Category, error)
	Create(ctx context.Context, category *models.Category) error
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id uint) error
}
