package repositories

import (
	"context"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type IBookmarkRepo interface {
	Create(ctx context.Context, userID, lessonID uint) error
	Delete(ctx context.Context, userID, lessonID uint) error
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Bookmark], error)
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

func (r *bookmarkRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Bookmark], error) {
	var bookmarks []*models.Bookmark

	base := r.db.WithContext(ctx).Model(&models.Bookmark{}).Preload("Lesson")

	var total int64
	query.Count(base).Count(&total)

	err := query.Apply(base).Find(&bookmarks).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.Bookmark]{
		Data: bookmarks,
		Meta: meta,
	}, nil
}
