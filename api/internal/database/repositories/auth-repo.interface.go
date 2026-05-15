package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
)

type IAuthRepo interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}
