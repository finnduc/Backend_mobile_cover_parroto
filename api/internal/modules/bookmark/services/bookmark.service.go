package services

import (
	"context"

	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/bookmark/dtos/req"
	"go-cover-parroto/internal/modules/bookmark/dtos/res"
	"go-cover-parroto/internal/modules/bookmark/repositories"
)

type IBookmarkService interface {
	AddBookmark(ctx context.Context, body req.AddBookmarkReq) *response.AppError
	RemoveBookmark(ctx context.Context, body req.RemoveBookmarkReq) *response.AppError
	List(ctx context.Context, query req.ListBookmarkQuery) (*response.PaginatedResponse[res.BookmarkRes], *response.AppError)
}

type bookmarkService struct {
	repo repositories.IBookmarkRepo
}

func NewBookmarkService(repo repositories.IBookmarkRepo) IBookmarkService {
	return &bookmarkService{repo: repo}
}

func (s *bookmarkService) AddBookmark(ctx context.Context, body req.AddBookmarkReq) *response.AppError {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return err
	}

	if createErr := s.repo.Create(ctx, userID, body.LessonID); createErr != nil {
		return response.Internal("failed to add bookmark")
	}
	return nil
}

func (s *bookmarkService) RemoveBookmark(ctx context.Context, body req.RemoveBookmarkReq) *response.AppError {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return err
	}

	if delErr := s.repo.Delete(ctx, userID, body.LessonID); delErr != nil {
		return response.Internal("failed to remove bookmark")
	}
	return nil
}

func (s *bookmarkService) List(ctx context.Context, query req.ListBookmarkQuery) (*response.PaginatedResponse[res.BookmarkRes], *response.AppError) {
	result, err := s.repo.FindAll(ctx, query.ToQuery())
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
