package res

type ShadowingStatusRes struct {
	UserID       string  `json:"user_id"`
	TranscriptID uint    `json:"transcript_id"`
	LessonID     uint    `json:"lesson_id"`
	BestScore    float64 `json:"best_score"`
	Feedback     string  `json:"feedback"`
	CompletedAt  string  `json:"completed_at"`
}
