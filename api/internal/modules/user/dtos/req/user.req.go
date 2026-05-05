package req

import "go-cover-parroto/internal/database"

type ListUserQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (q ListUserQuery) ToQuery() *database.Query {
	return database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
}

type UpdateUserReq struct {
	Name      *string `json:"name"`
	AvatarURL *string `json:"avatar_url"`
}
