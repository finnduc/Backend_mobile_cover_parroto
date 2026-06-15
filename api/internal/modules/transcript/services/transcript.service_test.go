package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	coreError "go-cover-parroto/internal/core/errors"
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
	transaction.IProvider
	transcriptRepo db_repos.ITranscriptRepo
}

func (m *mockProvider) Transcript() db_repos.ITranscriptRepo { return m.transcriptRepo }

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

func TestTranscriptService_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockTranscriptRepo)
		wantSeq  int
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(&models.Transcript{ID: 1, Sequence: 3, Content: "Test"}, nil)
			},
			wantSeq: 3,
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(repositories.MockTranscriptRepo)
			tt.setup(repo)
			svc := NewTranscriptService(repo, nil)

			result, appErr := svc.GetByID(context.Background(), tt.id)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Equal(t, tt.wantSeq, result.Sequence)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestTranscriptService_Create(t *testing.T) {
	tests := []struct {
		name     string
		body     req.CreateTranscriptReq
		setup    func(*repositories.MockTranscriptRepo)
		wantSeq  int
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			body: req.CreateTranscriptReq{LessonID: 1, Sequence: 5, Content: "New transcript"},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("Create", mock.Anything, mock.MatchedBy(func(t *models.Transcript) bool {
					return t.Sequence == 5 && t.Content == "New transcript"
				})).Return(nil)
			},
			wantSeq: 5,
		},
		{
			name: "db error returns 500",
			body: req.CreateTranscriptReq{LessonID: 1, Sequence: 1, Content: "x"},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("Create", mock.Anything, mock.Anything).Return(errors.New("db error"))
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

			result, appErr := svc.Create(context.Background(), tt.body)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Equal(t, tt.wantSeq, result.Sequence)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestTranscriptService_BulkCreate(t *testing.T) {
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
			lessonID: 1,
			body: req.BulkCreateTranscriptReq{
				Transcripts: []req.BulkCreateTranscriptItem{
					{Sequence: 1, Content: "First"},
					{Sequence: 2, Content: "Second"},
				},
			},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("BulkCreate", mock.Anything, mock.Anything).Return(nil)
			},
			wantLen: 2,
		},
		{
			name:     "db error returns 500",
			lessonID: 1,
			body: req.BulkCreateTranscriptReq{
				Transcripts: []req.BulkCreateTranscriptItem{
					{Sequence: 1, Content: "x"},
				},
			},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("BulkCreate", mock.Anything, mock.Anything).Return(errors.New("db error"))
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

			result, appErr := svc.BulkCreate(context.Background(), tt.lessonID, tt.body)

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

func TestTranscriptService_Update(t *testing.T) {
	existing := &models.Transcript{ID: 1, Sequence: 1, Content: "Original"}
	seq := 2
	content := "Updated"
	tests := []struct {
		name     string
		id       uint
		body     req.UpdateTranscriptReq
		setup    func(*repositories.MockTranscriptRepo)
		wantSeq  int
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			body: req.UpdateTranscriptReq{Sequence: &seq, Content: &content},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.MatchedBy(func(t *models.Transcript) bool {
					return t.Sequence == 2 && t.Content == "Updated"
				})).Return(nil)
			},
			wantSeq: 2,
		},
		{
			name: "not found returns 404",
			id:   999,
			body: req.UpdateTranscriptReq{Sequence: &seq},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "update error returns 500",
			id:   1,
			body: req.UpdateTranscriptReq{Sequence: &seq, Content: &content},
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(existing, nil)
				r.On("Update", mock.Anything, mock.Anything).Return(errors.New("db error"))
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

			result, appErr := svc.Update(context.Background(), tt.id, tt.body)

			if tt.wantErr {
				assert.Nil(t, result)
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
				assert.Equal(t, tt.wantSeq, result.Sequence)
			}
			repo.AssertExpectations(t)
		})
	}
}

func TestTranscriptService_Delete(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		setup    func(*repositories.MockTranscriptRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			id:   1,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(1)).Return(&models.Transcript{ID: 1}, nil)
				r.On("Delete", mock.Anything, uint(1)).Return(nil)
			},
		},
		{
			name: "not found returns 404",
			id:   999,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(999)).Return(nil, coreError.ErrNotFound)
			},
			wantErr:  true,
			wantCode: http.StatusNotFound,
		},
		{
			name: "db error returns 500",
			id:   2,
			setup: func(r *repositories.MockTranscriptRepo) {
				r.On("FindByID", mock.Anything, uint(2)).Return(&models.Transcript{ID: 2}, nil)
				r.On("Delete", mock.Anything, uint(2)).Return(errors.New("db error"))
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

			appErr := svc.Delete(context.Background(), tt.id)

			if tt.wantErr {
				assert.NotNil(t, appErr)
				assert.Equal(t, tt.wantCode, appErr.Code)
			} else {
				assert.Nil(t, appErr)
			}
			repo.AssertExpectations(t)
		})
	}
}
