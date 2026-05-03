package repositories

import (
	"context"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.User], error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &user, nil
}

func (r *userRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.User], error) {
	var users []*models.User
	base := r.db.WithContext(ctx).Model(&models.User{})

	var total int64
	query.Count(base).Count(&total)

	err := query.Apply(base).Find(&users).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.User]{
		Data: users,
		Meta: meta,
	}, nil
}

func (r *userRepo) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}
