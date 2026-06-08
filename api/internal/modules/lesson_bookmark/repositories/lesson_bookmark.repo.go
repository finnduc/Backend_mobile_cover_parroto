package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
)

type lessonBookmarkRepo struct {
	db *gorm.DB
}

func NewLessonBookmarkRepo(db *gorm.DB) db_repos.ILessonBookmarkRepo {
	return &lessonBookmarkRepo{db: db}
}

func (r *lessonBookmarkRepo) Create(ctx context.Context, bookmark *models.LessonBookmark) error {
	return r.db.WithContext(ctx).Create(bookmark).Error
}

func (r *lessonBookmarkRepo) FindByID(ctx context.Context, id uint) (*models.LessonBookmark, error) {
	var bookmark models.LessonBookmark
	err := r.db.WithContext(ctx).Preload("Lesson").First(&bookmark, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &bookmark, nil
}

func (r *lessonBookmarkRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LessonBookmark, error) {
	var bookmark models.LessonBookmark
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		First(&bookmark).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &bookmark, nil
}

func (r *lessonBookmarkRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.LessonBookmark], error) {
	var bookmarks []*models.LessonBookmark

	base := r.db.WithContext(ctx).Model(&models.LessonBookmark{}).Preload("Lesson")

	var total int64
	query.Count(base).Count(&total)

	err := query.Apply(base).Find(&bookmarks).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.LessonBookmark]{
		Data: bookmarks,
		Meta: meta,
	}, nil
}

func (r *lessonBookmarkRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.LessonBookmark{}, id).Error
}
