package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dictationStatusRepo struct {
	db *gorm.DB
}

func NewDictationStatusRepo(db *gorm.DB) db_repos.IDictationStatusRepo {
	return &dictationStatusRepo{db: db}
}

func (r *dictationStatusRepo) CreateOrIgnore(ctx context.Context, status *models.DictationStatus) error {
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(status).Error
}

func (r *dictationStatusRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.DictationStatus], error) {
	var statuses []*models.DictationStatus

	base := r.db.WithContext(ctx).Model(&models.DictationStatus{})
	var total int64
	query.Count(base).Count(&total)

	err := query.Apply(base).Find(&statuses).Error
	if err != nil {
		return nil, err
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.DictationStatus]{
		Data: statuses,
		Meta: meta,
	}, nil
}
