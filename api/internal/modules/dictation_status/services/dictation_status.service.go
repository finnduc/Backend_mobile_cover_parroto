package services

import (
	"context"
	"time"

	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/dictation_status/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "dictation_status")
}

type IDictationStatusService interface {
	Create(ctx context.Context, body req.CreateDictationStatusReq) (*models.DictationStatus, *response.AppError)
	List(ctx context.Context, query req.ListDictationStatusQuery) (*response.PaginatedResult[*models.DictationStatus], *response.AppError)
}

type dictationStatusService struct {
	repo db_repos.IDictationStatusRepo
}

func NewDictationStatusService(repo db_repos.IDictationStatusRepo) IDictationStatusService {
	return &dictationStatusService{repo: repo}
}

func (s *dictationStatusService) Create(ctx context.Context, body req.CreateDictationStatusReq) (*models.DictationStatus, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "transcriptId", body.TranscriptID)
	log.Infow("creating dictation status")

	status := &models.DictationStatus{
		UserID:       userID,
		TranscriptID: body.TranscriptID,
		LessonID:     body.LessonID,
		CompletedAt:  time.Now(),
	}

	if createErr := s.repo.CreateOrIgnore(ctx, status); createErr != nil {
		log.Errorw("failed to create dictation status", "error", createErr)
		return nil, response.Internal("failed to create dictation status")
	}

	log.Infow("dictation status created")
	return status, nil
}

func (s *dictationStatusService) List(ctx context.Context, query req.ListDictationStatusQuery) (*response.PaginatedResult[*models.DictationStatus], *response.AppError) {
	log := sLog()
	log.Infow("listing dictation status")

	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list dictation status", "error", err)
		return nil, response.Internal("failed to list dictation status")
	}
	return result, nil
}
