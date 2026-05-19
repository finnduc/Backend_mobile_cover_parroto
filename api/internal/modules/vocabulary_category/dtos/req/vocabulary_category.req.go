package req

import "go-cover-parroto/internal/database"

type ListVocabularyCategoryQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (q ListVocabularyCategoryQuery) ToQuery() *database.Query {
	return database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
}

type CreateVocabularyCategoryReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateVocabularyCategoryReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
