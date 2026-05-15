package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/bookmark/dtos/req"
	"go-cover-parroto/internal/modules/bookmark/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func testCtx(userID string, role enums.UserRole) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)
	return ctx
}

func TestBookmarkService_AddBookmark(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		body     req.AddBookmarkReq
		setup    func(*repositories.MockBookmarkRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			ctx:  testCtx("1", enums.UserRoleUser),
			body: req.AddBookmarkReq{LessonID: 10},
			setup: func(r *repositories.MockBookmarkRepo) {
				r.On("Create", mock.Anything, "1", uint(10)).Return(nil)
			},
		},
		{
			name:    "unauthenticated returns 401",
			ctx:     context.Background(),
			body:    req.AddBookmarkReq{LessonID: 10},
			setup:   func(r *repositories.MockBookmarkRepo) {},
			wantErr: true,
			wantCode: http.StatusUnauthorized,
		},
		{
			name: "db error returns 500",
			ctx:  testCtx("1", enums.UserRoleUser),
			body: req.AddBookmarkReq{LessonID: 10},
			setup: func(r *repositories.MockBookmarkRepo) {
				r.On("Create", mock.Anything, "1", uint(10)).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockBookmarkRepo)
			tt.setup(mockRepo)
			svc := NewBookmarkService(mockRepo)

			err := svc.AddBookmark(tt.ctx, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBookmarkService_RemoveBookmark(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		body     req.RemoveBookmarkReq
		setup    func(*repositories.MockBookmarkRepo)
		wantErr  bool
		wantCode int
	}{
		{
			name: "success",
			ctx:  testCtx("1", enums.UserRoleUser),
			body: req.RemoveBookmarkReq{LessonID: 10},
			setup: func(r *repositories.MockBookmarkRepo) {
				r.On("Delete", mock.Anything, "1", uint(10)).Return(nil)
			},
		},
		{
			name: "db error returns 500",
			ctx:  testCtx("1", enums.UserRoleUser),
			body: req.RemoveBookmarkReq{LessonID: 10},
			setup: func(r *repositories.MockBookmarkRepo) {
				r.On("Delete", mock.Anything, "1", uint(10)).Return(errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockBookmarkRepo)
			tt.setup(mockRepo)
			svc := NewBookmarkService(mockRepo)

			err := svc.RemoveBookmark(tt.ctx, tt.body)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBookmarkService_List(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		query    req.ListBookmarkQuery
		setup    func(*repositories.MockBookmarkRepo)
		wantLen  int
		wantErr  bool
		wantCode int
	}{
		{
			name:  "success",
			ctx:   testCtx("1", enums.UserRoleUser),
			query: req.ListBookmarkQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockBookmarkRepo) {
				paginated := &response.PaginatedResult[*models.Bookmark]{
					Data: []*models.Bookmark{
						{UserID: "1", LessonID: 10, Lesson: &models.Lesson{ID: 10, Title: "Test", ThumbnailURL: "u", Level: "A1", Duration: 5.0}},
					},
					Meta: response.NewMeta(1, 10, 1),
				}
				r.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)
			},
			wantLen: 1,
		},
		{
			name:  "db error returns 500",
			ctx:   testCtx("1", enums.UserRoleUser),
			query: req.ListBookmarkQuery{Page: 1, Limit: 10},
			setup: func(r *repositories.MockBookmarkRepo) {
				r.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
			},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repositories.MockBookmarkRepo)
			tt.setup(mockRepo)
			svc := NewBookmarkService(mockRepo)

			result, err := svc.List(tt.ctx, tt.query)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.Len(t, result.Data, tt.wantLen)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
