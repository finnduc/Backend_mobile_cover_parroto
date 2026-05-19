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

type vocabularyItemRepo struct {
	db *gorm.DB
}

func NewVocabularyItemRepo(db *gorm.DB) db_repos.IVocabularyItemRepo {
	return &vocabularyItemRepo{db: db}
}

func (r *vocabularyItemRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyItem], error) {
	var items []*models.VocabularyItem

	countQuery := query.Count(r.db.Model(&models.VocabularyItem{}))
	var total int64
	countQuery.Count(&total)

	result := query.Apply(r.db.WithContext(ctx).Model(&models.VocabularyItem{}))
	err := result.Preload("Deck").Find(&items).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.VocabularyItem]{
		Data: items,
		Meta: meta,
	}, nil
}

func (r *vocabularyItemRepo) FindByID(ctx context.Context, id uint) (*models.VocabularyItem, error) {
	var item models.VocabularyItem
	err := r.db.WithContext(ctx).Preload("Deck").First(&item, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &item, nil
}

func (r *vocabularyItemRepo) Create(ctx context.Context, item *models.VocabularyItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *vocabularyItemRepo) Update(ctx context.Context, item *models.VocabularyItem) error {
	return r.db.WithContext(ctx).Save(item).Error
}

func (r *vocabularyItemRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VocabularyItem{}, id).Error
}
