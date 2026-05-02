package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	req "go-cover-parroto/internal/modules/transcript/dtos/req"
	transcriptres "go-cover-parroto/internal/modules/transcript/dtos/res"
	"go-cover-parroto/internal/modules/transcript/repositories"
	"go-cover-parroto/internal/utils"
)

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
	transcript, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("transcript not found")
		}
		return nil, response.Internal("failed to get transcript")
	}
	var result transcriptres.TranscriptRes
	if err := utils.MapToDTO(transcript, &result); err != nil {
		return nil, response.Internal("failed to map transcript")
	}
	return &result, nil
}

func (s *transcriptService) GetByLesson(ctx context.Context, lessonID uint) ([]transcriptres.TranscriptRes, *response.AppError) {
	transcripts, err := s.repo.FindByLesson(ctx, lessonID)
	if err != nil {
		return nil, response.Internal("failed to get transcripts")
	}

	var result []transcriptres.TranscriptRes
	if err := utils.MapToDTOs(transcripts, &result); err != nil {
		return nil, response.Internal("failed to map transcripts")
	}
	return result, nil
}

func (s *transcriptService) Create(ctx context.Context, body req.CreateTranscriptReq) (*transcriptres.TranscriptRes, *response.AppError) {
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
		return nil, response.Internal("failed to create transcript")
	}
	var result transcriptres.TranscriptRes
	if err := utils.MapToDTO(transcript, &result); err != nil {
		return nil, response.Internal("failed to map transcript")
	}
	return &result, nil
}

func (s *transcriptService) Update(ctx context.Context, id uint, body req.UpdateTranscriptReq) (*transcriptres.TranscriptRes, *response.AppError) {
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
		return nil, response.Internal("failed to update transcript")
	}
	var result transcriptres.TranscriptRes
	if err := utils.MapToDTO(transcript, &result); err != nil {
		return nil, response.Internal("failed to map transcript")
	}
	return &result, nil
}

func (s *transcriptService) Delete(ctx context.Context, id uint) *response.AppError {
	if err := s.repo.Delete(ctx, id); err != nil {
		return response.Internal("failed to delete transcript")
	}
	return nil
}
