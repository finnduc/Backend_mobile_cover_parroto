package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.IVocabularyCategoryRepo = (*MockVocabularyCategoryRepo)(nil)

type MockVocabularyCategoryRepo struct {
	mock.Mock
}

func (m *MockVocabularyCategoryRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyCategory], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.VocabularyCategory]), args.Error(1)
}

func (m *MockVocabularyCategoryRepo) FindByID(ctx context.Context, id uint) (*models.VocabularyCategory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.VocabularyCategory), args.Error(1)
}

func (m *MockVocabularyCategoryRepo) Create(ctx context.Context, category *models.VocabularyCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockVocabularyCategoryRepo) Update(ctx context.Context, category *models.VocabularyCategory) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockVocabularyCategoryRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
