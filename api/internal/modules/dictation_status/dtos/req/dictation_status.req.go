package req

import "go-cover-parroto/internal/database"

type ListDictationStatusQuery struct {
	LessonID *uint `form:"lesson_id"`
	Page     int   `form:"page"`
	Limit    int   `form:"limit"`
}

func (q ListDictationStatusQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.LessonID != nil {
		query.SetFilter("lesson_id", *q.LessonID)
	}
	return query
}

type CreateDictationStatusReq struct {
	TranscriptID uint `uri:"transcriptId" binding:"required"`
}
