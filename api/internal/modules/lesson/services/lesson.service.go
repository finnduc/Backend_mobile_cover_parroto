package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/lesson/dtos/req"
	"go-cover-parroto/internal/modules/lesson/dtos/res"
	"go-cover-parroto/internal/modules/lesson/repositories"
	"go-cover-parroto/internal/utils"
)

type ILessonService interface {
	List(ctx context.Context, query req.ListLessonQuery) (*response.PaginatedResponse[res.LessonRes], *response.AppError)
	Get(ctx context.Context, body req.GetLessonReq) (*res.LessonRes, *response.AppError)
}

type lessonService struct {
	repo repositories.ILessonRepo
}

func NewLessonService(repo repositories.ILessonRepo) ILessonService {
	return &lessonService{repo: repo}
}

func (s *lessonService) List(ctx context.Context, query req.ListLessonQuery) (*response.PaginatedResponse[res.LessonRes], *response.AppError) {
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		return nil, response.Internal("failed to list lessons")
	}
	var lessons []res.LessonRes
	if err := utils.MapToDTOs(result.Data, &lessons); err != nil {
		return nil, response.Internal("failed to map lessons")
	}
	return &response.PaginatedResponse[res.LessonRes]{Data: lessons, Meta: result.Meta}, nil
}

func (s *lessonService) Get(ctx context.Context, body req.GetLessonReq) (*res.LessonRes, *response.AppError) {
	lesson, err := s.repo.FindByID(ctx, body.ID)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("lesson not found")
		}
		return nil, response.Internal("failed to get lesson")
	}
	var result res.LessonRes
	if err := utils.MapToDTO(lesson, &result); err != nil {
		return nil, response.Internal("failed to map lesson")
	}
	return &result, nil
}
