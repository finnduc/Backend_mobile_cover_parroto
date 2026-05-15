package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"gorm.io/gorm"
)

type learningHistoryRepo struct {
	db *gorm.DB
}

func NewLearningHistoryRepo(db *gorm.DB) db_repos.ILearningHistoryRepo {
	return &learningHistoryRepo{db: db}
}

func (r *learningHistoryRepo) Upsert(ctx context.Context, history *models.LearningHistory) error {
	var existing models.LearningHistory
	err := r.db.WithContext(ctx).Where("user_id = ? AND lesson_id = ?", history.UserID, history.LessonID).First(&existing).Error
	if err != nil {
		if errors.MapRepoError(err) == errors.ErrNotFound {
			return r.db.WithContext(ctx).Create(history).Error
		}
		return err
	}

	existing.DurationWatched = history.DurationWatched
	existing.Completed = history.Completed
	if err := r.db.WithContext(ctx).Save(&existing).Error; err != nil {
		return err
	}
	*history = existing
	return nil
}

func (r *learningHistoryRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.LearningHistory], error) {
	var histories []*models.LearningHistory

	base := r.db.WithContext(ctx).Model(&models.LearningHistory{})

	var total int64
	query.Count(base).Count(&total)

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

func (r *learningHistoryRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error) {
	var history models.LearningHistory
	err := r.db.WithContext(ctx).Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&history).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &history, nil
}
