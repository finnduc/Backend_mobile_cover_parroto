package repositories

import (
	"context"

	"go-cover-parroto/internal/core/errors"
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

func (r *learningHistoryRepo) Create(ctx context.Context, history *models.LearningHistory) error {
	return errors.MapRepoError(r.db.WithContext(ctx).Create(history).Error)
}

func (r *learningHistoryRepo) Update(ctx context.Context, history *models.LearningHistory) error {
	return errors.MapRepoError(r.db.WithContext(ctx).Save(history).Error)
}

func (r *learningHistoryRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error) {
	var history models.LearningHistory
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		First(&history).Error
	if err != nil {
		return nil, errors.MapRepoError(err)
	}
	return &history, nil
}

func (r *learningHistoryRepo) FindAllByUser(ctx context.Context, userID string, filter string) ([]*models.LearningHistory, error) {
	var histories []*models.LearningHistory
	query := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("updated_at DESC")
	switch filter {
	case "finished":
		query = query.Where("completed_dictation = ? AND completed_pronunciation = ?", true, true)
	case "unfinished":
		query = query.Where("NOT (completed_dictation = ? AND completed_pronunciation = ?)", true, true)
	}
	if err := query.Find(&histories).Error; err != nil {
		return nil, errors.MapRepoError(err)
	}
	return histories, nil
}

func (r *learningHistoryRepo) CountSummary(ctx context.Context, userID string) (int64, int64, error) {
	var completed int64
	if err := r.db.WithContext(ctx).Model(&models.LearningHistory{}).
		Where("user_id = ? AND completed_dictation = ? AND completed_pronunciation = ?", userID, true, true).
		Count(&completed).Error; err != nil {
		return 0, 0, errors.MapRepoError(err)
	}

	var unfinished int64
	if err := r.db.WithContext(ctx).Model(&models.LearningHistory{}).
		Where("user_id = ? AND NOT (completed_dictation = ? AND completed_pronunciation = ?)", userID, true, true).
		Count(&unfinished).Error; err != nil {
		return 0, 0, errors.MapRepoError(err)
	}

	return completed, unfinished, nil
}

func (r *learningHistoryRepo) CountCompletedLessonTranscripts(ctx context.Context, userID string, lessonID uint) (int64, int64, error) {
	var total int64
	if err := r.db.WithContext(ctx).Model(&models.Transcript{}).
		Where("lesson_id = ?", lessonID).
		Count(&total).Error; err != nil {
		return 0, 0, errors.MapRepoError(err)
	}

	var completed int64
	err := r.db.WithContext(ctx).Raw(`
		SELECT COUNT(DISTINCT transcript_id)
		FROM (
			SELECT transcript_id FROM dictation_status WHERE user_id = ? AND lesson_id = ?
			UNION
			SELECT transcript_id FROM shadowing_status WHERE user_id = ? AND lesson_id = ?
		) completed_transcripts
	`, userID, lessonID, userID, lessonID).Scan(&completed).Error
	if err != nil {
		return 0, 0, errors.MapRepoError(err)
	}

	return completed, total, nil
}
