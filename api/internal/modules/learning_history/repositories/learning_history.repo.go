package repositories

import (
	"context"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"gorm.io/gorm"
)

type ILearningHistoryRepo interface {
	Create(ctx context.Context, history *models.LearningHistory) error
	FindAllByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResult[*models.LearningHistory], error)
	FindByUserAndLesson(ctx context.Context, userID, lessonID uint) (*models.LearningHistory, error)
}

type learningHistoryRepo struct {
	db *gorm.DB
}

func NewLearningHistoryRepo(db *gorm.DB) ILearningHistoryRepo {
	return &learningHistoryRepo{db: db}
}

func (r *learningHistoryRepo) Create(ctx context.Context, history *models.LearningHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

func (r *learningHistoryRepo) FindAllByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResult[*models.LearningHistory], error) {
	var histories []*models.LearningHistory

	base := r.db.WithContext(ctx).Model(&models.LearningHistory{}).Where("user_id = ?", userID)

	var total int64
	base.Count(&total)

	err := query.Apply(base).Find(&histories).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}

	meta := response.NewMeta(query.Page, query.Limit, total)
	return &response.PaginatedResult[*models.LearningHistory]{
		Data: histories,
		Meta: meta,
	}, nil
}

func (r *learningHistoryRepo) FindByUserAndLesson(ctx context.Context, userID, lessonID uint) (*models.LearningHistory, error) {
	var history models.LearningHistory
	err := r.db.WithContext(ctx).Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&history).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &history, nil
}
