package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.ITranscriptRepo = (*MockTranscriptRepo)(nil)

type MockTranscriptRepo struct {
	mock.Mock
}

func (m *MockTranscriptRepo) FindByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, error) {
	args := m.Called(ctx, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Transcript), args.Error(1)
}

func (m *MockTranscriptRepo) FindByID(ctx context.Context, id uint) (*models.Transcript, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Transcript), args.Error(1)
}

func (m *MockTranscriptRepo) Create(ctx context.Context, transcript *models.Transcript) error {
	args := m.Called(ctx, transcript)
	return args.Error(0)
}

func (m *MockTranscriptRepo) BulkCreate(ctx context.Context, transcripts []*models.Transcript) error {
	args := m.Called(ctx, transcripts)
	return args.Error(0)
}

func (m *MockTranscriptRepo) Update(ctx context.Context, transcript *models.Transcript) error {
	args := m.Called(ctx, transcript)
	return args.Error(0)
}

func (m *MockTranscriptRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTranscriptRepo) DeleteByLesson(ctx context.Context, lessonID uint) error {
	args := m.Called(ctx, lessonID)
	return args.Error(0)
}
