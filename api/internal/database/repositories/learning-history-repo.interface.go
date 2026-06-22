package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
)

type ILearningHistoryRepo interface {
	Create(ctx context.Context, history *models.LearningHistory) error
	Update(ctx context.Context, history *models.LearningHistory) error
	FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error)
	FindAllByUser(ctx context.Context, userID string, filter string) ([]*models.LearningHistory, error)
	CountSummary(ctx context.Context, userID string) (completed int64, unfinished int64, err error)
	CountCompletedLessonTranscripts(ctx context.Context, userID string, lessonID uint) (completed int64, total int64, err error)
}
