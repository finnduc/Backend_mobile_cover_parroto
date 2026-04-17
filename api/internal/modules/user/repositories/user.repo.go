package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindByID(ctx context.Context, id uint) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &user, nil
}
