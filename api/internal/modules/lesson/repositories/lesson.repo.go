package repositories

import (
	"context"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type ILessonRepo interface {
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Lesson], error)
	FindByID(ctx context.Context, id uint) (*models.Lesson, error)
}

type lessonRepo struct {
	db *gorm.DB
}

func NewLessonRepo(db *gorm.DB) ILessonRepo {
	return &lessonRepo{db: db}
}

func (r *lessonRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Lesson], error) {
	var lessons []*models.Lesson

	countQuery := query.Count(r.db.Model(&models.Lesson{}))
	var total int64
	countQuery.Count(&total)

	result := query.Apply(r.db.WithContext(ctx).Model(&models.Lesson{}))
	err := result.Find(&lessons).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.Lesson]{
		Data: lessons,
		Meta: meta,
	}, nil
}

func (r *lessonRepo) FindByID(ctx context.Context, id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := r.db.WithContext(ctx).First(&lesson, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &lesson, nil
}
