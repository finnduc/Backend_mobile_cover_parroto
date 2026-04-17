package services

import (
	"context"
	"errors"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/category/dtos/res"
	"go-cover-parroto/internal/modules/category/repositories"
)

type ICategoryService interface {
	ListCategories(ctx context.Context) ([]res.CategoryRes, *response.AppError)
}

type categoryService struct {
	repo repositories.ICategoryRepo
}

func NewCategoryService(repo repositories.ICategoryRepo) ICategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) ListCategories(ctx context.Context) ([]res.CategoryRes, *response.AppError) {
	categories, err := s.repo.FindAll(ctx)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, response.NotFound("no categories found")
		}
		return nil, response.Internal("failed to list categories")
	}
	var result []res.CategoryRes
	for _, cat := range categories {
		result = append(result, res.CategoryRes{ID: cat.ID, Name: cat.Name})
	}
	return result, nil
}
