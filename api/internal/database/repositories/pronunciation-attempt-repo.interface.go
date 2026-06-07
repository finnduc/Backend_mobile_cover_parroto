package repositories

import (
	"context"
	"go-cover-parroto/internal/database/models"
)

type IPronunciationAttemptRepo interface {
	Create(ctx context.Context, attempt *models.PronunciationAttempt) error
	FindByID(ctx context.Context, id uint) (*models.PronunciationAttempt, error)
	FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) ([]*models.PronunciationAttempt, error)
	FindBestByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.PronunciationAttempt, error)
	Delete(ctx context.Context, id uint) error
}
