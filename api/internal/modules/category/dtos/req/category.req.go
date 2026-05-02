package req

import "go-cover-parroto/internal/core/database"

type ListCategoryQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (q ListCategoryQuery) ToQuery() *database.Query {
	return database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
}
