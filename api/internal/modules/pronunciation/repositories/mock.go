package repositories

import (
	"context"

	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.IPronunciationAttemptRepo = (*MockPronunciationAttemptRepo)(nil)
var _ db_repos.IPronunciationProgressRepo = (*MockPronunciationProgressRepo)(nil)

type MockPronunciationAttemptRepo struct {
	mock.Mock
}

func (m *MockPronunciationAttemptRepo) Create(ctx context.Context, attempt *models.PronunciationAttempt) error {
	args := m.Called(ctx, attempt)
	return args.Error(0)
}
func (m *MockPronunciationAttemptRepo) FindByID(ctx context.Context, id uint) (*models.PronunciationAttempt, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*models.PronunciationAttempt), args.Error(1)
}
func (m *MockPronunciationAttemptRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) ([]*models.PronunciationAttempt, error) {
	args := m.Called(ctx, userID, transcriptID)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).([]*models.PronunciationAttempt), args.Error(1)
}
func (m *MockPronunciationAttemptRepo) FindBestByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.PronunciationAttempt, error) {
	args := m.Called(ctx, userID, transcriptID)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*models.PronunciationAttempt), args.Error(1)
}
func (m *MockPronunciationAttemptRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockPronunciationProgressRepo struct {
	mock.Mock
}

func (m *MockPronunciationProgressRepo) Upsert(ctx context.Context, progress *models.PronunciationProgress) error {
	args := m.Called(ctx, progress)
	return args.Error(0)
}
func (m *MockPronunciationProgressRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) ([]*models.PronunciationProgress, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).([]*models.PronunciationProgress), args.Error(1)
}
func (m *MockPronunciationProgressRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.PronunciationProgress, error) {
	args := m.Called(ctx, userID, transcriptID)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*models.PronunciationProgress), args.Error(1)
}
