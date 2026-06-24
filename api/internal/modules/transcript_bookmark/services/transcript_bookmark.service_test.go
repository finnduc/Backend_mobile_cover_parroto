package services

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/transcript_bookmark/dtos/req"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTranscriptBookmarkRepo struct {
	mock.Mock
}

func (m *mockTranscriptBookmarkRepo) Create(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	args := m.Called(ctx, bookmark)
	return args.Error(0)
}

func (m *mockTranscriptBookmarkRepo) Update(ctx context.Context, bookmark *models.TranscriptBookmark) error {
	args := m.Called(ctx, bookmark)
	return args.Error(0)
}

func (m *mockTranscriptBookmarkRepo) FindByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) (*models.TranscriptBookmark, error) {
	args := m.Called(ctx, userID, transcriptID)
	if bookmark, ok := args.Get(0).(*models.TranscriptBookmark); ok {
		return bookmark, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTranscriptBookmarkRepo) FindAllByUser(ctx context.Context, userID string, lessonID *uint) ([]*models.TranscriptBookmark, error) {
	args := m.Called(ctx, userID, lessonID)
	if bookmarks, ok := args.Get(0).([]*models.TranscriptBookmark); ok {
		return bookmarks, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTranscriptBookmarkRepo) DeleteByUserAndTranscript(ctx context.Context, userID string, transcriptID uint) error {
	args := m.Called(ctx, userID, transcriptID)
	return args.Error(0)
}

type mockTranscriptRepo struct {
	mock.Mock
}

func (m *mockTranscriptRepo) FindByLesson(ctx context.Context, lessonID uint) ([]*models.Transcript, error) {
	args := m.Called(ctx, lessonID)
	if transcripts, ok := args.Get(0).([]*models.Transcript); ok {
		return transcripts, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTranscriptRepo) FindByID(ctx context.Context, id uint) (*models.Transcript, error) {
	args := m.Called(ctx, id)
	if transcript, ok := args.Get(0).(*models.Transcript); ok {
		return transcript, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTranscriptRepo) Create(ctx context.Context, transcript *models.Transcript) error {
	args := m.Called(ctx, transcript)
	return args.Error(0)
}

func (m *mockTranscriptRepo) BulkCreate(ctx context.Context, transcripts []*models.Transcript) error {
	args := m.Called(ctx, transcripts)
	return args.Error(0)
}

func (m *mockTranscriptRepo) Update(ctx context.Context, transcript *models.Transcript) error {
	args := m.Called(ctx, transcript)
	return args.Error(0)
}

func (m *mockTranscriptRepo) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTranscriptRepo) DeleteByLesson(ctx context.Context, lessonID uint) error {
	args := m.Called(ctx, lessonID)
	return args.Error(0)
}

func TestTranscriptBookmarkService_ListGroupsByLesson(t *testing.T) {
	createdAt := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	lessonID := uint(10)
	repo := new(mockTranscriptBookmarkRepo)
	transcriptRepo := new(mockTranscriptRepo)
	repo.On("FindAllByUser", mock.Anything, "user1", &lessonID).Return([]*models.TranscriptBookmark{
		{
			UserID:       "user1",
			LessonID:     10,
			TranscriptID: 1,
			Note:         "first",
			CreatedAt:    createdAt,
			Lesson:       &models.Lesson{Title: "Lesson A"},
			Transcript:   &models.Transcript{Content: "hello", Phonetic: "həˈloʊ", Vietnamese: "xin chao"},
		},
		{
			UserID:       "user1",
			LessonID:     10,
			TranscriptID: 2,
			Note:         "second",
			CreatedAt:    createdAt,
			Lesson:       &models.Lesson{Title: "Lesson A"},
			Transcript:   &models.Transcript{Content: "world"},
		},
	}, nil)

	result, err := NewTranscriptBookmarkService(repo, transcriptRepo).List(context.Background(), "user1", req.ListTranscriptBookmarkQuery{LessonID: &lessonID})

	assert.Nil(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(10), result[0].LessonID)
	assert.Equal(t, "Lesson A", result[0].LessonTitle)
	assert.Len(t, result[0].Transcripts, 2)
	assert.Equal(t, uint(1), result[0].Transcripts[0].TranscriptID)
	assert.Equal(t, "hello", result[0].Transcripts[0].Content)
	assert.Equal(t, createdAt.Format(time.RFC3339), result[0].Transcripts[0].CreatedAt)
	repo.AssertExpectations(t)
}

func TestTranscriptBookmarkService_CreateInfersLessonID(t *testing.T) {
	repo := new(mockTranscriptBookmarkRepo)
	transcriptRepo := new(mockTranscriptRepo)
	transcriptRepo.On("FindByID", mock.Anything, uint(7)).Return(&models.Transcript{ID: 7, LessonID: 20}, nil)
	repo.On("FindByUserAndTranscript", mock.Anything, "user1", uint(7)).Return(nil, coreError.ErrNotFound)
	repo.On("Create", mock.Anything, mock.MatchedBy(func(bookmark *models.TranscriptBookmark) bool {
		return bookmark.UserID == "user1" && bookmark.TranscriptID == 7 && bookmark.LessonID == 20 && bookmark.Note == "note"
	})).Return(nil)

	result, err := NewTranscriptBookmarkService(repo, transcriptRepo).Create(context.Background(), "user1", req.CreateTranscriptBookmarkReq{TranscriptID: 7, Note: "note"})

	assert.Nil(t, err)
	assert.Equal(t, uint(20), result.LessonID)
	assert.Equal(t, uint(7), result.TranscriptID)
	assert.Equal(t, "note", result.Note)
	repo.AssertExpectations(t)
	transcriptRepo.AssertExpectations(t)
}

func TestTranscriptBookmarkService_CreateUpdatesExistingBookmark(t *testing.T) {
	repo := new(mockTranscriptBookmarkRepo)
	transcriptRepo := new(mockTranscriptRepo)
	existing := &models.TranscriptBookmark{UserID: "user1", LessonID: 20, TranscriptID: 7, Note: "old"}
	transcriptRepo.On("FindByID", mock.Anything, uint(7)).Return(&models.Transcript{ID: 7, LessonID: 20}, nil)
	repo.On("FindByUserAndTranscript", mock.Anything, "user1", uint(7)).Return(existing, nil)
	repo.On("Update", mock.Anything, mock.MatchedBy(func(bookmark *models.TranscriptBookmark) bool {
		return bookmark.Note == "new"
	})).Return(nil)

	result, err := NewTranscriptBookmarkService(repo, transcriptRepo).Create(context.Background(), "user1", req.CreateTranscriptBookmarkReq{TranscriptID: 7, Note: "new"})

	assert.Nil(t, err)
	assert.Equal(t, "new", result.Note)
	repo.AssertExpectations(t)
	transcriptRepo.AssertExpectations(t)
}

func TestTranscriptBookmarkService_UpdateNotFound(t *testing.T) {
	repo := new(mockTranscriptBookmarkRepo)
	transcriptRepo := new(mockTranscriptRepo)
	repo.On("FindByUserAndTranscript", mock.Anything, "user1", uint(9)).Return(nil, coreError.ErrNotFound)

	result, err := NewTranscriptBookmarkService(repo, transcriptRepo).Update(context.Background(), "user1", 9, req.UpdateTranscriptBookmarkReq{Note: "note"})

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Code)
	repo.AssertExpectations(t)
}

func TestTranscriptBookmarkService_DeleteReturnsInternalOnRepoError(t *testing.T) {
	repo := new(mockTranscriptBookmarkRepo)
	transcriptRepo := new(mockTranscriptRepo)
	repo.On("DeleteByUserAndTranscript", mock.Anything, "user1", uint(9)).Return(errors.New("db error"))

	err := NewTranscriptBookmarkService(repo, transcriptRepo).Delete(context.Background(), "user1", 9)

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Code)
	repo.AssertExpectations(t)
}
