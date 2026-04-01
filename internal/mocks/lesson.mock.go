package mocks

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/internal/repo"

	"github.com/stretchr/testify/mock"
)

type MockLessonRepo struct {
	mock.Mock
}

func (m *MockLessonRepo) FindAll(ctx context.Context, filter repo.LessonFilter) ([]models.Lesson, int64, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]models.Lesson), args.Get(1).(int64), args.Error(2)
}

func (m *MockLessonRepo) Create(ctx context.Context, lesson *models.Lesson, categoryIDs []uint, transcripts []models.Transcript) error {
	args := m.Called(ctx, lesson, categoryIDs, transcripts)
	return args.Error(0)
}

func (m *MockLessonRepo) FindByID(ctx context.Context, id uint) (*models.Lesson, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Lesson), args.Error(1)
}

func (m *MockLessonRepo) FindTranscriptsByLessonID(ctx context.Context, lessonID uint) ([]models.Transcript, error) {
	args := m.Called(ctx, lessonID)
	return args.Get(0).([]models.Transcript), args.Error(1)
}