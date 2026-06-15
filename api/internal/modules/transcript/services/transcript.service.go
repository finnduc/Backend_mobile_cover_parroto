package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go-cover-parroto/internal/database/transaction"
	req "go-cover-parroto/internal/modules/transcript/dtos/req"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "transcript")
}

type ITranscriptService interface {
	GetByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, *response.AppError)
	GetByID(ctx context.Context, id uint) (*models.Transcript, *response.AppError)
	Create(ctx context.Context, body req.CreateTranscriptReq) (*models.Transcript, *response.AppError)
	BulkCreate(ctx context.Context, lessonID uint, body req.BulkCreateTranscriptReq) ([]*models.Transcript, *response.AppError)
	ReplaceByLesson(ctx context.Context, lessonID uint, body req.BulkCreateTranscriptReq) ([]*models.Transcript, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateTranscriptReq) (*models.Transcript, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
}

type transcriptService struct {
	repo db_repos.ITranscriptRepo
	uow  transaction.UnitOfWork
}

func NewTranscriptService(repo db_repos.ITranscriptRepo, uow transaction.UnitOfWork) ITranscriptService {
	return &transcriptService{repo: repo, uow: uow}
}

func (s *transcriptService) GetByID(ctx context.Context, id uint) (*models.Transcript, *response.AppError) {
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
	return transcript, nil
}

func (s *transcriptService) GetByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, *response.AppError) {
	log := sLog().With("lessonId", lessonID)
	log.Infow("getting transcripts by lesson")
	transcripts, err := s.repo.FindByLesson(ctx, lessonID)
	if err != nil {
		log.Errorw("failed to get transcripts by lesson", "error", err)
		return nil, response.Internal("failed to get transcripts")
	}
	return transcripts, nil
}

func (s *transcriptService) Create(ctx context.Context, body req.CreateTranscriptReq) (*models.Transcript, *response.AppError) {
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
	return transcript, nil
}

func (s *transcriptService) BulkCreate(ctx context.Context, lessonID uint, body req.BulkCreateTranscriptReq) ([]*models.Transcript, *response.AppError) {
	log := sLog().With("lessonId", lessonID, "count", len(body.Transcripts))
	log.Infow("bulk creating transcripts")

	transcripts := make([]*models.Transcript, len(body.Transcripts))
	for i, t := range body.Transcripts {
		transcripts[i] = &models.Transcript{
			LessonID:       lessonID,
			Sequence:       t.Sequence,
			Content:        t.Content,
			Phonetic:       t.Phonetic,
			Vietnamese:     t.Vietnamese,
			StartTimestamp: t.StartTimestamp,
			EndTimestamp:   t.EndTimestamp,
		}
	}

	if err := s.repo.BulkCreate(ctx, transcripts); err != nil {
		log.Errorw("failed to bulk create transcripts", "error", err)
		return nil, response.Internal("failed to create transcripts")
	}

	log.Infow("transcripts created", "count", len(transcripts))
	return transcripts, nil
}

func (s *transcriptService) ReplaceByLesson(ctx context.Context, lessonID uint, body req.BulkCreateTranscriptReq) ([]*models.Transcript, *response.AppError) {
	log := sLog().With("lessonId", lessonID, "count", len(body.Transcripts))
	log.Infow("replacing transcripts")

	var result []*models.Transcript
	err := s.uow.Do(ctx, func(ctx context.Context, p transaction.IProvider) error {
		tRepo := p.Transcript()

		if err := tRepo.DeleteByLesson(ctx, lessonID); err != nil {
			return err
		}

		transcripts := make([]*models.Transcript, len(body.Transcripts))
		for i, t := range body.Transcripts {
			transcripts[i] = &models.Transcript{
				LessonID:       lessonID,
				Sequence:       t.Sequence,
				Content:        t.Content,
				Phonetic:       t.Phonetic,
				Vietnamese:     t.Vietnamese,
				StartTimestamp: t.StartTimestamp,
				EndTimestamp:   t.EndTimestamp,
			}
		}

		if err := tRepo.BulkCreate(ctx, transcripts); err != nil {
			return err
		}

		result = transcripts
		return nil
	})

	if err != nil {
		log.Errorw("failed to replace transcripts", "error", err)
		return nil, response.Internal("failed to replace transcripts")
	}

	log.Infow("transcripts replaced")
	return result, nil
}

func (s *transcriptService) Update(ctx context.Context, id uint, body req.UpdateTranscriptReq) (*models.Transcript, *response.AppError) {
	log := sLog().With("transcriptId", id)
	log.Infow("updating transcript")
	transcript, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("transcript not found")
		}
		log.Errorw("failed to get transcript for update", "error", err)
		return nil, response.Internal("failed to get transcript")
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
	return transcript, nil
}

func (s *transcriptService) Delete(ctx context.Context, id uint) *response.AppError {
	log := sLog().With("transcriptId", id)
	log.Infow("deleting transcript")
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return response.NotFound("transcript not found")
		}
		log.Errorw("failed to get transcript for delete", "error", err)
		return response.Internal("failed to get transcript")
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Errorw("failed to delete transcript", "error", err)
		return response.Internal("failed to delete transcript")
	}
	log.Infow("transcript deleted")
	return nil
}
