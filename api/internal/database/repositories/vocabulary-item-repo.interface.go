package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type IVocabularyItemRepo interface {
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyItem], error)
	FindByID(ctx context.Context, id uint) (*models.VocabularyItem, error)
	Create(ctx context.Context, item *models.VocabularyItem) error
	Update(ctx context.Context, item *models.VocabularyItem) error
	Delete(ctx context.Context, id uint) error
}
