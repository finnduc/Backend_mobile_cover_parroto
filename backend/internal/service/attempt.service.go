package service

import (
	"context"
	"fmt"
	"go-cover-parroto/internal/models"
	"go-cover-parroto/internal/repo"
	pkgerrors "go-cover-parroto/pkg/errors"

	"gorm.io/gorm"
)

type attemptService struct {
	attemptRepo repo.IAttemptRepo
	lessonRepo  repo.ILessonRepo
}

func NewAttemptService(attemptRepo repo.IAttemptRepo, lessonRepo repo.ILessonRepo) IAttemptService {
	return &attemptService{attemptRepo: attemptRepo, lessonRepo: lessonRepo}
}

func (s *attemptService) CreateAttempt(ctx context.Context, userID, lessonID uint) (*models.Attempt, error) {
	// Verify lesson exists
	if _, err := s.lessonRepo.FindByID(ctx, lessonID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("attemptService.CreateAttempt: %w", pkgerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("attemptService.CreateAttempt: %w", pkgerrors.ErrInternalServer)
	}

	attempt := &models.Attempt{
		UserID:   userID,
		LessonID: lessonID,
	}
	if err := s.attemptRepo.Create(ctx, attempt); err != nil {
		return nil, fmt.Errorf("attemptService.CreateAttempt: %w", pkgerrors.ErrInternalServer)
	}
	return attempt, nil
}

func (s *attemptService) GetAttempt(ctx context.Context, id uint) (*models.Attempt, error) {
	attempt, err := s.attemptRepo.FindByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("attemptService.GetAttempt: %w", pkgerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("attemptService.GetAttempt: %w", pkgerrors.ErrInternalServer)
	}
	return attempt, nil
}
