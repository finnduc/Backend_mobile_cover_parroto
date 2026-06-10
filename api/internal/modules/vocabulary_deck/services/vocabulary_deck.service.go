package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go-cover-parroto/internal/modules/vocabulary_deck/dtos/req"

	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "vocabulary_deck")
}

type IVocabularyDeckService interface {
	List(ctx context.Context, query req.ListVocabularyDeckQuery) (*response.PaginatedResult[*models.VocabularyDeck], *response.AppError)
	ListDefault(ctx context.Context, query req.ListVocabularyDeckQuery) (*response.PaginatedResult[*models.VocabularyDeck], *response.AppError)
	ListByUser(ctx context.Context, userID string, query req.ListVocabularyDeckQuery) (*response.PaginatedResult[*models.VocabularyDeck], *response.AppError)
	GetByID(ctx context.Context, id uint) (*models.VocabularyDeck, *response.AppError)
	Create(ctx context.Context, userID string, body req.CreateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
	CreateAsSystem(ctx context.Context, body req.CreateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError)
	UpdateAsSystem(ctx context.Context, id uint, body req.UpdateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError)
	DeleteAsSystem(ctx context.Context, id uint) *response.AppError
}

type vocabularyDeckService struct {
	repo db_repos.IVocabularyDeckRepo
}

func NewVocabularyDeckService(repo db_repos.IVocabularyDeckRepo) IVocabularyDeckService {
	return &vocabularyDeckService{repo: repo}
}

func (s *vocabularyDeckService) List(ctx context.Context, query req.ListVocabularyDeckQuery) (*response.PaginatedResult[*models.VocabularyDeck], *response.AppError) {
	log := sLog()
	log.Infow("listing vocabulary decks")
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list decks", "error", err)
		return nil, response.Internal("failed to list decks")
	}
	return result, nil
}

func (s *vocabularyDeckService) ListDefault(ctx context.Context, query req.ListVocabularyDeckQuery) (*response.PaginatedResult[*models.VocabularyDeck], *response.AppError) {
	log := sLog()
	log.Infow("listing default vocabulary decks")
	isDefault := true
	query.IsDefault = &isDefault
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list default decks", "error", err)
		return nil, response.Internal("failed to list default decks")
	}
	return result, nil
}

func (s *vocabularyDeckService) ListByUser(ctx context.Context, userID string, query req.ListVocabularyDeckQuery) (*response.PaginatedResult[*models.VocabularyDeck], *response.AppError) {
	log := sLog().With("userId", userID)
	log.Infow("listing user vocabulary decks")
	dbQuery := query.ToQuery()
	dbQuery.SetFilter("user_id", userID)
	result, err := s.repo.FindAll(ctx, dbQuery)
	if err != nil {
		log.Errorw("failed to list user decks", "error", err)
		return nil, response.Internal("failed to list user decks")
	}
	return result, nil
}

func (s *vocabularyDeckService) GetByID(ctx context.Context, id uint) (*models.VocabularyDeck, *response.AppError) {
	log := sLog().With("id", id)
	log.Infow("getting vocabulary deck")
	deck, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("deck not found")
		}
		log.Errorw("failed to get deck", "error", err)
		return nil, response.Internal("failed to get deck")
	}
	return deck, nil
}

func (s *vocabularyDeckService) Create(ctx context.Context, userID string, body req.CreateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError) {

	log := sLog().With("userId", userID)
	log.Infow("creating user vocabulary deck")
	deck := &models.VocabularyDeck{
		Name:         body.Name,
		Description:  body.Description,
		ThumbnailURL: body.ThumbnailURL,
		Level:        body.Level,
		UserID:       &userID,
	}
	if createErr := s.repo.Create(ctx, deck); createErr != nil {
		log.Errorw("failed to create deck", "error", createErr)
		return nil, response.Internal("failed to create deck")
	}
	log.Infow("deck created", "id", deck.ID)
	return deck, nil
}

func (s *vocabularyDeckService) Update(ctx context.Context, id uint, body req.UpdateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError) {
	log := sLog().With("id", id)
	log.Infow("updating user vocabulary deck")
	deck, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("deck not found")
		}
		return nil, response.Internal("failed to get deck")
	}

	if deck.UserID != nil {
		actor, appErr := policy.ActorFromContext(ctx)
		if appErr != nil {
			return nil, appErr
		}
		if appErr := policy.CanMutate(actor, *deck.UserID); appErr != nil {
			return nil, appErr
		}
	}

	deck.Name = body.Name
	deck.Description = body.Description
	deck.ThumbnailURL = body.ThumbnailURL
	deck.Level = body.Level
	if updateErr := s.repo.Update(ctx, deck); updateErr != nil {
		log.Errorw("failed to update deck", "error", updateErr)
		return nil, response.Internal("failed to update deck")
	}
	log.Infow("deck updated")
	return deck, nil
}

func (s *vocabularyDeckService) Delete(ctx context.Context, id uint) *response.AppError {
	log := sLog().With("id", id)
	log.Infow("deleting user vocabulary deck")
	deck, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return response.NotFound("deck not found")
		}
		return response.Internal("failed to get deck")
	}

	if deck.UserID != nil {
		actor, appErr := policy.ActorFromContext(ctx)
		if appErr != nil {
			return appErr
		}
		if appErr := policy.CanMutate(actor, *deck.UserID); appErr != nil {
			return appErr
		}
	}

	if delErr := s.repo.Delete(ctx, id); delErr != nil {
		log.Errorw("failed to delete deck", "error", delErr)
		return response.Internal("failed to delete deck")
	}
	log.Infow("deck deleted")
	return nil
}

func (s *vocabularyDeckService) CreateAsSystem(ctx context.Context, body req.CreateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError) {
	if _, err := policy.ActorFromContext(ctx); err != nil {
		return nil, err
	}
	log := sLog()
	log.Infow("creating system vocabulary deck")
	deck := &models.VocabularyDeck{
		Name:         body.Name,
		Description:  body.Description,
		ThumbnailURL: body.ThumbnailURL,
		Level:        body.Level,
		IsDefault:    true,
	}
	if createErr := s.repo.Create(ctx, deck); createErr != nil {
		log.Errorw("failed to create system deck", "error", createErr)
		return nil, response.Internal("failed to create deck")
	}
	log.Infow("system deck created", "id", deck.ID)
	return deck, nil
}

func (s *vocabularyDeckService) UpdateAsSystem(ctx context.Context, id uint, body req.UpdateVocabularyDeckReq) (*models.VocabularyDeck, *response.AppError) {
	if _, err := policy.ActorFromContext(ctx); err != nil {
		return nil, err
	}
	log := sLog().With("id", id)
	log.Infow("updating system vocabulary deck")
	deck, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("deck not found")
		}
		return nil, response.Internal("failed to get deck")
	}

	deck.Name = body.Name
	deck.Description = body.Description
	deck.ThumbnailURL = body.ThumbnailURL
	deck.Level = body.Level
	if updateErr := s.repo.Update(ctx, deck); updateErr != nil {
		log.Errorw("failed to update deck", "error", updateErr)
		return nil, response.Internal("failed to update deck")
	}
	log.Infow("deck updated")
	return deck, nil
}

func (s *vocabularyDeckService) DeleteAsSystem(ctx context.Context, id uint) *response.AppError {
	if _, err := policy.ActorFromContext(ctx); err != nil {
		return err
	}
	log := sLog().With("id", id)
	log.Infow("deleting system vocabulary deck")
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return response.NotFound("deck not found")
		}
		return response.Internal("failed to get deck")
	}
	if delErr := s.repo.Delete(ctx, id); delErr != nil {
		log.Errorw("failed to delete deck", "error", delErr)
		return response.Internal("failed to delete deck")
	}
	log.Infow("deck deleted")
	return nil
}
