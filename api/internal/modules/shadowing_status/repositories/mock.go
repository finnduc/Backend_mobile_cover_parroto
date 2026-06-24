package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.IShadowingStatusRepo = (*MockShadowingStatusRepo)(nil)

type MockShadowingStatusRepo struct {
	mock.Mock
}

func (m *MockShadowingStatusRepo) Create(ctx context.Context, status *models.ShadowingStatus) error {
	args := m.Called(ctx, status)
	return args.Error(0)
}

func (m *MockShadowingStatusRepo) Update(ctx context.Context, status *models.ShadowingStatus) error {
	args := m.Called(ctx, status)
	return args.Error(0)
}

func (m *MockShadowingStatusRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.ShadowingStatus, error) {
	args := m.Called(ctx, userID, transcriptID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ShadowingStatus), args.Error(1)
}

func (m *MockShadowingStatusRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.ShadowingStatus], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.ShadowingStatus]), args.Error(1)
}
