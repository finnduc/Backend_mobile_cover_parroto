package req

type ListTranscriptBookmarkQuery struct {
	LessonID *uint `form:"lesson_id"`
}

type CreateTranscriptBookmarkReq struct {
	TranscriptID uint   `json:"transcript_id" binding:"required"`
	Note         string `json:"note"`
}

type UpdateTranscriptBookmarkReq struct {
	Note string `json:"note"`
}
