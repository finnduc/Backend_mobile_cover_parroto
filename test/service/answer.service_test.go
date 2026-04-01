package service

import (
	"context"
	"go-cover-parroto/internal/mocks"
	"go-cover-parroto/internal/models"
	"go-cover-parroto/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmitAnswer(t *testing.T) {
	mockAnswerRepo := new(mocks.MockAnswerRepo)
	mockAttemptRepo := new(mocks.MockAttemptRepo)
	mockLessonRepo := new(mocks.MockLessonRepo)
	mockProgressSvc := new(mocks.MockProgressService)

	s := service.NewAnswerService(
		mockAnswerRepo,
		mockAttemptRepo,
		mockLessonRepo,
		mockProgressSvc,
		0.8, // threshold
	)

	userID := uint(1)
	input := service.SubmitAnswerInput{
		AttemptID:    10,
		TranscriptID: 100,
		AnswerText:   "hello world",
	}

	// 1. Load attempt
	mockAttempt := &models.Attempt{Base: models.Base{ID: 10}, LessonID: 5, UserID: userID}
	mockAttemptRepo.On("FindByID", mock.Anything, input.AttemptID).Return(mockAttempt, nil)

	// 2. Validate transcript
	mockTranscripts := []models.Transcript{
		{Base: models.Base{ID: 100}, LessonID: 5, Content: "hello world", Sequence: 1},
	}
	mockLessonRepo.On("FindTranscriptsByLessonID", mock.Anything, mockAttempt.LessonID).Return(mockTranscripts, nil)

	// 3. Score & Save
	mockAnswerRepo.On("Upsert", mock.Anything, mock.AnythingOfType("*models.UserAnswer")).Return(nil)

	// 4. Update progress
	mockProgress := &models.UserProgress{UserID: userID, LessonID: 5, AnswerCount: 0, ScoreAvg: 0}
	mockProgressSvc.On("GetProgress", mock.Anything, userID, mockAttempt.LessonID).Return(mockProgress, nil)

	// 5. Check auto-complete
	mockAttemptRepo.On("CountAnswers", mock.Anything, input.AttemptID).Return(int64(1), nil)
	mockAttemptRepo.On("MarkCompleted", mock.Anything, input.AttemptID, mock.AnythingOfType("float64")).Return(nil)

	result, err := s.SubmitAnswer(context.Background(), userID, input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.IsCorrect)
	assert.Equal(t, 1.0, result.Score)

	mockAttemptRepo.AssertExpectations(t)
	mockLessonRepo.AssertExpectations(t)
	mockAnswerRepo.AssertExpectations(t)
	mockProgressSvc.AssertExpectations(t)
}
