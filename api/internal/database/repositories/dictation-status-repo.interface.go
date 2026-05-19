package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type IDictationStatusRepo interface {
	CreateOrIgnore(ctx context.Context, status *models.DictationStatus) error
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.DictationStatus], error)
}
