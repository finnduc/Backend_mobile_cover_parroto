package repo

import (
	"context"
	"go-cover-parroto/internal/models"

	"gorm.io/gorm"
)

type attemptRepo struct{ db *gorm.DB }

func NewAttemptRepo(db *gorm.DB) IAttemptRepo {
	return &attemptRepo{db: db}
}

func (r *attemptRepo) Create(ctx context.Context, attempt *models.Attempt) error {
	return r.db.WithContext(ctx).Create(attempt).Error
}

func (r *attemptRepo) FindByID(ctx context.Context, id uint) (*models.Attempt, error) {
	var attempt models.Attempt
	err := r.db.WithContext(ctx).Preload("Answers").First(&attempt, id).Error
	return &attempt, err
}

func (r *attemptRepo) CountAnswers(ctx context.Context, attemptID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.UserAnswer{}).
		Where("attempt_id = ?", attemptID).Count(&count).Error
	return count, err
}

func (r *attemptRepo) MarkCompleted(ctx context.Context, attemptID uint, totalScore float64) error {
	return r.db.WithContext(ctx).Model(&models.Attempt{}).
		Where("id = ?", attemptID).
		Updates(map[string]interface{}{"completed": true, "total_score": totalScore}).Error
}
