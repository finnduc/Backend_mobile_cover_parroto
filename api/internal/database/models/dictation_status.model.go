package models

import "time"

type DictationStatus struct {
	UserID       string    `gorm:"primaryKey" json:"user_id"`
	TranscriptID uint      `gorm:"primaryKey" json:"transcript_id"`
	LessonID     uint      `gorm:"not null;index" json:"lesson_id"`
	CompletedAt  time.Time `json:"completed_at"`
}

func (DictationStatus) TableName() string {
	return "dictation_status"
}
