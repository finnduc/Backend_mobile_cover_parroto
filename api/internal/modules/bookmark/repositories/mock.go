package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.IBookmarkRepo = (*MockBookmarkRepo)(nil)

type MockBookmarkRepo struct {
	mock.Mock
}

func (m *MockBookmarkRepo) Create(ctx context.Context, userID string, lessonID uint) error {
	args := m.Called(ctx, userID, lessonID)
	return args.Error(0)
}

func (m *MockBookmarkRepo) Delete(ctx context.Context, userID string, lessonID uint) error {
	args := m.Called(ctx, userID, lessonID)
	return args.Error(0)
}

func (m *MockBookmarkRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Bookmark], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.Bookmark]), args.Error(1)
}
