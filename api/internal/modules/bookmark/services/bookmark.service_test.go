package services

import (
	"context"
	"errors"
	"testing"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/bookmark/dtos/req"
	"github.com/stretchr/testify/mock"
)

type mockBookmarkRepo struct {
	mock.Mock
}

func (m *mockBookmarkRepo) Create(ctx context.Context, userID string, lessonID uint) error {
	args := m.Called(ctx, userID, lessonID)
	return args.Error(0)
}

func (m *mockBookmarkRepo) Delete(ctx context.Context, userID string, lessonID uint) error {
	args := m.Called(ctx, userID, lessonID)
	return args.Error(0)
}

func (m *mockBookmarkRepo) FindAll(ctx context.Context, query *database.Query) (*response.PaginatedResult[*models.Bookmark], error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.PaginatedResult[*models.Bookmark]), args.Error(1)
}

func testCtx(userID string, role enums.UserRole) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)
	return ctx
}

func TestAddBookmark_Success(t *testing.T) {
	mockRepo := new(mockBookmarkRepo)
	svc := NewBookmarkService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := req.AddBookmarkReq{LessonID: 10}

	mockRepo.On("Create", mock.Anything, "1", uint(10)).Return(nil)

	err := svc.AddBookmark(ctx, body)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	mockRepo.AssertExpectations(t)
}

func TestAddBookmark_Error(t *testing.T) {
	mockRepo := new(mockBookmarkRepo)
	svc := NewBookmarkService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := req.AddBookmarkReq{LessonID: 10}

	mockRepo.On("Create", mock.Anything, "1", uint(10)).Return(errors.New("db error"))

	err := svc.AddBookmark(ctx, body)
	if err == nil {
		t.Error("expected error, got nil")
	}
	mockRepo.AssertExpectations(t)
}

func TestAddBookmark_Unauthenticated(t *testing.T) {
	mockRepo := new(mockBookmarkRepo)
	svc := NewBookmarkService(mockRepo)
	ctx := context.Background()

	err := svc.AddBookmark(ctx, req.AddBookmarkReq{LessonID: 10})
	if err == nil || err.Code != 401 {
		t.Errorf("expected 401 unauthorized, got %v", err)
	}
}

func TestRemoveBookmark_Success(t *testing.T) {
	mockRepo := new(mockBookmarkRepo)
	svc := NewBookmarkService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := req.RemoveBookmarkReq{LessonID: 10}

	mockRepo.On("Delete", mock.Anything, "1", uint(10)).Return(nil)

	err := svc.RemoveBookmark(ctx, body)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	mockRepo.AssertExpectations(t)
}

func TestRemoveBookmark_Error(t *testing.T) {
	mockRepo := new(mockBookmarkRepo)
	svc := NewBookmarkService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	body := req.RemoveBookmarkReq{LessonID: 10}

	mockRepo.On("Delete", mock.Anything, "1", uint(10)).Return(errors.New("db error"))

	err := svc.RemoveBookmark(ctx, body)
	if err == nil {
		t.Error("expected error, got nil")
	}
	mockRepo.AssertExpectations(t)
}

func TestList_Success(t *testing.T) {
	mockRepo := new(mockBookmarkRepo)
	svc := NewBookmarkService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	q := req.ListBookmarkQuery{Page: 1, Limit: 10}

	paginated := &response.PaginatedResult[*models.Bookmark]{
		Data: []*models.Bookmark{
			{UserID: "1", LessonID: 10, Lesson: &models.Lesson{ID: 10, Title: "Test", ThumbnailURL: "u", Level: "A1", Duration: 5.0}},
		},
		Meta: response.NewMeta(1, 10, 1),
	}

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(paginated, nil)

	result, err := svc.List(ctx, q)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result.Data) != 1 {
		t.Errorf("expected 1 bookmark, got %d", len(result.Data))
	}
	mockRepo.AssertExpectations(t)
}

func TestList_Error(t *testing.T) {
	mockRepo := new(mockBookmarkRepo)
	svc := NewBookmarkService(mockRepo)
	ctx := testCtx("1", enums.UserRoleUser)
	q := req.ListBookmarkQuery{Page: 1, Limit: 10}

	mockRepo.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	_, err := svc.List(ctx, q)
	if err == nil {
		t.Error("expected error, got nil")
	}
	mockRepo.AssertExpectations(t)
}
