package services

import (
	"context"
	"errors"
	"time"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/lesson/dtos/res"
	"go-cover-parroto/internal/modules/lesson/repositories"
)

type ILessonService interface {
	ListLessons(ctx context.Context) ([]res.LessonRes, *response.AppError)
	GetLesson(ctx context.Context, id uint) (*res.LessonRes, *response.AppError)
}

type lessonService struct {
	repo repositories.ILessonRepo
}

func NewLessonService(repo repositories.ILessonRepo) ILessonService {
	return &lessonService{repo: repo}
}

func (s *lessonService) ListLessons(ctx context.Context) ([]res.LessonRes, *response.AppError) {
	lessons, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, response.Internal("failed to list lessons")
	}
	var result []res.LessonRes
	for _, lesson := range lessons {
		result = append(result, toLessonRes(lesson))
	}
	return result, nil
}

func (s *lessonService) GetLesson(ctx context.Context, id uint) (*res.LessonRes, *response.AppError) {
	lesson, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, response.NotFound("lesson not found")
		}
		return nil, response.Internal("failed to get lesson")
	}
	res := toLessonRes(*lesson)
	return &res, nil
}

func toLessonRes(lesson models.Lesson) res.LessonRes {
	return res.LessonRes{
		ID:           lesson.ID,
		CategoryID:   lesson.CategoryID,
		Title:        lesson.Title,
		Description:  lesson.Description,
		VideoURL:     lesson.VideoURL,
		ThumbnailURL: lesson.ThumbnailURL,
		Level:        lesson.Level,
		Duration:     lesson.Duration,
		CreatedAt:    lesson.CreatedAt.Format(time.RFC3339),
	}
}
