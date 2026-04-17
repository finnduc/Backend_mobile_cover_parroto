package req

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
