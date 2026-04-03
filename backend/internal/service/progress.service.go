package service

import (
	"context"
	"fmt"
	"go-cover-parroto/internal/models"
	"go-cover-parroto/internal/repo"
	pkgerrors "go-cover-parroto/pkg/errors"
)

type progressService struct {
	progressRepo repo.IProgressRepo
}

func NewProgressService(progressRepo repo.IProgressRepo) IProgressService {
	return &progressService{progressRepo: progressRepo}
}

func (s *progressService) GetProgress(ctx context.Context, userID, lessonID uint) (*models.UserProgress, error) {
	p, err := s.progressRepo.FindOrCreate(ctx, userID, lessonID)
	if err != nil {
		return nil, fmt.Errorf("progressService.GetProgress: %w", pkgerrors.ErrInternalServer)
	}
	return p, nil
}
