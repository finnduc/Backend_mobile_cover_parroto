package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/lesson/dtos/req"
	"go-cover-parroto/internal/modules/lesson/dtos/res"
	"go-cover-parroto/internal/modules/lesson/repositories"
	"go-cover-parroto/internal/utils"
)

type ILessonService interface {
	List(ctx context.Context, query req.ListLessonQuery) (*response.PaginatedResponse[res.LessonRes], *response.AppError)
	Get(ctx context.Context, body req.GetLessonReq) (*res.LessonRes, *response.AppError)
	Create(ctx context.Context, body req.CreateLessonReq) (*res.LessonRes, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateLessonReq) (*res.LessonRes, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
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

func (s *lessonService) Create(ctx context.Context, body req.CreateLessonReq) (*res.LessonRes, *response.AppError) {
	lesson := &models.Lesson{
		CategoryID:   body.CategoryID,
		Title:        body.Title,
		Description:  body.Description,
		VideoURL:     body.VideoURL,
		ThumbnailURL: body.ThumbnailURL,
		Level:        body.Level,
		Duration:     body.Duration,
	}
	if err := s.repo.Create(ctx, lesson); err != nil {
		return nil, response.Internal("failed to create lesson")
	}
	var result res.LessonRes
	if err := utils.MapToDTO(lesson, &result); err != nil {
		return nil, response.Internal("failed to map lesson")
	}
	return &result, nil
}

func (s *lessonService) Update(ctx context.Context, id uint, body req.UpdateLessonReq) (*res.LessonRes, *response.AppError) {
	lesson, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("lesson not found")
		}
		return nil, response.Internal("failed to get lesson")
	}
	lesson.CategoryID = body.CategoryID
	lesson.Title = body.Title
	lesson.Description = body.Description
	lesson.VideoURL = body.VideoURL
	lesson.ThumbnailURL = body.ThumbnailURL
	lesson.Level = body.Level
	lesson.Duration = body.Duration
	if updateErr := s.repo.Update(ctx, lesson); updateErr != nil {
		return nil, response.Internal("failed to update lesson")
	}
	var result res.LessonRes
	if err := utils.MapToDTO(lesson, &result); err != nil {
		return nil, response.Internal("failed to map lesson")
	}
	return &result, nil
}

func (s *lessonService) Delete(ctx context.Context, id uint) *response.AppError {
	if err := s.repo.Delete(ctx, id); err != nil {
		return response.Internal("failed to delete lesson")
	}
	return nil
}
