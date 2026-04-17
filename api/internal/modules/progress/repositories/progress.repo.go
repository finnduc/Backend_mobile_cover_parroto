package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type IProgressRepo interface {
	Upsert(ctx context.Context, userID, lessonID uint, durationWatched float64, completed bool) (*models.LearningHistory, error)
	FindByUserAndLesson(ctx context.Context, userID, lessonID uint) (*models.LearningHistory, error)
}

type progressRepo struct {
	db *gorm.DB
}

func NewProgressRepo(db *gorm.DB) IProgressRepo {
	return &progressRepo{db: db}
}

func (r *progressRepo) Upsert(ctx context.Context, userID, lessonID uint, durationWatched float64, completed bool) (*models.LearningHistory, error) {
	var progress models.LearningHistory
	err := r.db.WithContext(ctx).Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&progress).Error
	if err != nil {
		progress = models.LearningHistory{
			UserID:          userID,
			LessonID:        lessonID,
			DurationWatched: durationWatched,
			Completed:       completed,
		}
		err = r.db.WithContext(ctx).Create(&progress).Error
		if err != nil {
			return nil, err
		}
		return &progress, nil
	}

	progress.DurationWatched = durationWatched
	progress.Completed = completed
	err = r.db.WithContext(ctx).Save(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *progressRepo) FindByUserAndLesson(ctx context.Context, userID, lessonID uint) (*models.LearningHistory, error) {
	var progress models.LearningHistory
	err := r.db.WithContext(ctx).Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}
