package services

import (
	"context"
	"errors"

	"go-cover-parroto/internal/core/database"
	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	lhreq "go-cover-parroto/internal/modules/learning_history/dtos/req"
	lhres "go-cover-parroto/internal/modules/learning_history/dtos/res"
	"go-cover-parroto/internal/modules/learning_history/repositories"
	"go-cover-parroto/internal/utils"
)

type ILearningHistoryService interface {
	Record(ctx context.Context, userID uint, body lhreq.RecordHistoryReq) (*lhres.LearningHistoryRes, *response.AppError)
	ListByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResponse[lhres.LearningHistoryRes], *response.AppError)
	GetByLesson(ctx context.Context, userID, lessonID uint) (*lhres.LearningHistoryRes, *response.AppError)
}

type learningHistoryService struct {
	repo repositories.ILearningHistoryRepo
}

func NewLearningHistoryService(repo repositories.ILearningHistoryRepo) ILearningHistoryService {
	return &learningHistoryService{repo: repo}
}

func (s *learningHistoryService) Record(ctx context.Context, userID uint, body lhreq.RecordHistoryReq) (*lhres.LearningHistoryRes, *response.AppError) {
	history := &models.LearningHistory{
		UserID:          userID,
		LessonID:        body.LessonID,
		DurationWatched: body.DurationWatched,
		Completed:       body.Completed,
	}

	if err := s.repo.Create(ctx, history); err != nil {
		return nil, response.Internal("failed to record history")
	}

	var result lhres.LearningHistoryRes
	if err := utils.MapToDTO(history, &result); err != nil {
		return nil, response.Internal("failed to map history")
	}
	return &result, nil
}

func (s *learningHistoryService) ListByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResponse[lhres.LearningHistoryRes], *response.AppError) {
	result, err := s.repo.FindAllByUser(ctx, userID, query)
	if err != nil {
		return nil, response.Internal("failed to list history")
	}

	var histories []lhres.LearningHistoryRes
	if err := utils.MapToDTOs(result.Data, &histories); err != nil {
		return nil, response.Internal("failed to map history")
	}

	return &response.PaginatedResponse[lhres.LearningHistoryRes]{
		Data: histories,
		Meta: result.Meta,
	}, nil
}

func (s *learningHistoryService) GetByLesson(ctx context.Context, userID, lessonID uint) (*lhres.LearningHistoryRes, *response.AppError) {
	history, err := s.repo.FindByUserAndLesson(ctx, userID, lessonID)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("history not found")
		}
		return nil, response.Internal("failed to get history")
	}

	var result lhres.LearningHistoryRes
	if err := utils.MapToDTO(history, &result); err != nil {
		return nil, response.Internal("failed to map history")
	}
	return &result, nil
}
