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

type vocabularyDeckRepo struct {
	db *gorm.DB
}

func NewVocabularyDeckRepo(db *gorm.DB) db_repos.IVocabularyDeckRepo {
	return &vocabularyDeckRepo{db: db}
}

func (r *vocabularyDeckRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyDeck], error) {
	var decks []*models.VocabularyDeck

	countQuery := query.Count(r.db.Model(&models.VocabularyDeck{}))
	var total int64
	countQuery.Count(&total)

	result := query.Apply(r.db.WithContext(ctx).Model(&models.VocabularyDeck{}))
	err := result.Preload("Category").Find(&decks).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.VocabularyDeck]{
		Data: decks,
		Meta: meta,
	}, nil
}

func (r *vocabularyDeckRepo) FindByID(ctx context.Context, id uint) (*models.VocabularyDeck, error) {
	var deck models.VocabularyDeck
	err := r.db.WithContext(ctx).Preload("Category").First(&deck, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &deck, nil
}

func (r *vocabularyDeckRepo) Create(ctx context.Context, deck *models.VocabularyDeck) error {
	return r.db.WithContext(ctx).Create(deck).Error
}

func (r *vocabularyDeckRepo) Update(ctx context.Context, deck *models.VocabularyDeck) error {
	return r.db.WithContext(ctx).Save(deck).Error
}

func (r *vocabularyDeckRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.VocabularyDeck{}, id).Error
}
