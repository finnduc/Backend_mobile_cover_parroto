package services

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	"go-cover-parroto/internal/modules/category/dtos/req"
	"go-cover-parroto/internal/modules/category/dtos/res"
	"go-cover-parroto/internal/modules/category/repositories"
	"go-cover-parroto/internal/utils"
)

type ICategoryService interface {
	List(ctx context.Context, query req.ListCategoryQuery) (*response.PaginatedResponse[res.CategoryRes], *response.AppError)
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

func (s *categoryService) Create(ctx context.Context, body req.CreateCategoryReq) (*res.CategoryRes, *response.AppError) {
	category := &models.Category{Name: body.Name}
	if err := s.repo.Create(ctx, category); err != nil {
		return nil, response.Internal("failed to create category")
	}
	var result res.CategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		return nil, response.Internal("failed to map category")
	}
	return &result, nil
}

func (s *categoryService) Update(ctx context.Context, id uint, body req.UpdateCategoryReq) (*res.CategoryRes, *response.AppError) {
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, response.NotFound("category not found")
	}
	category.Name = body.Name
	if updateErr := s.repo.Update(ctx, category); updateErr != nil {
		return nil, response.Internal("failed to update category")
	}
	var result res.CategoryRes
	if err := utils.MapToDTO(category, &result); err != nil {
		return nil, response.Internal("failed to map category")
	}
	return &result, nil
}

func (s *categoryService) Delete(ctx context.Context, id uint) *response.AppError {
	if err := s.repo.Delete(ctx, id); err != nil {
		return response.Internal("failed to delete category")
	}
	return nil
}

func (s *categoryService) List(ctx context.Context, query req.ListCategoryQuery) (*response.PaginatedResponse[res.CategoryRes], *response.AppError) {
	result, err := s.repo.FindAll(ctx, query.ToQuery())
	if err != nil {
		return nil, response.Internal("failed to list categories")
	}
	var categories []res.CategoryRes
	if err := utils.MapToDTOs(result.Data, &categories); err != nil {
		return nil, response.Internal("failed to map categories")
	}
	return &response.PaginatedResponse[res.CategoryRes]{Data: categories, Meta: result.Meta}, nil
}
