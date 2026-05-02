package services

import (
	"context"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/category/dtos/req"
	"go-cover-parroto/internal/modules/category/dtos/res"
	"go-cover-parroto/internal/modules/category/repositories"
	"go-cover-parroto/internal/utils"
)

type ICategoryService interface {
	List(ctx context.Context, query req.ListCategoryQuery) (*response.PaginatedResponse[res.CategoryRes], *response.AppError)
}

type categoryService struct {
	repo repositories.ICategoryRepo
}

func NewCategoryService(repo repositories.ICategoryRepo) ICategoryService {
	return &categoryService{repo: repo}
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
