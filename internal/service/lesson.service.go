package service

import (
	"context"
	"fmt"
	"go-familytree/internal/models"
	"go-familytree/internal/repo"
	pkgerrors "go-familytree/pkg/errors"
	"go-familytree/pkg/utils"

	"gorm.io/gorm"
)

type lessonService struct {
	lessonRepo repo.ILessonRepo
}

func NewLessonService(lessonRepo repo.ILessonRepo) ILessonService {
	return &lessonService{lessonRepo: lessonRepo}
}

func (s *lessonService) ListLessons(ctx context.Context, level string, categoryID *uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	lessons, total, err := s.lessonRepo.FindAll(ctx, repo.LessonFilter{
		Level:           level,
		CategoryID:      categoryID,
		PaginationQuery: q,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("lessonService.ListLessons: %w", pkgerrors.ErrInternalServer)
	}
	return lessons, total, nil
}

func (s *lessonService) CreateLesson(ctx context.Context, input CreateLessonInput) (*models.Lesson, error) {
	videoURL := input.URL
	if videoURL == "" {
		videoURL = input.VideoURL
	}

	lesson := &models.Lesson{
		Title:        input.Title,
		Description:  input.Description,
		VideoURL:     videoURL,
		ThumbnailURL: input.ThumbnailURL,
		Level:        input.Level,
		Duration:     input.Duration,
	}

	transcripts := make([]models.Transcript, 0, len(input.Transcripts))
	for _, t := range input.Transcripts {
		transcripts = append(transcripts, models.Transcript{
			Sequence:       t.Sequence,
			Content:        t.Content,
			Phonetic:       t.Phonetic,
			Vietnamese:     t.Vietnamese,
			StartTimestamp: t.StartTimestamp,
			EndTimestamp:   t.EndTimestamp,
		})
	}

	if err := s.lessonRepo.Create(ctx, lesson, input.CategoryIDs, transcripts); err != nil {
		return nil, fmt.Errorf("lessonService.CreateLesson: %w", pkgerrors.ErrInternalServer)
	}
	return lesson, nil
}

func (s *lessonService) GetLesson(ctx context.Context, id uint) (*models.Lesson, error) {
	lesson, err := s.lessonRepo.FindByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("lessonService.GetLesson: %w", pkgerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("lessonService.GetLesson: %w", pkgerrors.ErrInternalServer)
	}
	return lesson, nil
}

func (s *lessonService) GetTranscripts(ctx context.Context, lessonID uint) ([]models.Transcript, error) {
	// Verify lesson exists first
	if _, err := s.lessonRepo.FindByID(ctx, lessonID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("lessonService.GetTranscripts: %w", pkgerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("lessonService.GetTranscripts: %w", pkgerrors.ErrInternalServer)
	}
	transcripts, err := s.lessonRepo.FindTranscriptsByLessonID(ctx, lessonID)
	if err != nil {
		return nil, fmt.Errorf("lessonService.GetTranscripts: %w", pkgerrors.ErrInternalServer)
	}
	return transcripts, nil
}
