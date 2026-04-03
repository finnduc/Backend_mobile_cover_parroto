package repo

import (
	"context"
	"go-cover-parroto/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type progressRepo struct{ db *gorm.DB }

func NewProgressRepo(db *gorm.DB) IProgressRepo {
	return &progressRepo{db: db}
}

func (r *progressRepo) FindOrCreate(ctx context.Context, userID, lessonID uint) (*models.UserProgress, error) {
	var p models.UserProgress
	err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		FirstOrCreate(&p, models.UserProgress{UserID: userID, LessonID: lessonID}).Error
	return &p, err
}

func (r *progressRepo) Save(ctx context.Context, p *models.UserProgress) error {
	return r.db.WithContext(ctx).Save(p).Error
}
