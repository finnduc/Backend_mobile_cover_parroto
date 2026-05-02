package req

import "go-cover-parroto/internal/core/database"

type CreateLessonReq struct {
	CategoryID   *uint   `json:"category_id"`
	Title        string  `json:"title" binding:"required"`
	Description  string  `json:"description"`
	VideoURL     string  `json:"video_url" binding:"required"`
	ThumbnailURL string  `json:"thumbnail_url"`
	Level        string  `json:"level"`
	Duration     float64 `json:"duration"`
}

type ListReq struct {
	CategoryID *uint  `json:"category_id"`
	Level      string `json:"level"`
	Page       int    `json:"page" binding:"min=1"`
	Limit      int    `json:"limit" binding:"min=1,max=100"`
}

type ListLessonQuery struct {
	CategoryID *uint   `form:"category_id"`
	Level      *string `form:"level"`
	Search     *string `form:"search"`
	Page       int     `form:"page"`
	Limit      int     `form:"limit"`
}

func (q ListLessonQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.CategoryID != nil {
		query.SetFilter("category_id", *q.CategoryID)
	}
	if q.Level != nil {
		query.SetFilter("level", *q.Level)
	}
	return query
}

type GetLessonReq struct {
	ID uint `uri:"lessonId"`
}
