package req

import "go-cover-parroto/internal/core/database"

type ListHistoryQuery struct {
	UserID   *uint `form:"user_id"`
	LessonID *uint `form:"lesson_id"`
	Page     int   `form:"page"`
	Limit    int   `form:"limit"`
}

func (q ListHistoryQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.UserID != nil {
		query.SetFilter("user_id", *q.UserID)
	}
	if q.LessonID != nil {
		query.SetFilter("lesson_id", *q.LessonID)
	}
	return query
}

type RecordHistoryReq struct {
	LessonID        uint    `json:"lesson_id" binding:"required"`
	DurationWatched float64 `json:"duration_watched" binding:"required,min=0"`
	Completed       bool    `json:"completed"`
}

type GetHistoryReq struct {
	LessonID uint `uri:"lessonId"`
}
