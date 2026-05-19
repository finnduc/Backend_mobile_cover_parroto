package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type IVocabularyDeckRepo interface {
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyDeck], error)
	FindByID(ctx context.Context, id uint) (*models.VocabularyDeck, error)
	Create(ctx context.Context, deck *models.VocabularyDeck) error
	Update(ctx context.Context, deck *models.VocabularyDeck) error
	Delete(ctx context.Context, id uint) error
}
