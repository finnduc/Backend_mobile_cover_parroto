package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.IDictationStatusRepo = (*MockDictationStatusRepo)(nil)

type MockDictationStatusRepo struct {
	mock.Mock
}

func (m *MockDictationStatusRepo) CreateOrIgnore(ctx context.Context, status *models.DictationStatus) error {
	args := m.Called(ctx, status)
	return args.Error(0)
}

func (m *MockDictationStatusRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.DictationStatus], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.DictationStatus]), args.Error(1)
}
