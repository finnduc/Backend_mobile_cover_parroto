package repositories

import (
	"context"
	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"gorm.io/gorm"
)

type transcriptBookmarkRepo struct {
	db *gorm.DB
}

func NewTranscriptBookmarkRepo(db *gorm.DB) db_repos.ITranscriptBookmarkRepo {
	return &transcriptBookmarkRepo{db: db}
}

func (r *transcriptBookmarkRepo) Create(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	return r.db.WithContext(ctx).Create(bookmark).Error
}

func (r *transcriptBookmarkRepo) FindByID(ctx context.Context, id uint) (*models.TranscriptBookmark, error) {
	var bookmark models.TranscriptBookmark
	err := r.db.WithContext(ctx).Preload("Transcript").Preload("Lesson").First(&bookmark, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &bookmark, nil
}

func (r *transcriptBookmarkRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.TranscriptBookmark, error) {
	var bookmark models.TranscriptBookmark
	err := r.db.WithContext(ctx).
		Preload("Transcript").
		Preload("Lesson").
		Where("user_id = ? AND transcript_id = ?", userID, transcriptID).
		First(&bookmark).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &bookmark, nil
}

func (r *transcriptBookmarkRepo) FindAllByUser(ctx context.Context, userID string, lessonID *uint) ([]*models.TranscriptBookmark, error) {
	var bookmarks []*models.TranscriptBookmark
	query := r.db.WithContext(ctx).
		Preload("Transcript").
		Preload("Lesson").
		Where("user_id = ?", userID)
	if lessonID != nil {
		query = query.Where("lesson_id = ?", *lessonID)
	}
	err := query.Find(&bookmarks).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return bookmarks, nil
}

func (r *transcriptBookmarkRepo) Update(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	return r.db.WithContext(ctx).Save(bookmark).Error
}

func (r *transcriptBookmarkRepo) DeleteByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND transcript_id = ?", userID, transcriptID).
		Delete(&models.TranscriptBookmark{}).Error
}
