package req

type CreateTranscriptProgressReq struct {
	TranscriptID uint `json:"transcript_id" binding:"required"`
}
