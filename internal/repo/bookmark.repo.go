package repo

import (
	"context"
	"errors"
	"go-familytree/internal/models"
	"go-familytree/pkg/utils"

	"gorm.io/gorm"
)

type bookmarkRepo struct{ db *gorm.DB }

func NewBookmarkRepo(db *gorm.DB) IBookmarkRepo {
	return &bookmarkRepo{db: db}
}

// Toggle adds bookmark if not exists, removes it if exists. Returns new state.
func (r *bookmarkRepo) Toggle(ctx context.Context, userID, lessonID uint) (bool, error) {
	var bm models.Bookmark
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&bm).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create bookmark
		newBm := models.Bookmark{UserID: userID, LessonID: lessonID}
		return true, r.db.WithContext(ctx).Create(&newBm).Error
	}
	if err != nil {
		return false, err
	}
	// Delete bookmark
	return false, r.db.WithContext(ctx).
		Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		Delete(&models.Bookmark{}).Error
}

func (r *bookmarkRepo) FindByUser(ctx context.Context, userID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	q.Normalize()
	var lessons []models.Lesson
	var total int64
	tx := r.db.WithContext(ctx).
		Joins("JOIN bookmarks b ON b.lesson_id = lessons.id").
		Where("b.user_id = ?", userID)

	if err := tx.Model(&models.Lesson{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := tx.Offset(q.Offset()).Limit(q.Limit).Find(&lessons).Error
	return lessons, total, err
}
