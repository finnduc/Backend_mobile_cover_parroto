package service

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/pkg/utils"
)

// ===== Auth Service =====

type RegisterInput struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name"     binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // seconds
}

type IAuthService interface {
	Register(ctx context.Context, input RegisterInput) (*models.User, error)
	Login(ctx context.Context, input LoginInput) (*TokenPair, error)
	Logout(ctx context.Context, accessToken string, userID uint) error
	Refresh(ctx context.Context, refreshToken string, userID uint) (*TokenPair, error)
}

// ===== Lesson Service =====

type ILessonService interface {
	ListLessons(ctx context.Context, level string, categoryID *uint, q utils.PaginationQuery) ([]models.Lesson, int64, error)
	GetLesson(ctx context.Context, id uint) (*models.Lesson, error)
	GetTranscripts(ctx context.Context, lessonID uint) ([]models.Transcript, error)
}

// ===== Attempt Service =====

type IAttemptService interface {
	CreateAttempt(ctx context.Context, userID, lessonID uint) (*models.Attempt, error)
	GetAttempt(ctx context.Context, id uint) (*models.Attempt, error)
}

// ===== Answer Service =====

type SubmitAnswerInput struct {
	AttemptID    uint   `json:"attempt_id"    binding:"required,gt=0"`
	TranscriptID uint   `json:"transcript_id" binding:"required,gt=0"`
	AnswerText   string `json:"answer_text"   binding:"required"`
}

type AnswerItemInput struct {
	TranscriptID uint   `json:"transcript_id" binding:"required,gt=0"`
	AnswerText   string `json:"answer_text"   binding:"required"`
}

type BulkSubmitInput struct {
	AttemptID uint              `json:"attempt_id" binding:"required,gt=0"`
	Answers   []AnswerItemInput `json:"answers"    binding:"required,min=1,dive"`
}

type AnswerResult struct {
	Score       float64 `json:"score"`
	IsCorrect   bool    `json:"is_correct"`
	CorrectText string  `json:"correct_text"`
}

type BulkAnswerResult struct {
	Results []AnswerResult `json:"results"`
}

type IAnswerService interface {
	SubmitAnswer(ctx context.Context, userID uint, input SubmitAnswerInput) (*AnswerResult, error)
	BulkSubmit(ctx context.Context, userID uint, input BulkSubmitInput) (*BulkAnswerResult, error)
}

// ===== Progress Service =====

type IProgressService interface {
	GetProgress(ctx context.Context, userID, lessonID uint) (*models.UserProgress, error)
}

// ===== Bookmark Service =====

type IBookmarkService interface {
	Toggle(ctx context.Context, userID, lessonID uint) (bool, error)
	ListBookmarks(ctx context.Context, userID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error)
}

// ===== Category Service =====

type ICategoryService interface {
	ListCategories(ctx context.Context) ([]models.Category, error)
	GetLessonsByCategory(ctx context.Context, categoryID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error)
}
