package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type ILearningHistoryRepo interface {
	Upsert(ctx context.Context, history *models.LearningHistory) error
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.LearningHistory], error)
	FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error)
}
