package service

import (
	"context"
	"fmt"
	"go-familytree/internal/models"
	"go-familytree/internal/repo"
	pkgerrors "go-familytree/pkg/errors"
	"go-familytree/pkg/utils"
)

type bookmarkService struct{ bookmarkRepo repo.IBookmarkRepo }

func NewBookmarkService(bookmarkRepo repo.IBookmarkRepo) IBookmarkService {
	return &bookmarkService{bookmarkRepo: bookmarkRepo}
}

func (s *bookmarkService) Toggle(ctx context.Context, userID, lessonID uint) (bool, error) {
	bookmarked, err := s.bookmarkRepo.Toggle(ctx, userID, lessonID)
	if err != nil {
		return false, fmt.Errorf("bookmarkService.Toggle: %w", pkgerrors.ErrInternalServer)
	}
	return bookmarked, nil
}

func (s *bookmarkService) ListBookmarks(ctx context.Context, userID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	lessons, total, err := s.bookmarkRepo.FindByUser(ctx, userID, q)
	if err != nil {
		return nil, 0, fmt.Errorf("bookmarkService.ListBookmarks: %w", pkgerrors.ErrInternalServer)
	}
	return lessons, total, nil
}

// ─── Category ───────────────────────────────────────────────────────────────

type categoryService struct{ categoryRepo repo.ICategoryRepo }

func NewCategoryService(categoryRepo repo.ICategoryRepo) ICategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) ListCategories(ctx context.Context) ([]models.Category, error) {
	cats, err := s.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("categoryService.ListCategories: %w", pkgerrors.ErrInternalServer)
	}
	return cats, nil
}

func (s *categoryService) GetLessonsByCategory(ctx context.Context, categoryID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error) {
	lessons, total, err := s.categoryRepo.FindLessonsByCategory(ctx, categoryID, q)
	if err != nil {
		return nil, 0, fmt.Errorf("categoryService.GetLessonsByCategory: %w", pkgerrors.ErrInternalServer)
	}
	return lessons, total, nil
}
