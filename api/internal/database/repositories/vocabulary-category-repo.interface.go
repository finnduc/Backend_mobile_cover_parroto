package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type IVocabularyCategoryRepo interface {
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyCategory], error)
	FindByID(ctx context.Context, id uint) (*models.VocabularyCategory, error)
	Create(ctx context.Context, category *models.VocabularyCategory) error
	Update(ctx context.Context, category *models.VocabularyCategory) error
	Delete(ctx context.Context, id uint) error
}
