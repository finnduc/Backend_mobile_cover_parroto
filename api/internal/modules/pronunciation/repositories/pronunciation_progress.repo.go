package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pronunciationProgressRepo struct {
	db *gorm.DB
}

func NewPronunciationProgressRepo(db *gorm.DB) db_repos.IPronunciationProgressRepo {
	return &pronunciationProgressRepo{db: db}
}

func (r *pronunciationProgressRepo) Upsert(ctx context.Context, progress *models.PronunciationProgress) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "transcript_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"lesson_id", "best_attempt_id", "best_score", "feedback", "updated_at"}),
		}).
		Create(progress).Error
}

func (r *pronunciationProgressRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) ([]*models.PronunciationProgress, error) {
	var progress []*models.PronunciationProgress
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		Find(&progress).Error
	return progress, err
}

func (r *pronunciationProgressRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.PronunciationProgress, error) {
	var progress models.PronunciationProgress
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND transcript_id = ?", userID, transcriptID).
		First(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}
