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
	return errors.MapRepoError(r.db.WithContext(ctx).Create(bookmark).Error)
}

func (r *transcriptBookmarkRepo) Update(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	return errors.MapRepoError(r.db.WithContext(ctx).Save(bookmark).Error)
}

func (r *transcriptBookmarkRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.TranscriptBookmark, error) {
	var bookmark models.TranscriptBookmark
	err := r.db.WithContext(ctx).
		Preload("Lesson").
		Preload("Transcript").
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
		Preload("Lesson").
		Preload("Transcript").
		Where("user_id = ?", userID).
		Order("lesson_id ASC, transcript_id ASC")
	if lessonID != nil {
		query = query.Where("lesson_id = ?", *lessonID)
	}
	if err := query.Find(&bookmarks).Error; err != nil {
		return nil, errors.MapRepoError(err)
	}
	return bookmarks, nil
}

func (r *transcriptBookmarkRepo) DeleteByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) error {
	return errors.MapRepoError(r.db.WithContext(ctx).
		Where("user_id = ? AND transcript_id = ?", userID, transcriptID).
		Delete(&models.TranscriptBookmark{}).Error)
}
