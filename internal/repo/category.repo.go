package repo

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/pkg/utils"

	"gorm.io/gorm"
)

type categoryRepo struct{ db *gorm.DB }

func NewCategoryRepo(db *gorm.DB) ICategoryRepo {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) FindAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (r *categoryRepo) FindLessonsByCategory(ctx context.Context, categoryID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	q.Normalize()
	var lessons []models.Lesson
	var total int64

	tx := r.db.WithContext(ctx).
		Joins("JOIN lesson_categories lc ON lc.lesson_id = lessons.id").
		Where("lc.category_id = ?", categoryID).
		Model(&models.Lesson{})

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := tx.Offset(q.Offset()).Limit(q.Limit).Find(&lessons).Error
	return lessons, total, err
}
