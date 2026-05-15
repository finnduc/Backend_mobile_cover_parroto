package repositories

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"

	"github.com/stretchr/testify/mock"
)

var _ db_repos.ILessonRepo = (*MockLessonRepo)(nil)

type MockLessonRepo struct {
	mock.Mock
}

func (m *MockLessonRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Lesson], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.Lesson]), args.Error(1)
}

func (m *MockLessonRepo) FindByID(ctx context.Context, id uint) (*models.Lesson, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Lesson), args.Error(1)
}

func (m *MockLessonRepo) Create(ctx context.Context, lesson *models.Lesson) error {
	args := m.Called(ctx, lesson)
	return args.Error(0)
}

func (m *MockLessonRepo) Update(ctx context.Context, lesson *models.Lesson) error {
	args := m.Called(ctx, lesson)
	return args.Error(0)
}

func (m *MockLessonRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
