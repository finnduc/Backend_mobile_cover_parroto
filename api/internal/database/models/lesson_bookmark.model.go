package models

import "time"

type LessonBookmark struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"not null;index" json:"user_id"`
	LessonID  uint      `gorm:"not null;index" json:"lesson_id"`
	CreatedAt time.Time `json:"created_at"`
	Lesson    *Lesson   `gorm:"foreignKey:LessonID" json:"-"`
}

func (LessonBookmark) TableName() string {
	return "lesson_bookmarks"
}
