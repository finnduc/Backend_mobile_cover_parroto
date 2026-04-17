package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type ITranscriptRepo interface {
	FindByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, error)
}

type transcriptRepo struct {
	db *gorm.DB
}

func NewTranscriptRepo(db *gorm.DB) ITranscriptRepo {
	return &transcriptRepo{db: db}
}

func (r *transcriptRepo) FindByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, error) {
	var transcripts []*models.Transcript
	err := r.db.WithContext(ctx).Where("lesson_id = ?", lessonID).Order("sequence ASC").Find(&transcripts).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return transcripts, nil
}
