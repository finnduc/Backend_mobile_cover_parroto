package services

import (
	"context"
	"errors"
	"time"

	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/shadowing_status/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "shadowing_status")
}

type IShadowingStatusService interface {
	Create(ctx context.Context, body req.CreateShadowingStatusReq) (*models.ShadowingStatus, *response.AppError)
	List(ctx context.Context, query req.ListShadowingStatusQuery) (*response.PaginatedResult[*models.ShadowingStatus], *response.AppError)
}

type shadowingStatusService struct {
	repo db_repos.IShadowingStatusRepo
}

func NewShadowingStatusService(repo db_repos.IShadowingStatusRepo) IShadowingStatusService {
	return &shadowingStatusService{repo: repo}
}

func (s *shadowingStatusService) Create(ctx context.Context, body req.CreateShadowingStatusReq) (*models.ShadowingStatus, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "transcriptId", body.TranscriptID)
	log.Infow("creating shadowing status")

	existing, findErr := s.repo.FindByUserAndTranscript(ctx, userID, body.TranscriptID)
	if findErr != nil && !errors.Is(findErr, gorm.ErrRecordNotFound) {
		log.Errorw("failed to check existing status", "error", findErr)
		return nil, response.Internal("failed to check shadowing status")
	}
	if existing != nil {
		return nil, response.Conflict("transcript already completed")
	}

	status := &models.ShadowingStatus{
		UserID:       userID,
		TranscriptID: body.TranscriptID,
		LessonID:     body.LessonID,
		CompletedAt:  time.Now(),
	}

	if createErr := s.repo.Create(ctx, status); createErr != nil {
		log.Errorw("failed to create shadowing status", "error", createErr)
		return nil, response.Internal("failed to create shadowing status")
	}

	log.Infow("shadowing status created")
	return status, nil
}

func (s *shadowingStatusService) List(ctx context.Context, query req.ListShadowingStatusQuery) (*response.PaginatedResult[*models.ShadowingStatus], *response.AppError) {
	log := sLog()
	log.Infow("listing shadowing status")

	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list shadowing status", "error", err)
		return nil, response.Internal("failed to list shadowing status")
	}
	return result, nil
}
