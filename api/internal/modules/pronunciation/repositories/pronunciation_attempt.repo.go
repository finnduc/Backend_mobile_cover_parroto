package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
)

type pronunciationAttemptRepo struct {
	db *gorm.DB
}

func NewPronunciationAttemptRepo(db *gorm.DB) db_repos.IPronunciationAttemptRepo {
	return &pronunciationAttemptRepo{db: db}
}

func (r *pronunciationAttemptRepo) Create(ctx context.Context, attempt *models.PronunciationAttempt) error {
	return r.db.WithContext(ctx).Create(attempt).Error
}

func (r *pronunciationAttemptRepo) FindByID(ctx context.Context, id uint) (*models.PronunciationAttempt, error) {
	var attempt models.PronunciationAttempt
	err := r.db.WithContext(ctx).First(&attempt, id).Error
	if err != nil {
		return nil, err
	}
	return &attempt, nil
}

func (r *pronunciationAttemptRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) ([]*models.PronunciationAttempt, error) {
	var attempts []*models.PronunciationAttempt
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND transcript_id = ?", userID, transcriptID).
		Order("created_at DESC").
		Find(&attempts).Error
	return attempts, err
}

func (r *pronunciationAttemptRepo) FindBestByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.PronunciationAttempt, error) {
	var attempt models.PronunciationAttempt
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND transcript_id = ?", userID, transcriptID).
		Order("overall_score DESC").
		First(&attempt).Error
	if err != nil {
		return nil, err
	}
	return &attempt, nil
}

func (r *pronunciationAttemptRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.PronunciationAttempt{}, id).Error
}
