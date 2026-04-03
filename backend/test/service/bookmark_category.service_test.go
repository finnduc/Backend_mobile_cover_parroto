package service

import (
	"context"
	"go-cover-parroto/internal/mocks"
	"go-cover-parroto/internal/models"
	"go-cover-parroto/internal/service"
	"go-cover-parroto/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookmarkToggle(t *testing.T) {
	mockRepo := new(mocks.MockBookmarkRepo)
	s := service.NewBookmarkService(mockRepo)

	userID := uint(1)
	lessonID := uint(100)

	mockRepo.On("Toggle", mock.Anything, userID, lessonID).Return(true, nil)

	bookmarked, err := s.Toggle(context.Background(), userID, lessonID)

	assert.NoError(t, err)
	assert.True(t, bookmarked)
	mockRepo.AssertExpectations(t)
}

func TestListBookmarks(t *testing.T) {
	mockRepo := new(mocks.MockBookmarkRepo)
	s := service.NewBookmarkService(mockRepo)

	userID := uint(1)
	pagination := utils.PaginationQuery{Page: 1, Limit: 10}
	mockLessons := []models.Lesson{{Base: models.Base{ID: 101}, Title: "Lesson 1"}}

	mockRepo.On("FindByUser", mock.Anything, userID, pagination).Return(mockLessons, int64(1), nil)

	lessons, total, err := s.ListBookmarks(context.Background(), userID, pagination)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, lessons, 1)
	mockRepo.AssertExpectations(t)
}

func TestListCategories(t *testing.T) {
	mockRepo := new(mocks.MockCategoryRepo)
	s := service.NewCategoryService(mockRepo)

	mockCats := []models.Category{{ID: 1, Name: "Grammar"}}
	mockRepo.On("FindAll", mock.Anything).Return(mockCats, nil)

	cats, err := s.ListCategories(context.Background())

	assert.NoError(t, err)
	assert.Len(t, cats, 1)
	assert.Equal(t, "Grammar", cats[0].Name)
	mockRepo.AssertExpectations(t)
}
