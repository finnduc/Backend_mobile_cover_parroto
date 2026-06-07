package models

import "time"

type TranscriptProgress struct {
	UserID       string    `gorm:"primaryKey" json:"user_id"`
	TranscriptID uint      `gorm:"primaryKey" json:"transcript_id"`
	LessonID     uint      `gorm:"not null;index" json:"lesson_id"`
	CompletedAt  time.Time `json:"completed_at"`
}

func (TranscriptProgress) TableName() string {
	return "transcript_progress"
}
