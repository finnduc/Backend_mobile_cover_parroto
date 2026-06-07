package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.ITranscriptProgressRepo = (*MockTranscriptProgressRepo)(nil)

type MockTranscriptProgressRepo struct {
	mock.Mock
}

func (m *MockTranscriptProgressRepo) CreateOrIgnore(ctx context.Context, progress *models.TranscriptProgress) error {
	args := m.Called(ctx, progress)
	return args.Error(0)
}

func (m *MockTranscriptProgressRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) ([]*models.TranscriptProgress, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.TranscriptProgress), args.Error(1)
}
