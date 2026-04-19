package models

import "time"

type Bookmark struct {
	UserID    uint      `gorm:"primaryKey" json:"user_id"`
	LessonID  uint      `gorm:"primaryKey" json:"lesson_id"`
	CreatedAt time.Time `json:"created_at"`
	Lesson    *Lesson   `gorm:"foreignKey:LessonID" json:"-"`
}

func (Bookmark) TableName() string {
	return "bookmarks"
}
