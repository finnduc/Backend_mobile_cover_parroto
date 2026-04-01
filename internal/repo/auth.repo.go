
package repo

import (
	"context"
	"go-familytree/internal/models"

	"gorm.io/gorm"
)

type authRepo struct{ db *gorm.DB }

func NewAuthRepo(db *gorm.DB) IAuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *authRepo) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *authRepo) GetUserRoles(ctx context.Context, userID uint) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.WithContext(ctx).
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}
