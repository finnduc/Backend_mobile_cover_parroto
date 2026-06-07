package services

import (
	"context"
	"encoding/json"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "pronunciation")
}

type IPronunciationService interface {
	Assess(ctx context.Context, lessonID, transcriptID uint, referenceText string, audioData []byte) (*PronunciationResult, *models.PronunciationAttempt, *response.AppError)
	DeleteAttempt(ctx context.Context, attemptID uint) *response.AppError
	ListProgress(ctx context.Context, lessonID uint) ([]*models.PronunciationProgress, *response.AppError)
	UpdateProgress(ctx context.Context, transcriptID uint) (*models.PronunciationProgress, *response.AppError)
}

type pronunciationService struct {
	attemptRepo  db_repos.IPronunciationAttemptRepo
	progressRepo db_repos.IPronunciationProgressRepo
}

func NewPronunciationService(
	attemptRepo db_repos.IPronunciationAttemptRepo,
	progressRepo db_repos.IPronunciationProgressRepo,
) IPronunciationService {
	return &pronunciationService{attemptRepo: attemptRepo, progressRepo: progressRepo}
}

func (s *pronunciationService) Assess(ctx context.Context, lessonID, transcriptID uint, referenceText string, audioData []byte) (*PronunciationResult, *models.PronunciationAttempt, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, nil, err
	}

	log := sLog().With("userId", userID, "lessonId", lessonID, "transcriptId", transcriptID)
	log.Infow("assessing pronunciation")

	result, azureErr := assessPronunciation(audioData, referenceText)
	if azureErr != nil {
		log.Errorw("azure assessment failed", "error", azureErr)
		return nil, nil, response.Internal("pronunciation assessment failed")
	}

	wordJSON, _ := json.Marshal(result.Words)

	attempt := &models.PronunciationAttempt{
		UserID:            userID,
		LessonID:          lessonID,
		TranscriptID:      transcriptID,
		ReferenceText:     referenceText,
		OverallScore:      result.OverallScore,
		AccuracyScore:     result.Scores.Accuracy,
		FluencyScore:      result.Scores.Fluency,
		CompletenessScore: result.Scores.Completeness,
		ProsodyScore:      result.Scores.Prosody,
		Feedback:          result.Feedback,
		WordResults:       string(wordJSON),
	}

	if createErr := s.attemptRepo.Create(ctx, attempt); createErr != nil {
		log.Errorw("failed to save attempt", "error", createErr)
		return nil, nil, response.Internal("failed to save pronunciation attempt")
	}

	best, bestErr := s.attemptRepo.FindBestByUserAndTranscript(ctx, userID, transcriptID)
	if bestErr == nil && best != nil {
		progress := &models.PronunciationProgress{
			UserID:        userID,
			TranscriptID:  transcriptID,
			LessonID:      lessonID,
			BestAttemptID: &best.ID,
			BestScore:     &best.OverallScore,
			Feedback:      best.Feedback,
		}
		if upsertErr := s.progressRepo.Upsert(ctx, progress); upsertErr != nil {
			log.Errorw("failed to update progress", "error", upsertErr)
		}
	}

	log.Infow("pronunciation assessed", "score", result.OverallScore)
	return result, attempt, nil
}

func (s *pronunciationService) DeleteAttempt(ctx context.Context, attemptID uint) *response.AppError {
	_, err := policy.GetUserID(ctx)
	if err != nil {
		return err
	}

	attempt, findErr := s.attemptRepo.FindByID(ctx, attemptID)
	if findErr != nil {
		if findErr.Error() == coreError.ErrNotFound.Error() {
			return response.NotFound("pronunciation attempt not found")
		}
		return response.Internal("failed to find attempt")
	}

	if accessErr := policy.Allow(ctx, attempt.UserID); accessErr != nil {
		return accessErr
	}

	if delErr := s.attemptRepo.Delete(ctx, attemptID); delErr != nil {
		return response.Internal("failed to delete attempt")
	}

	return nil
}

func (s *pronunciationService) ListProgress(ctx context.Context, lessonID uint) ([]*models.PronunciationProgress, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	progress, findErr := s.progressRepo.FindByUserAndLesson(ctx, userID, lessonID)
	if findErr != nil {
		return nil, response.Internal("failed to list progress")
	}
	return progress, nil
}

func (s *pronunciationService) UpdateProgress(ctx context.Context, transcriptID uint) (*models.PronunciationProgress, *response.AppError) {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	best, findErr := s.attemptRepo.FindBestByUserAndTranscript(ctx, userID, transcriptID)
	if findErr != nil {
		if findErr.Error() == coreError.ErrNotFound.Error() {
			return nil, response.NotFound("no attempts found for this transcript")
		}
		return nil, response.Internal("failed to find best attempt")
	}

	progress := &models.PronunciationProgress{
		UserID:        userID,
		TranscriptID:  transcriptID,
		LessonID:      best.LessonID,
		BestAttemptID: &best.ID,
		BestScore:     &best.OverallScore,
		Feedback:      best.Feedback,
	}

	if upsertErr := s.progressRepo.Upsert(ctx, progress); upsertErr != nil {
		return nil, response.Internal("failed to update progress")
	}

	saved, _ := s.progressRepo.FindByUserAndTranscript(ctx, userID, transcriptID)
	if saved == nil {
		saved = progress
	}

	return saved, nil
}
