package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/lesson/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "lesson")
}

type ILessonService interface {
	List(ctx context.Context, query req.ListLessonQuery) (*response.PaginatedResult[*models.Lesson], *response.AppError)
	Get(ctx context.Context, body req.GetLessonReq) (*models.Lesson, *response.AppError)
	Create(ctx context.Context, body req.CreateLessonReq) (*models.Lesson, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateLessonReq) (*models.Lesson, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
}

type lessonService struct {
	repo db_repos.ILessonRepo
}

func NewLessonService(repo db_repos.ILessonRepo) ILessonService {
	return &lessonService{repo: repo}
}

func (s *lessonService) List(ctx context.Context, query req.ListLessonQuery) (*response.PaginatedResult[*models.Lesson], *response.AppError) {
	log := sLog()
	log.Infow("listing lessons")
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list lessons", "error", err)
		return nil, response.Internal("failed to list lessons")
	}
	return result, nil
}

func (s *lessonService) Get(ctx context.Context, body req.GetLessonReq) (*models.Lesson, *response.AppError) {
	log := sLog().With("lessonId", body.ID)
	log.Infow("getting lesson")
	lesson, err := s.repo.FindByID(ctx, body.ID)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("lesson not found")
		}
		log.Errorw("failed to get lesson", "error", err)
		return nil, response.Internal("failed to get lesson")
	}
	return lesson, nil
}

func (s *lessonService) Create(ctx context.Context, body req.CreateLessonReq) (*models.Lesson, *response.AppError) {
	log := sLog()
	log.Infow("creating lesson")
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
		log.Errorw("failed to create lesson", "error", err)
		return nil, response.Internal("failed to create lesson")
	}
	log.Infow("lesson created", "lessonId", lesson.ID)
	return lesson, nil
}

func (s *lessonService) Update(ctx context.Context, id uint, body req.UpdateLessonReq) (*models.Lesson, *response.AppError) {
	log := sLog().With("lessonId", id)
	log.Infow("updating lesson")
	lesson, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("lesson not found")
		}
		log.Errorw("failed to get lesson for update", "error", err)
		return nil, response.Internal("failed to get lesson")
	}
	if body.CategoryID != nil {
		lesson.CategoryID = body.CategoryID
	}
	if body.Title != nil {
		lesson.Title = *body.Title
	}
	if body.Description != nil {
		lesson.Description = *body.Description
	}
	if body.VideoURL != nil {
		lesson.VideoURL = *body.VideoURL
	}
	if body.ThumbnailURL != nil {
		lesson.ThumbnailURL = *body.ThumbnailURL
	}
	if body.Level != nil {
		lesson.Level = *body.Level
	}
	if body.Duration != nil {
		lesson.Duration = *body.Duration
	}
	if updateErr := s.repo.Update(ctx, lesson); updateErr != nil {
		log.Errorw("failed to update lesson", "error", updateErr)
		return nil, response.Internal("failed to update lesson")
	}
	log.Infow("lesson updated")
	return lesson, nil
}

func (s *lessonService) Delete(ctx context.Context, id uint) *response.AppError {
	log := sLog().With("lessonId", id)
	log.Infow("deleting lesson")
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return response.NotFound("lesson not found")
		}
		log.Errorw("failed to get lesson for delete", "error", err)
		return response.Internal("failed to get lesson")
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Errorw("failed to delete lesson", "error", err)
		return response.Internal("failed to delete lesson")
	}
	log.Infow("lesson deleted")
	return nil
}
