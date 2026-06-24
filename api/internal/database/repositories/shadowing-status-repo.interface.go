package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type IShadowingStatusRepo interface {
	Create(ctx context.Context, status *models.ShadowingStatus) error
	Update(ctx context.Context, status *models.ShadowingStatus) error
	FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.ShadowingStatus, error)
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.ShadowingStatus], error)
}
