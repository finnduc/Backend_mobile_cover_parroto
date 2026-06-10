package services

import (
	"context"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/learning_history/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "learning_history")
}

type ILearningHistoryService interface {
	CreateOrUpdate(ctx context.Context, userID string, body req.CreateLearningHistoryReq) (*models.LearningHistory, *response.AppError)
	List(ctx context.Context, userID string) ([]*models.LearningHistory, *response.AppError)
	ListFinished(ctx context.Context, userID string) ([]*models.LearningHistory, *response.AppError)
	ListUnfinished(ctx context.Context, userID string) ([]*models.LearningHistory, *response.AppError)
	Summary(ctx context.Context, userID string) (int64, int64, *response.AppError)
	LessonSummary(ctx context.Context, userID string, lessonID uint) (int64, int64, *response.AppError)
}

type learningHistoryService struct {
	repo db_repos.ILearningHistoryRepo
}

func NewLearningHistoryService(repo db_repos.ILearningHistoryRepo) ILearningHistoryService {
	return &learningHistoryService{repo: repo}
}

func (s *learningHistoryService) CreateOrUpdate(ctx context.Context, userID string, body req.CreateLearningHistoryReq) (*models.LearningHistory, *response.AppError) {

	history := &models.LearningHistory{
		UserID:                 userID,
		LessonID:               body.LessonID,
		CompletedDictation:     body.CompletedDictation,
		CompletedPronunciation: body.CompletedPronunciation,
	}

	if upsertErr := s.repo.Upsert(ctx, history); upsertErr != nil {
		return nil, response.Internal("failed to save learning history")
	}

	saved, findErr := s.repo.FindByUserAndLesson(ctx, userID, body.LessonID)
	if findErr != nil {
		return nil, response.Internal("failed to retrieve learning history")
	}

	return saved, nil
}

func (s *learningHistoryService) List(ctx context.Context, userID string) ([]*models.LearningHistory, *response.AppError) {
	result, findErr := s.repo.FindByUser(ctx, userID)
	if findErr != nil {
		return nil, response.Internal("failed to list")
	}
	return result, nil
}

func (s *learningHistoryService) ListFinished(ctx context.Context, userID string) ([]*models.LearningHistory, *response.AppError) {
	result, findErr := s.repo.FindFinished(ctx, userID)
	if findErr != nil {
		return nil, response.Internal("failed to list finished")
	}
	return result, nil
}

func (s *learningHistoryService) ListUnfinished(ctx context.Context, userID string) ([]*models.LearningHistory, *response.AppError) {
	result, findErr := s.repo.FindUnfinished(ctx, userID)
	if findErr != nil {
		return nil, response.Internal("failed to list unfinished")
	}
	return result, nil
}

func (s *learningHistoryService) Summary(ctx context.Context, userID string) (int64, int64, *response.AppError) {
	finished, err1 := s.repo.CountFinished(ctx, userID)
	unfinished, err2 := s.repo.CountUnfinished(ctx, userID)
	if err1 != nil || err2 != nil {
		return 0, 0, response.Internal("failed to get summary")
	}
	return finished, unfinished, nil
}

func (s *learningHistoryService) LessonSummary(ctx context.Context, userID string, lessonID uint) (int64, int64, *response.AppError) {

	history, findErr := s.repo.FindByUserAndLesson(ctx, userID, lessonID)
	if findErr != nil {
		if findErr.Error() == coreError.ErrNotFound.Error() {
			return 0, 0, nil
		}
		return 0, 0, response.Internal("failed to get lesson summary")
	}

	if history.CompletedDictation && history.CompletedPronunciation != nil && *history.CompletedPronunciation {
		return 1, 0, nil
	}
	return 0, 1, nil
}
