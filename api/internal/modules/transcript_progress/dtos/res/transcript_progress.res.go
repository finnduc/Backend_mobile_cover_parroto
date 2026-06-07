package res

type TranscriptProgressRes struct {
	UserID       string `json:"user_id"`
	TranscriptID uint   `json:"transcript_id"`
	LessonID     uint   `json:"lesson_id"`
	CompletedAt  string `json:"completed_at"`
}
