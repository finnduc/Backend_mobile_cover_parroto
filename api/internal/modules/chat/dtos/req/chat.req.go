package req

type ListMessagesQuery struct {
	BeforeID uint64 `form:"before_id" example:"100"`
	Limit    int    `form:"limit,default=20" example:"20"`
}

type SendMessagePayload struct {
	Content string `json:"content" binding:"required,min=1,max=1000" example:"Hello world"`
}
