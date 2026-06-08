package req

type ListMessagesQuery struct {
	BeforeID uint64 `form:"before_id" example:"100"`
	Limit    int    `form:"limit,default=20" example:"20"`
}

type SendMessageReq struct {
	Content string `json:"content" binding:"required" example:"Hello world"`
}

