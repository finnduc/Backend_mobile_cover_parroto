package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	db_repos "go-cover-parroto/internal/database/repositories"
	"go-cover-parroto/internal/database/transaction"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/transcript/dtos/req"
	"go-cover-parroto/internal/modules/transcript/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUnitOfWork struct {
	provider transaction.IProvider
}

func (m *mockUnitOfWork) Do(ctx context.Context, fn func(ctx context.Context, p transaction.IProvider) error) error {
	return fn(ctx, m.provider)
}

type mockProvider struct {
	transcriptRepo db_repos.ITranscriptRepo
}

func (m *mockProvider) Auth() db_repos.IAuthRepo                       { panic("not implemented") }
func (m *mockProvider) Bookmark() db_repos.IBookmarkRepo               { panic("not implemented") }
func (m *mockProvider) Category() db_repos.ICategoryRepo               { panic("not implemented") }
func (m *mockProvider) LearningHistory() db_repos.ILearningHistoryRepo { panic("not implemented") }
func (m *mockProvider) Lesson() db_repos.ILessonRepo                   { panic("not implemented") }
func (m *mockProvider) Transcript() db_repos.ITranscriptRepo           { return m.transcriptRepo }

func TestTranscriptService_GetByLesson(t *testing.T) {
	tests := []struct {
		name     string
		lessonID uint
		setup    func(*repositories.MockTranscriptRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:     "success",
			lessonID: 2,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByLesson", mock.Anything, uint(2)).Return([]*models.Transcript{
					{ID: 1, LessonID: 2, Sequence: 1, Content: "Hello world", Phonetic: "ˈhɛloʊ wɜːld", Vietnamese: "Xin chào thế giới", StartTimestamp: 0.0, EndTimestamp: 2.5},
					{ID: 2, LessonID: 2, Sequence: 2, Content: "How are you", Phonetic: "haʊ ɑːr juː", Vietnamese: "Bạn có khỏe không", StartTimestamp: 3.0, EndTimestamp: 5.0},
				}, nil)
			},
			wantLen: 2,
		},
		{
			name:     "empty result",
			lessonID: 1,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByLesson", mock.Anything, uint(1)).Return([]*models.Transcript{}, nil)
			},
			wantLen: 0,
		},
		{
			name:     "db error returns 500",
			lessonID: 99,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByLesson", mock.Anything, uint(99)).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(repositories.MockTranscriptRepo)
			tt.setup(repo)
			svc := NewTranscriptService(repo, nil)

			result, appErr := svc.GetByLesson(context.Background(), tt.lessonID)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Len(t, result, tt.wantLen)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestTranscriptService_ReplaceByLesson(t *testing.T) {
	tests := []struct {
		name     string
		lessonID uint
		body     req.BulkCreateTranscriptReq
		setup    func(*repositories.MockTranscriptRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:     "success",
			lessonID: 5,
			body: req.BulkCreateTranscriptReq{
				Transcripts: []req.BulkCreateTranscriptItem{
					{Sequence: 1, Content: "New content", Phonetic: "nuː", Vietnamese: "Mới", StartTimestamp: 0.0, EndTimestamp: 2.0},
					{Sequence: 2, Content: "Next line", Phonetic: "nɛkst", Vietnamese: "Tiếp theo", StartTimestamp: 2.5, EndTimestamp: 4.0},
				},
			},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("DeleteByLesson", mock.Anything, uint(5)).Return(nil)
				r.On("BulkCreate", mock.Anything, mock.Anything).Return(nil)
			},
			wantLen: 2,
		},
		{
			name:     "delete error returns 500",
			lessonID: 5,
			body: req.BulkCreateTranscriptReq{
				Transcripts: []req.BulkCreateTranscriptItem{
					{Sequence: 1, Content: "x", Vietnamese: "y"},
				},
			},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("DeleteByLesson", mock.Anything, uint(5)).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "bulk create error returns 500",
			lessonID: 5,
			body: req.BulkCreateTranscriptReq{
				Transcripts: []req.BulkCreateTranscriptItem{
					{Sequence: 1, Content: "x", Vietnamese: "y"},
				},
			},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("DeleteByLesson", mock.Anything, uint(5)).Return(nil)
				r.On("BulkCreate", mock.Anything, mock.Anything).Return(errors.New("bulk create failed"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tRepo := new(repositories.MockTranscriptRepo)
			uow := &mockUnitOfWork{provider: &mockProvider{transcriptRepo: tRepo}}
			tt.setup(tRepo)
			svc := NewTranscriptService(tRepo, uow)

			result, appErr := svc.ReplaceByLesson(context.Background(), tt.lessonID, tt.body)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Len(t, result, tt.wantLen)
			}
			tRepo.AssertExpectations(t)
		})
	}
}
