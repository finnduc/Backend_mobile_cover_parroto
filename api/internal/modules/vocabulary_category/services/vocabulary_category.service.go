package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/vocabulary_category/dtos/req"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "vocabulary_category")
}

type IVocabularyCategoryService interface {
	List(ctx context.Context, query req.ListVocabularyCategoryQuery) (*response.PaginatedResult[*models.VocabularyCategory], *response.AppError)
	GetByID(ctx context.Context, id uint) (*models.VocabularyCategory, *response.AppError)
	Create(ctx context.Context, body req.CreateVocabularyCategoryReq) (*models.VocabularyCategory, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateVocabularyCategoryReq) (*models.VocabularyCategory, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
}

type vocabularyCategoryService struct {
	repo db_repos.IVocabularyCategoryRepo
}

func NewVocabularyCategoryService(repo db_repos.IVocabularyCategoryRepo) IVocabularyCategoryService {
	return &vocabularyCategoryService{repo: repo}
}

func (s *vocabularyCategoryService) List(ctx context.Context, query req.ListVocabularyCategoryQuery) (*response.PaginatedResult[*models.VocabularyCategory], *response.AppError) {
	log := sLog()
	log.Infow("listing vocabulary categories")
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list vocabulary categories", "error", err)
		return nil, response.Internal("failed to list vocabulary categories")
	}
	return result, nil
}

func (s *vocabularyCategoryService) GetByID(ctx context.Context, id uint) (*models.VocabularyCategory, *response.AppError) {
	log := sLog().With("id", id)
	log.Infow("getting vocabulary category")
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("category not found")
		}
		log.Errorw("failed to get category", "error", err)
		return nil, response.Internal("failed to get category")
	}
	return category, nil
}

func (s *vocabularyCategoryService) Create(ctx context.Context, body req.CreateVocabularyCategoryReq) (*models.VocabularyCategory, *response.AppError) {
	log := sLog()
	log.Infow("creating vocabulary category")
	category := &models.VocabularyCategory{Name: body.Name, Description: body.Description}
	if err := s.repo.Create(ctx, category); err != nil {
		log.Errorw("failed to create category", "error", err)
		return nil, response.Internal("failed to create category")
	}
	log.Infow("category created", "id", category.ID)
	return category, nil
}

func (s *vocabularyCategoryService) Update(ctx context.Context, id uint, body req.UpdateVocabularyCategoryReq) (*models.VocabularyCategory, *response.AppError) {
	log := sLog().With("id", id)
	log.Infow("updating vocabulary category")
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("category not found")
		}
		log.Errorw("failed to get category for update", "error", err)
		return nil, response.Internal("failed to get category")
	}
	if body.Name != nil {
		category.Name = *body.Name
	}
	if body.Description != nil {
		category.Description = *body.Description
	}
	if updateErr := s.repo.Update(ctx, category); updateErr != nil {
		log.Errorw("failed to update category", "error", updateErr)
		return nil, response.Internal("failed to update category")
	}
	log.Infow("category updated")
	return category, nil
}

func (s *vocabularyCategoryService) Delete(ctx context.Context, id uint) *response.AppError {
	log := sLog().With("id", id)
	log.Infow("deleting vocabulary category")
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return response.NotFound("category not found")
		}
		log.Errorw("failed to get category for delete", "error", err)
		return response.Internal("failed to get category")
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Errorw("failed to delete category", "error", err)
		return response.Internal("failed to delete category")
	}
	log.Infow("category deleted")
	return nil
}
