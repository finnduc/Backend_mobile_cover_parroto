package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
)

type shadowingStatusRepo struct {
	db *gorm.DB
}

func NewShadowingStatusRepo(db *gorm.DB) db_repos.IShadowingStatusRepo {
	return &shadowingStatusRepo{db: db}
}

func (r *shadowingStatusRepo) Create(ctx context.Context, status *models.ShadowingStatus) error {
	return r.db.WithContext(ctx).Create(status).Error
}

func (r *shadowingStatusRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.ShadowingStatus, error) {
	var status models.ShadowingStatus
	err := r.db.WithContext(ctx).Where("user_id = ? AND transcript_id = ?", userID, transcriptID).First(&status).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *shadowingStatusRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.ShadowingStatus], error) {
	var statuses []*models.ShadowingStatus

	base := r.db.WithContext(ctx).Model(&models.ShadowingStatus{})
	var total int64
	query.Count(base).Count(&total)

	err := query.Apply(base).Find(&statuses).Error
	if err != nil {
		return nil, err
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.ShadowingStatus]{
		Data: statuses,
		Meta: meta,
	}, nil
}
