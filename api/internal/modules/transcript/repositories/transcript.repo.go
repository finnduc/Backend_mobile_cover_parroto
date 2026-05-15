package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
)

type transcriptRepo struct {
	db *gorm.DB
}

func NewTranscriptRepo(db *gorm.DB) db_repos.ITranscriptRepo {
	return &transcriptRepo{db: db}
}

func (r *transcriptRepo) FindByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, error) {
	var transcripts []*models.Transcript
	err := r.db.WithContext(ctx).Where("lesson_id = ?", lessonID).Order("sequence ASC").Find(&transcripts).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return transcripts, nil
}

func (r *transcriptRepo) FindByID(ctx context.Context, id uint) (*models.Transcript, error) {
	var transcript models.Transcript
	err := r.db.WithContext(ctx).First(&transcript, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &transcript, nil
}

func (r *transcriptRepo) Create(ctx context.Context, transcript *models.Transcript) error {
	return r.db.WithContext(ctx).Create(transcript).Error
}

func (r *transcriptRepo) BulkCreate(ctx context.Context, transcripts []*models.Transcript) error {
	return r.db.WithContext(ctx).Create(transcripts).Error
}

func (r *transcriptRepo) Update(ctx context.Context, transcript *models.Transcript) error {
	return r.db.WithContext(ctx).Save(transcript).Error
}

func (r *transcriptRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Transcript{}, id).Error
}

func (r *transcriptRepo) DeleteByLesson(ctx context.Context, lessonID uint) error {
	return r.db.WithContext(ctx).Where("lesson_id = ?", lessonID).Delete(&models.Transcript{}).Error
}
