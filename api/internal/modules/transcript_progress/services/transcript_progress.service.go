package services

import (
	"context"
	"time"

	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/transcript_progress/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "transcript_progress")
}

type ITranscriptProgressService interface {
	Create(ctx context.Context, lessonID uint, body req.CreateTranscriptProgressReq) (*models.TranscriptProgress, *response.AppError)
	List(ctx context.Context, lessonID uint) ([]*models.TranscriptProgress, *response.AppError)
}

type transcriptProgressService struct {
	repo db_repos.ITranscriptProgressRepo
}

func NewTranscriptProgressService(repo db_repos.ITranscriptProgressRepo) ITranscriptProgressService {
	return &transcriptProgressService{repo: repo}
}

func (s *transcriptProgressService) Create(ctx context.Context, lessonID uint, body req.CreateTranscriptProgressReq) (*models.TranscriptProgress, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "transcriptId", body.TranscriptID, "lessonId", lessonID)
	log.Infow("creating transcript progress")

	progress := &models.TranscriptProgress{
		UserID:       userID,
		TranscriptID: body.TranscriptID,
		LessonID:     lessonID,
		CompletedAt:  time.Now(),
	}

	if createErr := s.repo.CreateOrIgnore(ctx, progress); createErr != nil {
		log.Errorw("failed to create transcript progress", "error", createErr)
		return nil, response.Internal("failed to create transcript progress")
	}

	log.Infow("transcript progress created")
	return progress, nil
}

func (s *transcriptProgressService) List(ctx context.Context, lessonID uint) ([]*models.TranscriptProgress, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "lessonId", lessonID)
	log.Infow("listing transcript progress")

	progress, findErr := s.repo.FindByUserAndLesson(ctx, userID, lessonID)
	if findErr != nil {
		log.Errorw("failed to list transcript progress", "error", findErr)
		return nil, response.Internal("failed to list transcript progress")
	}

	return progress, nil
}
