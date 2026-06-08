package res

type LessonBookmarkRes struct {
	ID        uint   `json:"id"`
	UserID    string `json:"user_id"`
	LessonID  uint   `json:"lesson_id"`
	CreatedAt string `json:"created_at"`
}
