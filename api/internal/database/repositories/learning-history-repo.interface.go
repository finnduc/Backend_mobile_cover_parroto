package repositories

import (
	"context"
	"go-cover-parroto/internal/database/models"
)

type ILearningHistoryRepo interface {
	Upsert(ctx context.Context, history *models.LearningHistory) error
	FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error)
	FindByUser(ctx context.Context, userID string) ([]*models.LearningHistory, error)
	FindFinished(ctx context.Context, userID string) ([]*models.LearningHistory, error)
	FindUnfinished(ctx context.Context, userID string) ([]*models.LearningHistory, error)
	CountFinished(ctx context.Context, userID string) (int64, error)
	CountUnfinished(ctx context.Context, userID string) (int64, error)
}
