package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/vocabulary_item/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "vocabulary_item")
}

type IVocabularyItemService interface {
	List(ctx context.Context, query req.ListVocabularyItemQuery) (*response.PaginatedResult[*models.VocabularyItem], *response.AppError)
	GetByID(ctx context.Context, id uint) (*models.VocabularyItem, *response.AppError)
	Create(ctx context.Context, body req.CreateVocabularyItemFromDeckReq) (*models.VocabularyItem, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateVocabularyItemReq) (*models.VocabularyItem, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
}

type vocabularyItemService struct {
	repo     db_repos.IVocabularyItemRepo
	deckRepo db_repos.IVocabularyDeckRepo
}

func NewVocabularyItemService(repo db_repos.IVocabularyItemRepo, deckRepo db_repos.IVocabularyDeckRepo) IVocabularyItemService {
	return &vocabularyItemService{repo: repo, deckRepo: deckRepo}
}

func (s *vocabularyItemService) List(ctx context.Context, query req.ListVocabularyItemQuery) (*response.PaginatedResult[*models.VocabularyItem], *response.AppError) {
	log := sLog()
	log.Infow("listing vocabulary items")
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list items", "error", err)
		return nil, response.Internal("failed to list items")
	}
	return result, nil
}

func (s *vocabularyItemService) GetByID(ctx context.Context, id uint) (*models.VocabularyItem, *response.AppError) {
	log := sLog().With("id", id)
	log.Infow("getting vocabulary item")
	item, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("item not found")
		}
		log.Errorw("failed to get item", "error", err)
		return nil, response.Internal("failed to get item")
	}
	return item, nil
}

func (s *vocabularyItemService) Create(ctx context.Context, body req.CreateVocabularyItemFromDeckReq) (*models.VocabularyItem, *response.AppError) {
	actor, err := policy.ActorFromContext(ctx)
	if err != nil {
		return nil, err
	}

	log := sLog().With("userId", actor.UserID, "deckId", body.DeckID)
	log.Infow("creating vocabulary item in user deck")

	deck, deckErr := s.deckRepo.FindByID(ctx, body.DeckID)
	if deckErr != nil {
		if errors.Is(deckErr, coreError.ErrNotFound) {
			return nil, response.NotFound("deck not found")
		}
		return nil, response.Internal("failed to verify deck")
	}

	if deck.UserID != nil {
		if appErr := policy.CanMutate(actor, *deck.UserID); appErr != nil {
			return nil, appErr
		}
	}

	item := &models.VocabularyItem{
		DeckID:           body.DeckID,
		Phrase:           body.Phrase,
		NormalizedPhrase: body.NormalizedPhrase,
		Meaning:          body.Meaning,
		ExampleSentence:  body.ExampleSentence,
		Note:             body.Note,
		LessonID:         body.LessonID,
		TranscriptID:     body.TranscriptID,
	}
	if createErr := s.repo.Create(ctx, item); createErr != nil {
		log.Errorw("failed to create item", "error", createErr)
		return nil, response.Internal("failed to create item")
	}
	log.Infow("item created", "id", item.ID)
	return item, nil
}

func (s *vocabularyItemService) Update(ctx context.Context, id uint, body req.UpdateVocabularyItemReq) (*models.VocabularyItem, *response.AppError) {
	log := sLog().With("id", id)
	log.Infow("updating vocabulary item")

	actor, appErr := policy.ActorFromContext(ctx)
	if appErr != nil {
		return nil, appErr
	}

	item, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("item not found")
		}
		return nil, response.Internal("failed to get item")
	}

	if item.Deck != nil && item.Deck.UserID != nil {
		if appErr := policy.CanMutate(actor, *item.Deck.UserID); appErr != nil {
			return nil, appErr
		}
	}

	item.Phrase = body.Phrase
	item.NormalizedPhrase = body.NormalizedPhrase
	item.Meaning = body.Meaning
	item.ExampleSentence = body.ExampleSentence
	item.Note = body.Note
	if updateErr := s.repo.Update(ctx, item); updateErr != nil {
		log.Errorw("failed to update item", "error", updateErr)
		return nil, response.Internal("failed to update item")
	}
	log.Infow("item updated")
	return item, nil
}

func (s *vocabularyItemService) Delete(ctx context.Context, id uint) *response.AppError {
	log := sLog().With("id", id)
	log.Infow("deleting vocabulary item")

	actor, appErr := policy.ActorFromContext(ctx)
	if appErr != nil {
		return appErr
	}

	item, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return response.NotFound("item not found")
		}
		return response.Internal("failed to get item")
	}

	if item.Deck != nil && item.Deck.UserID != nil {
		if appErr := policy.CanMutate(actor, *item.Deck.UserID); appErr != nil {
			return appErr
		}
	}

	if delErr := s.repo.Delete(ctx, id); delErr != nil {
		log.Errorw("failed to delete item", "error", delErr)
		return response.Internal("failed to delete item")
	}
	log.Infow("item deleted")
	return nil
}
