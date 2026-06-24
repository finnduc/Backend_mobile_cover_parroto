package models

import "time"

type TranscriptBookmark struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	UserID       string      `gorm:"not null;index;uniqueIndex:idx_transcript_bookmark_user_transcript" json:"user_id"`
	LessonID     uint        `gorm:"not null;index" json:"lesson_id"`
	TranscriptID uint        `gorm:"not null;index;uniqueIndex:idx_transcript_bookmark_user_transcript" json:"transcript_id"`
	Note         string      `gorm:"type:text" json:"note"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Lesson       *Lesson     `gorm:"foreignKey:LessonID" json:"-"`
	Transcript   *Transcript `gorm:"foreignKey:TranscriptID" json:"-"`
}

func (TranscriptBookmark) TableName() string {
	return "transcript_bookmarks"
}
