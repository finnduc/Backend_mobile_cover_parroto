package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
)

type ITranscriptRepo interface {
	FindByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, error)
	FindByID(ctx context.Context, id uint) (*models.Transcript, error)
	Create(ctx context.Context, transcript *models.Transcript) error
	BulkCreate(ctx context.Context, transcripts []*models.Transcript) error
	Update(ctx context.Context, transcript *models.Transcript) error
	Delete(ctx context.Context, id uint) error
	DeleteByLesson(ctx context.Context, lessonID uint) error
}
