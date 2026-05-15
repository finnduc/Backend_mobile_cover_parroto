package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.ILearningHistoryRepo = (*MockLearningHistoryRepo)(nil)

type MockLearningHistoryRepo struct {
	mock.Mock
}

func (m *MockLearningHistoryRepo) Upsert(ctx context.Context, history *models.LearningHistory) error {
	args := m.Called(ctx, history)
	return args.Error(0)
}

func (m *MockLearningHistoryRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.LearningHistory], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.LearningHistory]), args.Error(1)
}

func (m *MockLearningHistoryRepo) FindByUserAndLesson(ctx context.Context, userID string, lessonID uint) (*models.LearningHistory, error) {
	args := m.Called(ctx, userID, lessonID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.LearningHistory), args.Error(1)
}
