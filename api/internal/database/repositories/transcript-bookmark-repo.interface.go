package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
)

type ITranscriptBookmarkRepo interface {
	Create(ctx context.Context, bookmark *models.TranscriptBookmark) error
	Update(ctx context.Context, bookmark *models.TranscriptBookmark) error
	FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.TranscriptBookmark, error)
	FindAllByUser(ctx context.Context, userID string, lessonID *uint) ([]*models.TranscriptBookmark, error)
	DeleteByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) error
}
