package req

import "go-cover-parroto/internal/core/database"

type ListBookmarkQuery struct {
	UserID   *string `form:"user_id"`
	LessonID *uint   `form:"lesson_id"`
	Page     int     `form:"page"`
	Limit    int     `form:"limit"`
}

func (q ListBookmarkQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.UserID != nil {
		query.SetFilter("user_id", *q.UserID)
	}
	if q.LessonID != nil {
		query.SetFilter("lesson_id", *q.LessonID)
	}
	return query
}

type AddBookmarkReq struct {
	LessonID uint `uri:"lessonId"`
}

type RemoveBookmarkReq struct {
	LessonID uint `uri:"lessonId"`
}
