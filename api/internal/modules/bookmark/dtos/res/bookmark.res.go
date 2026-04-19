package res

import "time"

type LessonInfo struct {
	ID           uint    `json:"id"`
	Title        string  `json:"title"`
	ThumbnailURL string  `json:"thumbnail_url"`
	Level        string  `json:"level"`
	Duration     float64 `json:"duration"`
}

type BookmarkRes struct {
	UserID    uint        `json:"user_id"`
	LessonID  uint        `json:"lesson_id"`
	CreatedAt time.Time   `json:"created_at"`
	Lesson    *LessonInfo `json:"lesson,omitempty"`
}
