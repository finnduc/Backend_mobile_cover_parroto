package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type ILessonBookmarkRepo interface {
	Create(ctx context.Context, bookmark *models.LessonBookmark) error
	FindByID(ctx context.Context, id uint) (*models.LessonBookmark, error)
	FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LessonBookmark, error)
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.LessonBookmark], error)
	Delete(ctx context.Context, id uint) error
}
