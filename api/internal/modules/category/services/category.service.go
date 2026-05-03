package services

import (
	"context"
	"errors"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/category/dtos/req"
	"go-cover-parroto/internal/modules/category/dtos/res"
	"go-cover-parroto/internal/modules/category/repositories"
	"go-cover-parroto/internal/utils"
	"go.uber.org/zap"
)

func sLog() *zap.SugaredLogger {
	return logger.S().With("service", "category")
}

type ICategoryService interface {
	List(ctx context.Context, query req.ListCategoryQuery) (*response.PaginatedResponse[res.CategoryRes], *response.AppError)
	GetByID(ctx context.Context, id uint) (*res.CategoryRes, *response.AppError)
	Create(ctx context.Context, body req.CreateCategoryReq) (*res.CategoryRes, *response.AppError)
	Update(ctx context.Context, id uint, body req.UpdateCategoryReq) (*res.CategoryRes, *response.AppError)
	Delete(ctx context.Context, id uint) *response.AppError
}

type categoryService struct {
	repo repositories.ICategoryRepo
}

func NewCategoryService(repo repositories.ICategoryRepo) ICategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetByID(ctx context.Context, id uint) (*res.CategoryRes, *response.AppError) {
	log := sLog().With("categoryId", id)
	log.Infow("getting category by id")
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("category not found")
		}
		log.Errorw("failed to get category", "error", err)
		return nil, response.Internal("failed to get category")
	}
	var result res.CategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		return nil, response.Internal("failed to map category")
	}
	return &result, nil
}

func (s *categoryService) Create(ctx context.Context, body req.CreateCategoryReq) (*res.CategoryRes, *response.AppError) {
	log := sLog()
	log.Infow("creating category")
	category := &models.Category{Name: body.Name}
	if err := s.repo.Create(ctx, category); err != nil {
		log.Errorw("failed to create category", "error", err)
		return nil, response.Internal("failed to create category")
	}
	log.Infow("category created", "categoryId", category.ID)
	var result res.CategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		return nil, response.Internal("failed to map category")
	}
	return &result, nil
}

func (s *categoryService) Update(ctx context.Context, id uint, body req.UpdateCategoryReq) (*res.CategoryRes, *response.AppError) {
	log := sLog().With("categoryId", id)
	log.Infow("updating category")
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, response.NotFound("category not found")
	}
	category.Name = body.Name
	if updateErr := s.repo.Update(ctx, category); updateErr != nil {
		log.Errorw("failed to update category", "error", updateErr)
		return nil, response.Internal("failed to update category")
	}
	log.Infow("category updated")
	var result res.CategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		return nil, response.Internal("failed to map category")
	}
	return &result, nil
}

func (s *categoryService) Delete(ctx context.Context, id uint) *response.AppError {
	log := sLog().With("categoryId", id)
	log.Infow("deleting category")
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Errorw("failed to delete category", "error", err)
		return response.Internal("failed to delete category")
	}
	log.Infow("category deleted")
	return nil
}

func (s *categoryService) List(ctx context.Context, query req.ListCategoryQuery) (*response.PaginatedResponse[res.CategoryRes], *response.AppError) {
	log := sLog()
	log.Infow("listing categories")
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		log.Errorw("failed to list categories", "error", err)
		return nil, response.Internal("failed to list categories")
	}
	var categories []res.CategoryRes
	if err := utils.MapToDTOs(result.Data, &categories); err != nil {
		return nil, response.Internal("failed to map categories")
	}
	return &response.PaginatedResponse[res.CategoryRes]{Data: categories, Meta: result.Meta}, nil
}
