package models

import "time"

type LearningHistory struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null;index" json:"user_id"`
	LessonID        uint      `gorm:"not null;index" json:"lesson_id"`
	DurationWatched float64   `json:"duration_watched"`
	Completed       bool      `gorm:"default:true" json:"completed"`
	CreatedAt       time.Time `json:"created_at"`
}

func (LearningHistory) TableName() string {
	return "learning_history"
}
