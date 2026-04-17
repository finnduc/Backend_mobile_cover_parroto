package repositories

import (
	"context"
	"errors"

	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("record not found")

type ICategoryRepo interface {
	FindAll(ctx context.Context) ([]models.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) ICategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) FindAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return categories, err
}
