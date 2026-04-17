package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type IBookmarkRepo interface {
	Create(ctx context.Context, userID, lessonID uint) error
	Delete(ctx context.Context, userID, lessonID uint) error
}

type bookmarkRepo struct {
	db *gorm.DB
}

func NewBookmarkRepo(db *gorm.DB) IBookmarkRepo {
	return &bookmarkRepo{db: db}
}

func (r *bookmarkRepo) Create(ctx context.Context, userID, lessonID uint) error {
	bookmark := &models.Bookmark{
		UserID:   userID,
		LessonID: lessonID,
	}
	return r.db.WithContext(ctx).Create(bookmark).Error
}

func (r *bookmarkRepo) Delete(ctx context.Context, userID, lessonID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ? AND lesson_id = ?", userID, lessonID).Delete(&models.Bookmark{}).Error
}
