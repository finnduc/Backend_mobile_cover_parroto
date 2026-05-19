package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.IVocabularyDeckRepo = (*MockVocabularyDeckRepo)(nil)

type MockVocabularyDeckRepo struct {
	mock.Mock
}

func (m *MockVocabularyDeckRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.VocabularyDeck], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.VocabularyDeck]), args.Error(1)
}

func (m *MockVocabularyDeckRepo) FindByID(ctx context.Context, id uint) (*models.VocabularyDeck, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.VocabularyDeck), args.Error(1)
}

func (m *MockVocabularyDeckRepo) Create(ctx context.Context, deck *models.VocabularyDeck) error {
	args := m.Called(ctx, deck)
	return args.Error(0)
}

func (m *MockVocabularyDeckRepo) Update(ctx context.Context, deck *models.VocabularyDeck) error {
	args := m.Called(ctx, deck)
	return args.Error(0)
}

func (m *MockVocabularyDeckRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
