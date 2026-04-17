package services

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/bookmark/repositories"
)

type IBookmarkService interface {
	AddBookmark(ctx context.Context, userID, lessonID uint) *response.AppError
	RemoveBookmark(ctx context.Context, userID, lessonID uint) *response.AppError
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
