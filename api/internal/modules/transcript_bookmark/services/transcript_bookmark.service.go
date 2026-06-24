package services

import (
	"context"
	"errors"
	"time"

	coreError "go-cover-parroto/internal/core/errors"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	db_repos "go-cover-parroto/internal/database/repositories"
	"go-cover-parroto/internal/modules/transcript_bookmark/dtos/req"
	"go-cover-parroto/internal/modules/transcript_bookmark/dtos/res"
)

type ITranscriptBookmarkService interface {
	List(ctx context.Context, userID string, query req.ListTranscriptBookmarkQuery) ([]res.TranscriptBookmarkGroupRes, *response.AppError)
	Create(ctx context.Context, userID string, body req.CreateTranscriptBookmarkReq) (*res.TranscriptBookmarkRes, *response.AppError)
	Update(ctx context.Context, userID string, transcriptID uint, body req.UpdateTranscriptBookmarkReq) (*res.TranscriptBookmarkRes, *response.AppError)
	Delete(ctx context.Context, userID string, transcriptID uint) *response.AppError
}

type transcriptBookmarkService struct {
	repo           db_repos.ITranscriptBookmarkRepo
	transcriptRepo db_repos.ITranscriptRepo
}

func NewTranscriptBookmarkService(repo db_repos.ITranscriptBookmarkRepo, transcriptRepo db_repos.ITranscriptRepo) ITranscriptBookmarkService {
	return &transcriptBookmarkService{repo: repo, transcriptRepo: transcriptRepo}
}

func (s *transcriptBookmarkService) List(ctx context.Context, userID string, query req.ListTranscriptBookmarkQuery) ([]res.TranscriptBookmarkGroupRes, *response.AppError) {
	bookmarks, err := s.repo.FindAllByUser(ctx, userID, query.LessonID)
	if err != nil {
		return nil, response.Internal("failed to list transcript bookmarks")
	}

	groups := make([]res.TranscriptBookmarkGroupRes, 0)
	groupIndex := make(map[uint]int)
	for _, bookmark := range bookmarks {
		if _, ok := groupIndex[bookmark.LessonID]; !ok {
			title := ""
			if bookmark.Lesson != nil {
				title = bookmark.Lesson.Title
			}
			groupIndex[bookmark.LessonID] = len(groups)
			groups = append(groups, res.TranscriptBookmarkGroupRes{
				LessonID:    bookmark.LessonID,
				LessonTitle: title,
				Transcripts: []res.TranscriptBookmarkLineRes{},
			})
		}

		line := res.TranscriptBookmarkLineRes{
			TranscriptID: bookmark.TranscriptID,
			Note:         bookmark.Note,
			CreatedAt:    formatTime(bookmark.CreatedAt),
		}
		if bookmark.Transcript != nil {
			line.Content = bookmark.Transcript.Content
			line.Phonetic = bookmark.Transcript.Phonetic
			line.Vietnamese = bookmark.Transcript.Vietnamese
		}
		idx := groupIndex[bookmark.LessonID]
		groups[idx].Transcripts = append(groups[idx].Transcripts, line)
	}
	return groups, nil
}

func (s *transcriptBookmarkService) Create(ctx context.Context, userID string, body req.CreateTranscriptBookmarkReq) (*res.TranscriptBookmarkRes, *response.AppError) {
	transcript, err := s.transcriptRepo.FindByID(ctx, body.TranscriptID)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("transcript not found")
		}
		return nil, response.Internal("failed to get transcript")
	}

	existing, err := s.repo.FindByUserAndTranscript(ctx, userID, body.TranscriptID)
	if err == nil && existing != nil {
		existing.Note = body.Note
		if updateErr := s.repo.Update(ctx, existing); updateErr != nil {
			return nil, response.Internal("failed to update transcript bookmark")
		}
		return toBookmarkRes(existing), nil
	}
	if err != nil && !errors.Is(err, coreError.ErrNotFound) {
		return nil, response.Internal("failed to get transcript bookmark")
	}

	bookmark := &models.TranscriptBookmark{
		UserID:       userID,
		LessonID:     transcript.LessonID,
		TranscriptID: body.TranscriptID,
		Note:         body.Note,
	}
	if err := s.repo.Create(ctx, bookmark); err != nil {
		return nil, response.Internal("failed to create transcript bookmark")
	}
	return toBookmarkRes(bookmark), nil
}

func (s *transcriptBookmarkService) Update(ctx context.Context, userID string, transcriptID uint, body req.UpdateTranscriptBookmarkReq) (*res.TranscriptBookmarkRes, *response.AppError) {
	bookmark, err := s.repo.FindByUserAndTranscript(ctx, userID, transcriptID)
	if err != nil {
		if errors.Is(err, coreError.ErrNotFound) {
			return nil, response.NotFound("transcript bookmark not found")
		}
		return nil, response.Internal("failed to get transcript bookmark")
	}
	bookmark.Note = body.Note
	if err := s.repo.Update(ctx, bookmark); err != nil {
		return nil, response.Internal("failed to update transcript bookmark")
	}
	return toBookmarkRes(bookmark), nil
}

func (s *transcriptBookmarkService) Delete(ctx context.Context, userID string, transcriptID uint) *response.AppError {
	if err := s.repo.DeleteByUserAndTranscript(ctx, userID, transcriptID); err != nil {
		return response.Internal("failed to delete transcript bookmark")
	}
	return nil
}

func toBookmarkRes(bookmark *models.TranscriptBookmark) *res.TranscriptBookmarkRes {
	return &res.TranscriptBookmarkRes{
		LessonID:     bookmark.LessonID,
		TranscriptID: bookmark.TranscriptID,
		Note:         bookmark.Note,
		CreatedAt:    formatTime(bookmark.CreatedAt),
	}
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
