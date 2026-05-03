package services

import (
	"context"

	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/bookmark/dtos/req"
	"go-cover-parroto/internal/modules/bookmark/dtos/res"
	"go-cover-parroto/internal/modules/bookmark/repositories"
	"go-cover-parroto/internal/utils"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "bookmark")
}

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

	log := sLog().With("userId", userID, "lessonId", body.LessonID)
	log.Infow("adding bookmark")
	if createErr := s.repo.Create(ctx, userID, body.LessonID); createErr != nil {
		log.Errorw("failed to add bookmark", "error", createErr)
		return response.Internal("failed to add bookmark")
	}
	log.Infow("bookmark added")
	return nil
}

func (s *bookmarkService) RemoveBookmark(ctx context.Context, body req.RemoveBookmarkReq) *response.AppError {
	userID, err := policy.GetUserID(ctx)
	if err != nil {
		return err
	}

	log := sLog().With("userId", userID, "lessonId", body.LessonID)
	log.Infow("removing bookmark")
	if delErr := s.repo.Delete(ctx, userID, body.LessonID); delErr != nil {
		log.Errorw("failed to remove bookmark", "error", delErr)
		return response.Internal("failed to remove bookmark")
	}
	log.Infow("bookmark removed")
	return nil
}

func (s *bookmarkService) List(ctx context.Context, query req.ListBookmarkQuery) (*response.PaginatedResponse[res.BookmarkRes], *response.AppError) {
	log := sLog()
	log.Infow("listing bookmarks")
	if query.UserID != nil {
		if err := policy.Allow(ctx, *query.UserID); err != nil {
			return nil, err
		}
	}

	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list bookmarks", "error", err)
		return nil, response.Internal("failed to list bookmarks")
	}

	var bookmarks []res.BookmarkRes
	if err := utils.MapToDTOs(result.Data, &bookmarks); err != nil {
		log.Errorw("failed to map bookmarks", "error", err)
		return nil, response.Internal("failed to map bookmarks")
	}

	return &response.PaginatedResponse[res.BookmarkRes]{
		Data: bookmarks,
		Meta: result.Meta,
	}, nil
}
