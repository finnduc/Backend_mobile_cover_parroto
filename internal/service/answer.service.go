package service

import (
	"context"
	"fmt"
	"go-cover-parroto/global"
	"go-cover-parroto/internal/models"
	"go-cover-parroto/internal/repo"
	pkgerrors "go-cover-parroto/pkg/errors"
	"go-cover-parroto/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type answerService struct {
	answerRepo  repo.IAnswerRepo
	attemptRepo repo.IAttemptRepo
	lessonRepo  repo.ILessonRepo
	progressSvc IProgressService
	threshold   float64
}

func NewAnswerService(
	answerRepo repo.IAnswerRepo,
	attemptRepo repo.IAttemptRepo,
	lessonRepo repo.ILessonRepo,
	progressSvc IProgressService,
	threshold float64,
) IAnswerService {
	return &answerService{
		answerRepo:  answerRepo,
		attemptRepo: attemptRepo,
		lessonRepo:  lessonRepo,
		progressSvc: progressSvc,
		threshold:   threshold,
	}
}

func (s *answerService) SubmitAnswer(ctx context.Context, userID uint, input SubmitAnswerInput) (*AnswerResult, error) {
	// 1. Load attempt
	attempt, err := s.attemptRepo.FindByID(ctx, input.AttemptID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("answerService.SubmitAnswer: %w", pkgerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("answerService.SubmitAnswer: %w", pkgerrors.ErrInternalServer)
	}

	// 2. Validate transcript belongs to lesson
	transcripts, err := s.lessonRepo.FindTranscriptsByLessonID(ctx, attempt.LessonID)
	if err != nil {
		return nil, fmt.Errorf("answerService.SubmitAnswer: %w", pkgerrors.ErrInternalServer)
	}

	var correctTranscript *models.Transcript
	for i := range transcripts {
		if transcripts[i].ID == input.TranscriptID {
			correctTranscript = &transcripts[i]
			break
		}
	}
	if correctTranscript == nil {
		return nil, fmt.Errorf("answerService.SubmitAnswer: transcript not in lesson: %w", pkgerrors.ErrInvalidInput)
	}

	// 3. Score
	score, isCorrect := utils.CalculateScore(input.AnswerText, correctTranscript.Content, s.threshold)

	// 4. Save answer
	answer := &models.UserAnswer{
		AttemptID:    input.AttemptID,
		TranscriptID: input.TranscriptID,
		AnswerText:   input.AnswerText,
		Score:        score,
		IsCorrect:    isCorrect,
	}
	if err := s.answerRepo.Upsert(ctx, answer); err != nil {
		return nil, fmt.Errorf("answerService.SubmitAnswer: %w", pkgerrors.ErrInternalServer)
	}

	// 5. Update progress (running average)
	if err := s.updateProgress(ctx, userID, attempt.LessonID, correctTranscript.Sequence, score); err != nil {
		global.Logger.Warn("answerService.SubmitAnswer: failed to update progress", zap.Error(err))
	}

	// 6. Check auto-complete
	s.checkAutoComplete(ctx, input.AttemptID, len(transcripts))

	return &AnswerResult{
		Score:       score,
		IsCorrect:   isCorrect,
		CorrectText: correctTranscript.Content,
	}, nil
}

func (s *answerService) BulkSubmit(ctx context.Context, userID uint, input BulkSubmitInput) (*BulkAnswerResult, error) {
	attempt, err := s.attemptRepo.FindByID(ctx, input.AttemptID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("answerService.BulkSubmit: %w", pkgerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("answerService.BulkSubmit: %w", pkgerrors.ErrInternalServer)
	}

	transcripts, err := s.lessonRepo.FindTranscriptsByLessonID(ctx, attempt.LessonID)
	if err != nil {
		return nil, fmt.Errorf("answerService.BulkSubmit: %w", pkgerrors.ErrInternalServer)
	}

	// Build ID -> transcript map
	tMap := make(map[uint]*models.Transcript, len(transcripts))
	for i := range transcripts {
		tMap[transcripts[i].ID] = &transcripts[i]
	}

	var answers []models.UserAnswer
	var results []AnswerResult

	for _, item := range input.Answers {
		t, ok := tMap[item.TranscriptID]
		if !ok {
			return nil, fmt.Errorf("answerService.BulkSubmit: transcript %d not in lesson: %w", item.TranscriptID, pkgerrors.ErrInvalidInput)
		}
		score, isCorrect := utils.CalculateScore(item.AnswerText, t.Content, s.threshold)
		answers = append(answers, models.UserAnswer{
			AttemptID:    input.AttemptID,
			TranscriptID: item.TranscriptID,
			AnswerText:   item.AnswerText,
			Score:        score,
			IsCorrect:    isCorrect,
		})
		results = append(results, AnswerResult{Score: score, IsCorrect: isCorrect, CorrectText: t.Content})
	}

	if err := s.answerRepo.BulkUpsert(ctx, answers); err != nil {
		return nil, fmt.Errorf("answerService.BulkSubmit: %w", pkgerrors.ErrInternalServer)
	}

	s.checkAutoComplete(ctx, input.AttemptID, len(transcripts))

	return &BulkAnswerResult{Results: results}, nil
}

func (s *answerService) updateProgress(ctx context.Context, userID, lessonID uint, sequence int, score float64) error {
	p, err := s.progressSvc.GetProgress(ctx, userID, lessonID)
	if err != nil {
		return err
	}
	// Running average
	p.ScoreAvg = (p.ScoreAvg*float64(p.AnswerCount) + score) / float64(p.AnswerCount+1)
	p.AnswerCount++
	p.LastSequence = sequence
	return nil
}

func (s *answerService) checkAutoComplete(ctx context.Context, attemptID uint, totalTranscripts int) {
	count, err := s.attemptRepo.CountAnswers(ctx, attemptID)
	if err != nil || int(count) < totalTranscripts {
		return
	}
	// Calculate average score and mark completed
	_ = s.attemptRepo.MarkCompleted(ctx, attemptID, float64(count)/float64(totalTranscripts))
}
