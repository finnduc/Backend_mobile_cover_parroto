package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.IVocabularyItemRepo = (*MockVocabularyItemRepo)(nil)

type MockVocabularyItemRepo struct {
	mock.Mock
}

func (m *MockVocabularyItemRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyItem], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.VocabularyItem]), args.Error(1)
}

func (m *MockVocabularyItemRepo) FindByID(ctx context.Context, id uint) (*models.VocabularyItem, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.VocabularyItem), args.Error(1)
}

func (m *MockVocabularyItemRepo) Create(ctx context.Context, item *models.VocabularyItem) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func (m *MockVocabularyItemRepo) Update(ctx context.Context, item *models.VocabularyItem) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func (m *MockVocabularyItemRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
