package req

import "go-cover-parroto/internal/database"

type ListShadowingStatusQuery struct {
	LessonID *uint `form:"lesson_id"`
	Page     int   `form:"page"`
	Limit    int   `form:"limit"`
}

func (q ListShadowingStatusQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.LessonID != nil {
		query.SetFilter("lesson_id", *q.LessonID)
	}
	return query
}

type CreateShadowingStatusReq struct {
	TranscriptID uint    `json:"transcript_id" binding:"required"`
	LessonID     uint    `json:"lesson_id" binding:"required"`
	BestScore    float64 `json:"best_score"`
	Feedback     string  `json:"feedback"`
}
