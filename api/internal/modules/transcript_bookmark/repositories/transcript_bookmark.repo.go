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
	err := r.db.WithContext(ctx).Preload("Transcript").First(&bookmark, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &bookmark, nil
}

func (r *transcriptBookmarkRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.TranscriptBookmark], error) {
	var bookmarks []*models.TranscriptBookmark

	base := r.db.WithContext(ctx).Model(&models.TranscriptBookmark{}).Preload("Transcript")

	if _, ok := query.Filters["lesson_id"]; ok {
		lessonID := query.Filters["lesson_id"]
		delete(query.Filters, "lesson_id")
		base = base.Joins("JOIN transcripts ON transcripts.id = transcript_bookmarks.transcript_id").
			Where("transcripts.lesson_id = ?", lessonID)
	}

	var total int64
	query.Count(base).Count(&total)

	err := query.Apply(base).Find(&bookmarks).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.TranscriptBookmark]{
		Data: bookmarks,
		Meta: meta,
	}, nil
}

func (r *transcriptBookmarkRepo) Update(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	return r.db.WithContext(ctx).Save(bookmark).Error
}

func (r *transcriptBookmarkRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.TranscriptBookmark{}, id).Error
}
