package req

import "go-cover-parroto/internal/database"

type ListCategoryQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (q ListCategoryQuery) ToQuery() *database.Query {
	return database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
}

type CreateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}
