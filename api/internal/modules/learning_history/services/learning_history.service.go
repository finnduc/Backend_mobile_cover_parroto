package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	lhreq "go-cover-parroto/internal/modules/learning_history/dtos/req"
	lhres "go-cover-parroto/internal/modules/learning_history/dtos/res"
	"go-cover-parroto/internal/modules/learning_history/repositories"
	"go-cover-parroto/internal/utils"
)

type ILearningHistoryService interface {
	Record(ctx context.Context, body lhreq.RecordHistoryReq) (*lhres.LearningHistoryRes, *response.AppError)
	List(ctx context.Context, query lhreq.ListHistoryQuery) (*response.PaginatedResponse[lhres.LearningHistoryRes], *response.AppError)
	GetByLesson(ctx context.Context, body lhreq.GetHistoryReq) (*lhres.LearningHistoryRes, *response.AppError)
}

type learningHistoryService struct {
	repo repositories.ILearningHistoryRepo
}

func NewLearningHistoryService(repo repositories.ILearningHistoryRepo) ILearningHistoryService {
	return &learningHistoryService{repo: repo}
}

func (s *learningHistoryService) Record(ctx context.Context, body lhreq.RecordHistoryReq) (*lhres.LearningHistoryRes, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	history := &models.LearningHistory{
		UserID:          userID,
		LessonID:        body.LessonID,
		DurationWatched: body.DurationWatched,
		Completed:       body.Completed,
	}
	if upsertErr := s.repo.Upsert(ctx, history); upsertErr != nil {
		return nil, response.Internal("failed to record history")
	}
	var result lhres.LearningHistoryRes
	if err := utils.MapToDTO(history, &result); err != nil {
		return nil, response.Internal("failed to map history")
	}
	return &result, nil
}

func (s *learningHistoryService) List(ctx context.Context, query lhreq.ListHistoryQuery) (*response.PaginatedResponse[lhres.LearningHistoryRes], *response.AppError) {
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		return nil, response.Internal("failed to list history")
	}
	var histories []lhres.LearningHistoryRes
	if err := utils.MapToDTOs(result.Data, &histories); err != nil {
		return nil, response.Internal("failed to map history")
	}
	return &response.PaginatedResponse[lhres.LearningHistoryRes]{Data: histories, Meta: result.Meta}, nil
}

func (s *learningHistoryService) GetByLesson(ctx context.Context, body lhreq.GetHistoryReq) (*lhres.LearningHistoryRes, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	history, findErr := s.repo.FindByUserAndLesson(ctx, userID, body.LessonID)
	if findErr != nil {
		if errors.Is(findErr, coreError.ErrNotFound) {
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
