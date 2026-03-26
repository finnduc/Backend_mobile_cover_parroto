package repo

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/pkg/utils"
)

// ===== Auth Repo =====

type IAuthRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserRoles(ctx context.Context, userID uint) ([]models.Role, error)
}

// ===== Lesson Repo =====

type LessonFilter struct {
	Level      string
	CategoryID *uint
	utils.PaginationQuery
}

type ILessonRepo interface {
	FindAll(ctx context.Context, filter LessonFilter) ([]models.Lesson, int64, error)
	FindByID(ctx context.Context, id uint) (*models.Lesson, error)
	FindTranscriptsByLessonID(ctx context.Context, lessonID uint) ([]models.Transcript, error)
}

// ===== Attempt Repo =====

type IAttemptRepo interface {
	Create(ctx context.Context, attempt *models.Attempt) error
	FindByID(ctx context.Context, id uint) (*models.Attempt, error)
	CountAnswers(ctx context.Context, attemptID uint) (int64, error)
	MarkCompleted(ctx context.Context, attemptID uint, totalScore float64) error
}

// ===== Answer Repo =====

type IAnswerRepo interface {
	Upsert(ctx context.Context, answer *models.UserAnswer) error
	BulkUpsert(ctx context.Context, answers []models.UserAnswer) error
	FindByAttemptAndTranscript(ctx context.Context, attemptID, transcriptID uint) (*models.UserAnswer, error)
}

// ===== Progress Repo =====

type IProgressRepo interface {
	FindOrCreate(ctx context.Context, userID, lessonID uint) (*models.UserProgress, error)
	Save(ctx context.Context, p *models.UserProgress) error
}

// ===== Bookmark Repo =====

type IBookmarkRepo interface {
	Toggle(ctx context.Context, userID, lessonID uint) (bool, error)
	FindByUser(ctx context.Context, userID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error)
}

// ===== Category Repo =====

type ICategoryRepo interface {
	FindAll(ctx context.Context) ([]models.Category, error)
	FindLessonsByCategory(ctx context.Context, categoryID uint, q utils.PaginationQuery) ([]models.Lesson, int64, error)
}
