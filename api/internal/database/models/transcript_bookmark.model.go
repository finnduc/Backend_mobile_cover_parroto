package models

import "time"

type TranscriptBookmark struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	UserID       string      `gorm:"not null;index" json:"user_id"`
	TranscriptID uint        `gorm:"not null;index" json:"transcript_id"`
	LessonID     uint        `gorm:"not null;index" json:"lesson_id"`
	Note         string      `gorm:"type:text" json:"note"`
	CreatedAt    time.Time   `json:"created_at"`
	Transcript   *Transcript `gorm:"foreignKey:TranscriptID" json:"-"`
	Lesson       *Lesson     `gorm:"foreignKey:LessonID" json:"-"`
}

func (TranscriptBookmark) TableName() string {
	return "transcript_bookmarks"
}
