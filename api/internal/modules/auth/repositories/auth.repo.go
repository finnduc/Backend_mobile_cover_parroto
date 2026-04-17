package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type IAuthRepo interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) IAuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *authRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &user, nil
}
