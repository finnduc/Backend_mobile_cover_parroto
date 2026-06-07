package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type ITranscriptBookmarkRepo interface {
	Create(ctx context.Context, bookmark *models.TranscriptBookmark) error
	FindByID(ctx context.Context, id uint) (*models.TranscriptBookmark, error)
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.TranscriptBookmark], error)
	Update(ctx context.Context, bookmark *models.TranscriptBookmark) error
	Delete(ctx context.Context, id uint) error
}
