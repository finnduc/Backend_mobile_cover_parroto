package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/database/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTranscriptRepo struct{ mock.Mock }

func (m *mockTranscriptRepo) FindByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, error) {
	args := m.Called(ctx, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Transcript), args.Error(1)
}

func TestGetByLesson_Success(t *testing.T) {
	repo := new(mockTranscriptRepo)
	svc := NewTranscriptService(repo)

	transcripts := []*models.Transcript{
		{ID: 1, LessonID: 2, Sequence: 1, Content: "Hello world", Phonetic: "ˈhɛloʊ wɜːld", Vietnamese: "Xin chào thế giới", StartTimestamp: 0.0, EndTimestamp: 2.5},
		{ID: 2, LessonID: 2, Sequence: 2, Content: "How are you", Phonetic: "haʊ ɑːr juː", Vietnamese: "Bạn có khỏe không", StartTimestamp: 3.0, EndTimestamp: 5.0},
	}
	repo.On("FindByLesson", mock.Anything, uint(2)).Return(transcripts, nil)

	result, appErr := svc.GetByLesson(context.Background(), 2)

	assert.Nil(t, appErr)
	assert.Len(t, result, 2)
	assert.Equal(t, 1, result[0].Sequence)
	assert.Equal(t, "Hello world", result[0].Content)
	assert.Equal(t, "Xin chào thế giới", result[0].Vietnamese)
	repo.AssertExpectations(t)
}

func TestGetByLesson_Error(t *testing.T) {
	repo := new(mockTranscriptRepo)
	svc := NewTranscriptService(repo)

	repo.On("FindByLesson", mock.Anything, uint(99)).Return(nil, errors.New("db error"))

	result, appErr := svc.GetByLesson(context.Background(), 99)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}

func TestGetByLesson_Empty(t *testing.T) {
	repo := new(mockTranscriptRepo)
	svc := NewTranscriptService(repo)

	repo.On("FindByLesson", mock.Anything, uint(1)).Return([]*models.Transcript{}, nil)

	result, appErr := svc.GetByLesson(context.Background(), 1)

	assert.Nil(t, appErr)
	assert.Len(t, result, 0)
}
