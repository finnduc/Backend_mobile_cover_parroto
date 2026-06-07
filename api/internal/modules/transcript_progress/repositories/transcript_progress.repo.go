package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type transcriptProgressRepo struct {
	db *gorm.DB
}

func NewTranscriptProgressRepo(db *gorm.DB) db_repos.ITranscriptProgressRepo {
	return &transcriptProgressRepo{db: db}
}

func (r *transcriptProgressRepo) CreateOrIgnore(ctx context.Context, progress *models.TranscriptProgress) error {
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(progress).Error
}

func (r *transcriptProgressRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) ([]*models.TranscriptProgress, error) {
	var progresses []*models.TranscriptProgress
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		Find(&progresses).Error
	return progresses, err
}
