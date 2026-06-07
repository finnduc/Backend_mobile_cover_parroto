package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type learningHistoryRepo struct {
	db *gorm.DB
}

func NewLearningHistoryRepo(db *gorm.DB) db_repos.ILearningHistoryRepo {
	return &learningHistoryRepo{db: db}
}

func (r *learningHistoryRepo) Upsert(ctx context.Context, history *models.LearningHistory) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "lesson_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"completed_dictation", "completed_pronunciation", "updated_at"}),
		}).
		Create(history).Error
}

func (r *learningHistoryRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error) {
	var history models.LearningHistory
	err := r.db.WithContext(ctx).Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&history).Error
	if err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *learningHistoryRepo) FindByUser(ctx context.Context, userID string) ([]*models.LearningHistory, error) {
	var history []*models.LearningHistory
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&history).Error
	return history, err
}

func (r *learningHistoryRepo) FindFinished(ctx context.Context, userID string) ([]*models.LearningHistory, error) {
	var history []*models.LearningHistory
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND completed_dictation = true AND completed_pronunciation = true", userID).
		Find(&history).Error
	return history, err
}

func (r *learningHistoryRepo) FindUnfinished(ctx context.Context, userID string) ([]*models.LearningHistory, error) {
	var history []*models.LearningHistory
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND (completed_dictation = false OR completed_pronunciation = false OR completed_pronunciation IS NULL)", userID).
		Find(&history).Error
	return history, err
}

func (r *learningHistoryRepo) CountFinished(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.LearningHistory{}).
		Where("user_id = ? AND completed_dictation = true AND completed_pronunciation = true", userID).
		Count(&count).Error
	return count, err
}

func (r *learningHistoryRepo) CountUnfinished(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.LearningHistory{}).
		Where("user_id = ? AND (completed_dictation = false OR completed_pronunciation = false OR completed_pronunciation IS NULL)", userID).
		Count(&count).Error
	return count, err
}
