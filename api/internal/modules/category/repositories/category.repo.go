package repositories

import (
	"context"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type ICategoryRepo interface {
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Category], error)
	FindByID(ctx context.Context, id uint) (*models.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) ICategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Category], error) {
	var categories []*models.Category

	countQuery := query.Count(r.db.Model(&models.Category{}))
	var total int64
	countQuery.Count(&total)

	result := query.Apply(r.db.WithContext(ctx).Model(&models.Category{}))
	err := result.Find(&categories).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.Category]{
		Data: categories,
		Meta: meta,
	}, nil
}

func (r *categoryRepo) FindByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &category, nil
}
