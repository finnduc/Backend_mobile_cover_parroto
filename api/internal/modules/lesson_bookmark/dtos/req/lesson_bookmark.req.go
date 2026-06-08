package req

import "go-cover-parroto/internal/database"

type ListLessonBookmarkQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (q ListLessonBookmarkQuery) ToQuery() *database.Query {
	return database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
}

type CreateLessonBookmarkReq struct {
	LessonID uint `json:"lesson_id" binding:"required"`
}
