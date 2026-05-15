package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type ILessonRepo interface {
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Lesson], error)
	FindByID(ctx context.Context, id uint) (*models.Lesson, error)
	Create(ctx context.Context, lesson *models.Lesson) error
	Update(ctx context.Context, lesson *models.Lesson) error
	Delete(ctx context.Context, id uint) error
}
