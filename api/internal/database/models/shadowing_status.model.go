package models

import "time"

type ShadowingStatus struct {
	UserID       string    `gorm:"primaryKey" json:"user_id"`
	TranscriptID uint      `gorm:"primaryKey" json:"transcript_id"`
	LessonID     uint      `gorm:"not null;index" json:"lesson_id"`
	CompletedAt  time.Time `json:"completed_at"`
}

func (ShadowingStatus) TableName() string {
	return "shadowing_status"
}
