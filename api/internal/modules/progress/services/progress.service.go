package services

import (
	"context"
	"time"

	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/progress/dtos/res"
	"go-cover-parroto/internal/modules/progress/repositories"

	"gorm.io/gorm"
)

var progressRepo repositories.IProgressRepo

func UpdateProgress(userID, lessonID uint, durationWatched float64, completed bool) (*res.ProgressRes, error) {
	progress, err := progressRepo.Upsert(context.Background(), userID, lessonID, durationWatched, completed)
	if err != nil {
		return nil, err
	}
	return toProgressRes(progress), nil
}

func GetProgress(userID, lessonID uint) (*res.ProgressRes, error) {
	progress, err := progressRepo.FindByUserAndLesson(context.Background(), userID, lessonID)
	if err != nil {
		return nil, err
	}
	return toProgressRes(progress), nil
}

func toProgressRes(progress *models.LearningHistory) *res.ProgressRes {
	return &res.ProgressRes{
		UserID:          progress.UserID,
		LessonID:        progress.LessonID,
		DurationWatched: progress.DurationWatched,
		Completed:       progress.Completed,
		CreatedAt:       progress.CreatedAt.Format(time.RFC3339),
	}
}

func SetProgressRepo(db *gorm.DB) {
	progressRepo = repositories.NewProgressRepo(db)
}
