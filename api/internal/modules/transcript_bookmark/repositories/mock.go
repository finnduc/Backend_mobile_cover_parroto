package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.ITranscriptBookmarkRepo = (*MockTranscriptBookmarkRepo)(nil)

type MockTranscriptBookmarkRepo struct {
	mock.Mock
}

func (m *MockTranscriptBookmarkRepo) Create(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	args := m.Called(ctx, bookmark)
	return args.Error(0)
}

func (m *MockTranscriptBookmarkRepo) FindByID(ctx context.Context, id uint) (*models.TranscriptBookmark, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.TranscriptBookmark), args.Error(1)
}

func (m *MockTranscriptBookmarkRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.TranscriptBookmark], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.TranscriptBookmark]), args.Error(1)
}

func (m *MockTranscriptBookmarkRepo) Update(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	args := m.Called(ctx, bookmark)
	return args.Error(0)
}

func (m *MockTranscriptBookmarkRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
