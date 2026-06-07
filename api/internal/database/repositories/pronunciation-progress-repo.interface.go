package repositories

import (
	"context"
	"go-cover-parroto/internal/database/models"
)

type IPronunciationProgressRepo interface {
	Upsert(ctx context.Context, progress *models.PronunciationProgress) error
	FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) ([]*models.PronunciationProgress, error)
	FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.PronunciationProgress, error)
}
