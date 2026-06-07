package req

import "go-cover-parroto/internal/database"

type ListTranscriptBookmarkQuery struct {
	UserID   *string `form:"user_id"`
	LessonID *uint   `form:"lesson_id"`
	Page     int     `form:"page"`
	Limit    int     `form:"limit"`
}

func (q ListTranscriptBookmarkQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.UserID != nil {
		query.SetFilter("user_id", *q.UserID)
	}
	if q.LessonID != nil {
		query.SetFilter("lesson_id", *q.LessonID)
	}
	return query
}

type CreateTranscriptBookmarkReq struct {
	TranscriptID uint   `json:"transcript_id" binding:"required"`
	Note         string `json:"note"`
}

type UpdateTranscriptBookmarkNoteReq struct {
	Note string `json:"note" binding:"required"`
}
