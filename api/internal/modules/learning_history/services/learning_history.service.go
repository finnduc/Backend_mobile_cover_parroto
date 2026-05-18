package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	lhreq "go-cover-parroto/internal/modules/learning_history/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "learning_history")
}

type ILearningHistoryService interface {
	Record(ctx context.Context, body lhreq.RecordHistoryReq) (*models.LearningHistory, *response.AppError)
	List(ctx context.Context, query lhreq.ListHistoryQuery) (*response.PaginatedResult[*models.LearningHistory], *response.AppError)
	GetByLesson(ctx context.Context, body lhreq.GetHistoryReq) (*models.LearningHistory, *response.AppError)
}

type learningHistoryService struct {
	repo db_repos.ILearningHistoryRepo
}

func NewLearningHistoryService(repo db_repos.ILearningHistoryRepo) ILearningHistoryService {
	return &learningHistoryService{repo: repo}
}

func (s *learningHistoryService) Record(ctx context.Context, body lhreq.RecordHistoryReq) (*models.LearningHistory, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "lessonId", body.LessonID)
	log.Infow("recording learning history")
	history := &models.LearningHistory{
		UserID:          userID,
		LessonID:        body.LessonID,
		DurationWatched: body.DurationWatched,
		Completed:       body.Completed,
	}
	if upsertErr := s.repo.Upsert(ctx, history); upsertErr != nil {
		log.Errorw("failed to record history", "error", upsertErr)
		return nil, response.Internal("failed to record history")
	}
	log.Infow("learning history recorded")
	return history, nil
}

func (s *learningHistoryService) List(ctx context.Context, query lhreq.ListHistoryQuery) (*response.PaginatedResult[*models.LearningHistory], *response.AppError) {
	log := sLog()
	log.Infow("listing learning history")
	if query.UserID != nil {
		if err := policy.Allow(ctx, *query.UserID); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list history", "error", err)
		return nil, response.Internal("failed to list history")
	}
	return result, nil
}

func (s *learningHistoryService) GetByLesson(ctx context.Context, body lhreq.GetHistoryReq) (*models.LearningHistory, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "lessonId", body.LessonID)
	log.Infow("getting learning history by lesson")
	history, findErr := s.repo.FindByUserAndLesson(ctx, userID, body.LessonID)
	if findErr != nil {
		if errors.Is(findErr, coreError.ErrNotFound) {
			return nil, response.NotFound("history not found")
		}
		log.Errorw("failed to get history", "error", findErr)
		return nil, response.Internal("failed to get history")
	}
	return history, nil
}
