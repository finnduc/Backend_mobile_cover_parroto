package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	req "go-cover-parroto/internal/modules/transcript/dtos/req"
	transcriptres "go-cover-parroto/internal/modules/transcript/dtos/res"
	"go-cover-parroto/internal/modules/transcript/repositories"
	"go-cover-parroto/internal/utils"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "transcript")
}

type ITranscriptService interface {
	GetByLesson(ctx context.Context, lessonID uint) ([]transcriptres.TranscriptRes, *response.AppError)
	GetByID(ctx context.Context, id uint) (*transcriptres.TranscriptRes, *response.AppError)
	Create(ctx context.Context, body req.CreateTranscriptReq) (*transcriptres.TranscriptRes, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateTranscriptReq) (*transcriptres.TranscriptRes, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
}

type transcriptService struct {
	repo repositories.ITranscriptRepo
}

func NewTranscriptService(repo repositories.ITranscriptRepo) ITranscriptService {
	return &transcriptService{repo: repo}
}

func (s *transcriptService) GetByID(ctx context.Context, id uint) (*transcriptres.TranscriptRes, *response.AppError) {
	log := sLog().With("transcriptId", id)
	log.Infow("getting transcript by id")
	transcript, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("transcript not found")
		}
		log.Errorw("failed to get transcript", "error", err)
		return nil, response.Internal("failed to get transcript")
	}
	var result transcriptres.TranscriptRes
	if err := utils.MapToDTO(transcript, &result); err != nil {
		return nil, response.Internal("failed to map transcript")
	}
	return &result, nil
}

func (s *transcriptService) GetByLesson(ctx context.Context, lessonID uint) ([]transcriptres.TranscriptRes, *response.AppError) {
	log := sLog().With("lessonId", lessonID)
	log.Infow("getting transcripts by lesson")
	transcripts, err := s.repo.FindByLesson(ctx, lessonID)
	if err != nil {
		log.Errorw("failed to get transcripts by lesson", "error", err)
		return nil, response.Internal("failed to get transcripts")
	}

	var result []transcriptres.TranscriptRes
	if err := utils.MapToDTOs(transcripts, &result); err != nil {
		return nil, response.Internal("failed to map transcripts")
	}
	return result, nil
}

func (s *transcriptService) Create(ctx context.Context, body req.CreateTranscriptReq) (*transcriptres.TranscriptRes, *response.AppError) {
	log := sLog().With("lessonId", body.LessonID)
	log.Infow("creating transcript")
	transcript := &models.Transcript{
		LessonID:       body.LessonID,
		Sequence:       body.Sequence,
		Content:        body.Content,
		Phonetic:       body.Phonetic,
		Vietnamese:     body.Vietnamese,
		StartTimestamp: body.StartTimestamp,
		EndTimestamp:   body.EndTimestamp,
	}
	if err := s.repo.Create(ctx, transcript); err != nil {
		log.Errorw("failed to create transcript", "error", err)
		return nil, response.Internal("failed to create transcript")
	}
	log.Infow("transcript created", "transcriptId", transcript.ID)
	var result transcriptres.TranscriptRes
	if err := utils.MapToDTO(transcript, &result); err != nil {
		return nil, response.Internal("failed to map transcript")
	}
	return &result, nil
}

func (s *transcriptService) Update(ctx context.Context, id uint, body req.UpdateTranscriptReq) (*transcriptres.TranscriptRes, *response.AppError) {
	log := sLog().With("transcriptId", id)
	log.Infow("updating transcript")
	transcript, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, response.NotFound("transcript not found")
	}
	if body.Sequence != nil {
		transcript.Sequence = *body.Sequence
	}
	if body.Content != nil {
		transcript.Content = *body.Content
	}
	if body.Phonetic != nil {
		transcript.Phonetic = *body.Phonetic
	}
	if body.Vietnamese != nil {
		transcript.Vietnamese = *body.Vietnamese
	}
	if body.StartTimestamp != nil {
		transcript.StartTimestamp = *body.StartTimestamp
	}
	if body.EndTimestamp != nil {
		transcript.EndTimestamp = *body.EndTimestamp
	}
	if updateErr := s.repo.Update(ctx, transcript); updateErr != nil {
		log.Errorw("failed to update transcript", "error", updateErr)
		return nil, response.Internal("failed to update transcript")
	}
	log.Infow("transcript updated")
	var result transcriptres.TranscriptRes
	if err := utils.MapToDTO(transcript, &result); err != nil {
		return nil, response.Internal("failed to map transcript")
	}
	return &result, nil
}

func (s *transcriptService) Delete(ctx context.Context, id uint) *response.AppError {
	log := sLog().With("transcriptId", id)
	log.Infow("deleting transcript")
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Errorw("failed to delete transcript", "error", err)
		return response.Internal("failed to delete transcript")
	}
	log.Infow("transcript deleted")
	return nil
}
