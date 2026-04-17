package services

import (
	"context"

	"go-cover-parroto/internal/core/response"
	transcriptres "go-cover-parroto/internal/modules/transcript/dtos/res"
	"go-cover-parroto/internal/modules/transcript/repositories"
	"go-cover-parroto/internal/utils"
)

type ITranscriptService interface {
	GetByLesson(ctx context.Context, lessonID uint) ([]transcriptres.TranscriptRes, *response.AppError)
}

type transcriptService struct {
	repo repositories.ITranscriptRepo
}

func NewTranscriptService(repo repositories.ITranscriptRepo) ITranscriptService {
	return &transcriptService{repo: repo}
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
