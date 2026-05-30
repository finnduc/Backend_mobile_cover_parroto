package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
)

type IChatRepo interface {
	Create(ctx context.Context, msg *models.GlobalChatMessage) error
	FindHistory(ctx context.Context, beforeID uint64, limit int) ([]*models.GlobalChatMessage, error)
}
