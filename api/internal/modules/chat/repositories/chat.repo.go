package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
)

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) db_repos.IChatRepo {
	return &chatRepo{db: db}
}

func (r *chatRepo) Create(ctx context.Context, msg *models.GlobalChatMessage) error {
	if err := r.db.WithContext(ctx).Create(msg).Error; err != nil {
		return errors.MapRepoError(err)
	}
	return nil
}

func (r *chatRepo) FindHistory(ctx context.Context, beforeID uint64, limit int) ([]*models.GlobalChatMessage, error) {
	var messages []*models.GlobalChatMessage

	tx := r.db.WithContext(ctx).
		Model(&models.GlobalChatMessage{}).
		Preload("User").
		Order("id DESC").
		Limit(limit)

	if beforeID > 0 {
		tx = tx.Where("id < ?", beforeID)
	}

	if err := tx.Find(&messages).Error; err != nil {
		return nil, errors.MapRepoError(err)
	}
	return messages, nil
}
