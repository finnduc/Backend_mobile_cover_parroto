package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
)

type IBookmarkRepo interface {
	Create(ctx context.Context, userID string, lessonID uint) error
	Delete(ctx context.Context, userID string, lessonID uint) error
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Bookmark], error)
}
