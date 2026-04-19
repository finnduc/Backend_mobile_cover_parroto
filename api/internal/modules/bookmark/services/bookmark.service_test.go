package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBookmarkRepo struct{ mock.Mock }

func (m *mockBookmarkRepo) Create(ctx context.Context, userID, lessonID uint) error {
	args := m.Called(ctx, userID, lessonID)
	return args.Error(0)
}

func (m *mockBookmarkRepo) Delete(ctx context.Context, userID, lessonID uint) error {
	args := m.Called(ctx, userID, lessonID)
	return args.Error(0)
}

func (m *mockBookmarkRepo) ListByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResult[*models.Bookmark], error) {
	args := m.Called(ctx, userID, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.Bookmark]), args.Error(1)
}

func TestAddBookmark_Success(t *testing.T) {
	repo := new(mockBookmarkRepo)
	svc := NewBookmarkService(repo)

	repo.On("Create", mock.Anything, uint(1), uint(10)).Return(nil)

	appErr := svc.AddBookmark(context.Background(), 1, 10)

	assert.Nil(t, appErr)
	repo.AssertExpectations(t)
}

func TestAddBookmark_Error(t *testing.T) {
	repo := new(mockBookmarkRepo)
	svc := NewBookmarkService(repo)

	repo.On("Create", mock.Anything, uint(1), uint(10)).Return(errors.New("db error"))

	appErr := svc.AddBookmark(context.Background(), 1, 10)

	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}

func TestRemoveBookmark_Success(t *testing.T) {
	repo := new(mockBookmarkRepo)
	svc := NewBookmarkService(repo)

	repo.On("Delete", mock.Anything, uint(1), uint(10)).Return(nil)

	appErr := svc.RemoveBookmark(context.Background(), 1, 10)

	assert.Nil(t, appErr)
	repo.AssertExpectations(t)
}

func TestRemoveBookmark_Error(t *testing.T) {
	repo := new(mockBookmarkRepo)
	svc := NewBookmarkService(repo)

	repo.On("Delete", mock.Anything, uint(1), uint(10)).Return(errors.New("db error"))

	appErr := svc.RemoveBookmark(context.Background(), 1, 10)

	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}

func TestListByUser_Success(t *testing.T) {
	repo := new(mockBookmarkRepo)
	svc := NewBookmarkService(repo)

	lesson := &models.Lesson{ID: 5, Title: "Go Basics", ThumbnailURL: "https://thumb.jpg", Level: "beginner", Duration: 30.5}
	bookmarks := []*models.Bookmark{
		{UserID: 1, LessonID: 5, CreatedAt: time.Now(), Lesson: lesson},
	}
	query := database.NewQuery().SetPage(1).SetLimit(10)
	paginatedResult := &response.PaginatedResult[*models.Bookmark]{
		Data: bookmarks,
		Meta: response.NewMeta(1, 10, 1),
	}
	repo.On("ListByUser", mock.Anything, uint(1), query).Return(paginatedResult, nil)

	result, appErr := svc.ListByUser(context.Background(), 1, query)

	assert.Nil(t, appErr)
	assert.Len(t, result.Data, 1)
	assert.Equal(t, uint(5), result.Data[0].LessonID)
	assert.NotNil(t, result.Data[0].Lesson)
	assert.Equal(t, "Go Basics", result.Data[0].Lesson.Title)
	repo.AssertExpectations(t)
}

func TestListByUser_Error(t *testing.T) {
	repo := new(mockBookmarkRepo)
	svc := NewBookmarkService(repo)

	query := database.NewQuery()
	repo.On("ListByUser", mock.Anything, uint(1), query).Return(nil, errors.New("db error"))

	result, appErr := svc.ListByUser(context.Background(), 1, query)

	assert.Nil(t, result)
	assert.NotNil(t, appErr)
	assert.Equal(t, 500, appErr.Code)
}
