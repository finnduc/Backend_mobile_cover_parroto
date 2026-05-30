package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepo interface {
	FindByID(ctx context.Context, id string) (*models.User, error)
	Upsert(ctx context.Context, user *models.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &user, nil
}

func (r *userRepo) Upsert(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"email", "name", "avatar_url",
		}),
	}).Create(user).Error
}
