package repositories

import (
	"context"
	"go-cover-parroto/internal/database/models"
)

type ITranscriptProgressRepo interface {
	CreateOrIgnore(ctx context.Context, progress *models.TranscriptProgress) error
	FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) ([]*models.TranscriptProgress, error)
}
