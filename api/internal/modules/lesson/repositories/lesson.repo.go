package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type ILessonRepo interface {
	FindAll(ctx context.Context) ([]models.Lesson, error)
	FindByID(ctx context.Context, id uint) (*models.Lesson, error)
}

type lessonRepo struct {
	db *gorm.DB
}

func NewLessonRepo(db *gorm.DB) ILessonRepo {
	return &lessonRepo{db: db}
}

func (r *lessonRepo) FindAll(ctx context.Context) ([]models.Lesson, error) {
	var lessons []models.Lesson
	err := r.db.WithContext(ctx).Find(&lessons).Error
	return lessons, err
}

func (r *lessonRepo) FindByID(ctx context.Context, id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := r.db.WithContext(ctx).First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}
