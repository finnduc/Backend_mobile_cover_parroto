package req

import "go-cover-parroto/internal/database"

type ListVocabularyDeckQuery struct {
	CategoryID *uint `form:"category_id"`
	IsDefault  *bool `form:"is_default"`
	Page       int   `form:"page"`
	Limit      int   `form:"limit"`
}

func (q ListVocabularyDeckQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.CategoryID != nil {
		query.SetFilter("category_id", *q.CategoryID)
	}
	if q.IsDefault != nil {
		query.SetFilter("is_default", *q.IsDefault)
	}
	return query
}

type CreateVocabularyDeckReq struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
	Level        string `json:"level"`
}

type UpdateVocabularyDeckReq struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
	Level        string `json:"level"`
}
