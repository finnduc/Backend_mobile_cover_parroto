package res

type ProgressRes struct {
	UserID          uint    `json:"user_id"`
	LessonID        uint    `json:"lesson_id"`
	DurationWatched float64 `json:"duration_watched"`
	Completed       bool    `json:"completed"`
	CreatedAt       string  `json:"created_at"`
}
