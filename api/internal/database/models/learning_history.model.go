package models

import "time"

type LearningHistory struct {
	ID                     uint      `gorm:"primaryKey" json:"id"`
	UserID                 string    `gorm:"not null;index" json:"user_id"`
	LessonID               uint      `gorm:"not null;index" json:"lesson_id"`
	CompletedDictation     bool      `gorm:"default:false" json:"completed_dictation"`
	CompletedPronunciation *bool     `json:"completed_pronunciation"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

func (LearningHistory) TableName() string {
	return "learning_history"
}
