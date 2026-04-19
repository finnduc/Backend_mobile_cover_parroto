package services

import (
	"context"

	"go-cover-parroto/internal/core/database"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/bookmark/dtos/res"
	"go-cover-parroto/internal/modules/bookmark/repositories"
)

type IBookmarkService interface {
	AddBookmark(ctx context.Context, userID, lessonID uint) *response.AppError
	RemoveBookmark(ctx context.Context, userID, lessonID uint) *response.AppError
	ListByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResponse[res.BookmarkRes], *response.AppError)
}

type bookmarkService struct {
	repo repositories.IBookmarkRepo
}

func NewBookmarkService(repo repositories.IBookmarkRepo) IBookmarkService {
	return &bookmarkService{repo: repo}
}

func (s *bookmarkService) AddBookmark(ctx context.Context, userID, lessonID uint) *response.AppError {
	if err := s.repo.Create(ctx, userID, lessonID); err != nil {
		return response.Internal("failed to add bookmark")
	}
	return nil
}

func (s *bookmarkService) RemoveBookmark(ctx context.Context, userID, lessonID uint) *response.AppError {
	if err := s.repo.Delete(ctx, userID, lessonID); err != nil {
		return response.Internal("failed to remove bookmark")
	}
	return nil
}

func (s *bookmarkService) ListByUser(ctx context.Context, userID uint, query *database.Query) (*response.PaginatedResponse[res.BookmarkRes], *response.AppError) {
	result, err := s.repo.ListByUser(ctx, userID, query)
	if err != nil {
		return nil, response.Internal("failed to list bookmarks")
	}

	bookmarks := make([]res.BookmarkRes, len(result.Data))
	for i, b := range result.Data {
		bookmarks[i] = mapToBookmarkRes(b)
	}

	return &response.PaginatedResponse[res.BookmarkRes]{
		Data: bookmarks,
		Meta: result.Meta,
	}, nil
}

func mapToBookmarkRes(b *models.Bookmark) res.BookmarkRes {
	r := res.BookmarkRes{
		UserID:    b.UserID,
		LessonID:  b.LessonID,
		CreatedAt: b.CreatedAt,
	}
	if b.Lesson != nil {
		r.Lesson = &res.LessonInfo{
			ID:           b.Lesson.ID,
			Title:        b.Lesson.Title,
			ThumbnailURL: b.Lesson.ThumbnailURL,
			Level:        b.Lesson.Level,
			Duration:     b.Lesson.Duration,
		}
	}
	return r
}
