package repositories

import (
	"context"
	"errors"

	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("record not found")

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
