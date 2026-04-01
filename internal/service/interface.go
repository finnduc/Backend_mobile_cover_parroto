package service

import (
	"context"
	"go-cover-parroto/internal/models"
	"go-cover-parroto/pkg/utils"
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
	CreateLesson(ctx context.Context, input CreateLessonInput) (*models.Lesson, error)
	GetLesson(ctx context.Context, id uint) (*models.Lesson, error)
	GetTranscripts(ctx context.Context, lessonID uint) ([]models.Transcript, error)
}

type CreateLessonInput struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description"`
	URL          string `json:"url" binding:"required_without=VideoURL,url"`
	VideoURL     string `json:"video_url" binding:"omitempty,url"`
	ThumbnailURL string `json:"thumbnail_url" binding:"omitempty,url"`
	Level        string `json:"level" binding:"omitempty,oneof=easy medium hard"`
	Duration     float64 `json:"duration" binding:"omitempty,gte=0"`
	CategoryIDs  []uint `json:"category_ids"`
	Transcripts  []CreateTranscriptInput `json:"transcripts" binding:"required,min=1,dive"`
}

type CreateTranscriptInput struct {
	Sequence       int     `json:"sequence" binding:"required,gte=1"`
	Content        string  `json:"content" binding:"required"`
	Phonetic       string  `json:"phonetic"`
	Vietnamese     string  `json:"vietnamese"`
	StartTimestamp float64 `json:"start_timestamp" binding:"required,gte=0"`
	EndTimestamp   float64 `json:"end_timestamp" binding:"required,gte=0"`
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
