package models

import "time"

type PronunciationProgress struct {
	UserID        string    `gorm:"primaryKey" json:"user_id"`
	TranscriptID  uint      `gorm:"primaryKey" json:"transcript_id"`
	LessonID      uint      `gorm:"not null;index" json:"lesson_id"`
	BestAttemptID *uint     `json:"best_attempt_id"`
	BestScore     *float64  `gorm:"type:decimal(5,2)" json:"best_score"`
	Feedback      string    `gorm:"type:text" json:"feedback"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (PronunciationProgress) TableName() string {
	return "pronunciation_progress"
}
