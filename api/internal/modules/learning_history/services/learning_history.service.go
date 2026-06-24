package services

import (
	"context"
	"errors"
	"time"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go-cover-parroto/internal/modules/learning_history/dtos/req"
	"go-cover-parroto/internal/modules/learning_history/dtos/res"
)

type ILearningHistoryService interface {
	List(ctx context.Context, userID string, filter string) ([]res.LearningHistoryRes, *response.AppError)
	Create(ctx context.Context, userID string, body req.CreateLearningHistoryReq) (*res.LearningHistoryRes, *response.AppError)
	GetByLesson(ctx context.Context, userID string, lessonID uint) (*res.LearningHistoryRes, *response.AppError)
	Summary(ctx context.Context, userID string) (*res.LearningHistorySummaryRes, *response.AppError)
	LessonSummary(ctx context.Context, userID string, lessonID uint) (*res.LessonProgressSummaryRes, *response.AppError)
}

type learningHistoryService struct {
	repo db_repos.ILearningHistoryRepo
}

func NewLearningHistoryService(repo db_repos.ILearningHistoryRepo) ILearningHistoryService {
	return &learningHistoryService{repo: repo}
}

func (s *learningHistoryService) List(ctx context.Context, userID string, filter string) ([]res.LearningHistoryRes, *response.AppError) {
	histories, err := s.repo.FindAllByUser(ctx, userID, filter)
	if err != nil {
		return nil, response.Internal("failed to list learning history")
	}
	return toHistoryResList(histories), nil
}

func (s *learningHistoryService) Create(ctx context.Context, userID string, body req.CreateLearningHistoryReq) (*res.LearningHistoryRes, *response.AppError) {
	history, err := s.repo.FindByUserAndLesson(ctx, userID, body.LessonID)
	if err != nil && !errors.Is(err, coreError.ErrNotFound) {
		return nil, response.Internal("failed to get learning history")
	}

	if history == nil {
		history = &models.LearningHistory{UserID: userID, LessonID: body.LessonID}
		if body.CompletedDictation != nil {
			history.CompletedDictation = *body.CompletedDictation
		}
		if body.CompletedPronunciation != nil {
			history.CompletedPronunciation = *body.CompletedPronunciation
		}
		if err := s.repo.Create(ctx, history); err != nil {
			return nil, response.Internal("failed to create learning history")
		}
		return toHistoryRes(history), nil
	}

	if body.CompletedDictation != nil {
		history.CompletedDictation = *body.CompletedDictation
	}
	if body.CompletedPronunciation != nil {
		history.CompletedPronunciation = *body.CompletedPronunciation
	}
	if err := s.repo.Update(ctx, history); err != nil {
		return nil, response.Internal("failed to update learning history")
	}
	return toHistoryRes(history), nil
}

func (s *learningHistoryService) GetByLesson(ctx context.Context, userID string, lessonID uint) (*res.LearningHistoryRes, *response.AppError) {
	history, err := s.repo.FindByUserAndLesson(ctx, userID, lessonID)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("learning history not found")
		}
		return nil, response.Internal("failed to get learning history")
	}
	return toHistoryRes(history), nil
}

func (s *learningHistoryService) Summary(ctx context.Context, userID string) (*res.LearningHistorySummaryRes, *response.AppError) {
	completed, unfinished, err := s.repo.CountSummary(ctx, userID)
	if err != nil {
		return nil, response.Internal("failed to summarize learning history")
	}
	return &res.LearningHistorySummaryRes{CompletedCount: completed, UnfinishedCount: unfinished}, nil
}

func (s *learningHistoryService) LessonSummary(ctx context.Context, userID string, lessonID uint) (*res.LessonProgressSummaryRes, *response.AppError) {
	completed, total, err := s.repo.CountCompletedLessonTranscripts(ctx, userID, lessonID)
	if err != nil {
		return nil, response.Internal("failed to summarize lesson progress")
	}
	uncompleted := total - completed
	if uncompleted < 0 {
		uncompleted = 0
	}
	return &res.LessonProgressSummaryRes{Completed: completed, Uncompleted: uncompleted}, nil
}

func toHistoryResList(histories []*models.LearningHistory) []res.LearningHistoryRes {
	result := make([]res.LearningHistoryRes, 0, len(histories))
	for _, history := range histories {
		result = append(result, *toHistoryRes(history))
	}
	return result
}

func toHistoryRes(history *models.LearningHistory) *res.LearningHistoryRes {
	return &res.LearningHistoryRes{
		ID:                     history.ID,
		UserID:                 history.UserID,
		LessonID:               history.LessonID,
		CompletedDictation:     history.CompletedDictation,
		CompletedPronunciation: history.CompletedPronunciation,
		CreatedAt:              formatTime(history.CreatedAt),
		UpdatedAt:              formatTime(history.UpdatedAt),
	}
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
