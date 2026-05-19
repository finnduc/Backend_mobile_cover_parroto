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

type vocabularyCategoryRepo struct {
	db *gorm.DB
}

func NewVocabularyCategoryRepo(db *gorm.DB) db_repos.IVocabularyCategoryRepo {
	return &vocabularyCategoryRepo{db: db}
}

func (r *vocabularyCategoryRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyCategory], error) {
	var categories []*models.VocabularyCategory

	countQuery := query.Count(r.db.Model(&models.VocabularyCategory{}))
	var total int64
	countQuery.Count(&total)

	result := query.Apply(r.db.WithContext(ctx).Model(&models.VocabularyCategory{}))
	err := result.Find(&categories).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.VocabularyCategory]{
		Data: categories,
		Meta: meta,
	}, nil
}

func (r *vocabularyCategoryRepo) FindByID(ctx context.Context, id uint) (*models.VocabularyCategory, error) {
	var category models.VocabularyCategory
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &category, nil
}

func (r *vocabularyCategoryRepo) Create(ctx context.Context, category *models.VocabularyCategory) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *vocabularyCategoryRepo) Update(ctx context.Context, category *models.VocabularyCategory) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *vocabularyCategoryRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VocabularyCategory{}, id).Error
}
