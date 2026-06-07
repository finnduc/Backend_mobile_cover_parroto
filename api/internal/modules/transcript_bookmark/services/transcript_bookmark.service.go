package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/transcript_bookmark/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "transcript_bookmark")
}

type ITranscriptBookmarkService interface {
	Create(ctx context.Context, body req.CreateTranscriptBookmarkReq) (*models.TranscriptBookmark, *response.AppError)
	List(ctx context.Context, query req.ListTranscriptBookmarkQuery) (*response.PaginatedResult[*models.TranscriptBookmark], *response.AppError)
	UpdateNote(ctx context.Context, id uint, body req.UpdateTranscriptBookmarkNoteReq) (*models.TranscriptBookmark, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
}

type transcriptBookmarkService struct {
	repo db_repos.ITranscriptBookmarkRepo
}

func NewTranscriptBookmarkService(repo db_repos.ITranscriptBookmarkRepo) ITranscriptBookmarkService {
	return &transcriptBookmarkService{repo: repo}
}

func (s *transcriptBookmarkService) Create(ctx context.Context, body req.CreateTranscriptBookmarkReq) (*models.TranscriptBookmark, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "transcriptId", body.TranscriptID)
	log.Infow("creating transcript bookmark")

	bookmark := &models.TranscriptBookmark{
		UserID:       userID,
		TranscriptID: body.TranscriptID,
		Note:         body.Note,
	}

	if createErr := s.repo.Create(ctx, bookmark); createErr != nil {
		log.Errorw("failed to create transcript bookmark", "error", createErr)
		return nil, response.Internal("failed to create transcript bookmark")
	}

	log.Infow("transcript bookmark created")
	return bookmark, nil
}

func (s *transcriptBookmarkService) List(ctx context.Context, query req.ListTranscriptBookmarkQuery) (*response.PaginatedResult[*models.TranscriptBookmark], *response.AppError) {
	log := sLog()
	log.Infow("listing transcript bookmarks")

	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list transcript bookmarks", "error", err)
		return nil, response.Internal("failed to list transcript bookmarks")
	}
	return result, nil
}

func (s *transcriptBookmarkService) UpdateNote(ctx context.Context, id uint, body req.UpdateTranscriptBookmarkNoteReq) (*models.TranscriptBookmark, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", userID, "id", id)
	log.Infow("updating transcript bookmark note")

	bookmark, findErr := s.repo.FindByID(ctx, id)
	if findErr != nil {
		if errors.Is(findErr, coreError.ErrNotFound) {
			return nil, response.NotFound("transcript bookmark not found")
		}
		log.Errorw("failed to find transcript bookmark", "error", findErr)
		return nil, response.Internal("failed to find transcript bookmark")
	}

	if accessErr := policy.Allow(ctx, bookmark.UserID); accessErr != nil {
		return nil, accessErr
	}

	bookmark.Note = body.Note

	if updateErr := s.repo.Update(ctx, bookmark); updateErr != nil {
		log.Errorw("failed to update transcript bookmark", "error", updateErr)
		return nil, response.Internal("failed to update transcript bookmark")
	}

	log.Infow("transcript bookmark updated")
	return bookmark, nil
}

func (s *transcriptBookmarkService) Delete(ctx context.Context, id uint) *response.AppError {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return err
	}

	log := sLog().With("userId", userID, "id", id)
	log.Infow("deleting transcript bookmark")

	bookmark, findErr := s.repo.FindByID(ctx, id)
	if findErr != nil {
		if errors.Is(findErr, coreError.ErrNotFound) {
			return response.NotFound("transcript bookmark not found")
		}
		log.Errorw("failed to find transcript bookmark", "error", findErr)
		return response.Internal("failed to find transcript bookmark")
	}

	if accessErr := policy.Allow(ctx, bookmark.UserID); accessErr != nil {
		return accessErr
	}

	if delErr := s.repo.Delete(ctx, id); delErr != nil {
		log.Errorw("failed to delete transcript bookmark", "error", delErr)
		return response.Internal("failed to delete transcript bookmark")
	}

	log.Infow("transcript bookmark deleted")
	return nil
}
