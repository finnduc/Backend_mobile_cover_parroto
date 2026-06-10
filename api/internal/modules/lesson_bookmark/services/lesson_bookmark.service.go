package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/lesson_bookmark/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "lesson_bookmark")
}

type ILessonBookmarkService interface {
	Create(ctx context.Context, userID string, body req.CreateLessonBookmarkReq) (*models.LessonBookmark, *response.AppError)
	List(ctx context.Context, userID string, query req.ListLessonBookmarkQuery) (*response.PaginatedResult[*models.LessonBookmark], *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
	Toggle(ctx context.Context, userID string, lessonID uint) (*models.LessonBookmark, *response.AppError)
}

type lessonBookmarkService struct {
	repo db_repos.ILessonBookmarkRepo
}

func NewLessonBookmarkService(repo db_repos.ILessonBookmarkRepo) ILessonBookmarkService {
	return &lessonBookmarkService{repo: repo}
}

func (s *lessonBookmarkService) Create(ctx context.Context, userID string, body req.CreateLessonBookmarkReq) (*models.LessonBookmark, *response.AppError) {
	log := sLog().With("userId", userID, "lessonId", body.LessonID)
	log.Infow("creating lesson bookmark")

	bookmark := &models.LessonBookmark{
		UserID:   userID,
		LessonID: body.LessonID,
	}

	if createErr := s.repo.Create(ctx, bookmark); createErr != nil {
		log.Errorw("failed to create lesson bookmark", "error", createErr)
		return nil, response.Internal("failed to create lesson bookmark")
	}

	log.Infow("lesson bookmark created")
	return bookmark, nil
}

func (s *lessonBookmarkService) List(ctx context.Context, userID string, query req.ListLessonBookmarkQuery) (*response.PaginatedResult[*models.LessonBookmark], *response.AppError) {
	log := sLog().With("userId", userID)
	log.Infow("listing lesson bookmarks")

	q := query.ToQuery()
	q.SetFilter("user_id", userID)

	result, findErr := s.repo.FindAll(ctx, q)
	if findErr != nil {
		log.Errorw("failed to list lesson bookmarks", "error", findErr)
		return nil, response.Internal("failed to list lesson bookmarks")
	}
	return result, nil
}

func (s *lessonBookmarkService) Delete(ctx context.Context, id uint) *response.AppError {
	actor, err := policy.ActorFromContext(ctx)
	if err != nil {
		return err
	}

	log := sLog().With("userId", actor.UserID, "id", id)
	log.Infow("deleting lesson bookmark")

	bookmark, findErr := s.repo.FindByID(ctx, id)
	if findErr != nil {
		if errors.Is(findErr, coreError.ErrNotFound) {
			return response.NotFound("lesson bookmark not found")
		}
		log.Errorw("failed to find lesson bookmark", "error", findErr)
		return response.Internal("failed to find lesson bookmark")
	}

	if accessErr := policy.CanMutate(actor, bookmark.UserID); accessErr != nil {
		return accessErr
	}

	if delErr := s.repo.Delete(ctx, id); delErr != nil {
		log.Errorw("failed to delete lesson bookmark", "error", delErr)
		return response.Internal("failed to delete lesson bookmark")
	}

	log.Infow("lesson bookmark deleted")
	return nil
}

func (s *lessonBookmarkService) Toggle(ctx context.Context, userID string, lessonID uint) (*models.LessonBookmark, *response.AppError) {
	log := sLog().With("userId", userID, "lessonId", lessonID)

	existing, findErr := s.repo.FindByUserAndLesson(ctx, userID, lessonID)
	if findErr == nil && existing != nil {
		log.Infow("removing existing lesson bookmark")
		if delErr := s.repo.Delete(ctx, existing.ID); delErr != nil {
			log.Errorw("failed to delete lesson bookmark", "error", delErr)
			return nil, response.Internal("failed to delete lesson bookmark")
		}
		return nil, nil
	}

	log.Infow("creating lesson bookmark")
	bookmark := &models.LessonBookmark{
		UserID:   userID,
		LessonID: lessonID,
	}

	if createErr := s.repo.Create(ctx, bookmark); createErr != nil {
		log.Errorw("failed to create lesson bookmark", "error", createErr)
		return nil, response.Internal("failed to create lesson bookmark")
	}

	return bookmark, nil
}
