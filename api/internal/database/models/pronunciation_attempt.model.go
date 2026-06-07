package models

import "time"

type PronunciationAttempt struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	UserID            string    `gorm:"not null;index" json:"user_id"`
	LessonID          uint      `gorm:"not null;index" json:"lesson_id"`
	TranscriptID      uint      `gorm:"not null;index" json:"transcript_id"`
	ReferenceText     string    `gorm:"type:text;not null" json:"reference_text"`
	OverallScore      float64   `gorm:"type:decimal(5,2)" json:"overall_score"`
	AccuracyScore     float64   `gorm:"type:decimal(5,2)" json:"accuracy_score"`
	FluencyScore      float64   `gorm:"type:decimal(5,2)" json:"fluency_score"`
	CompletenessScore float64   `gorm:"type:decimal(5,2)" json:"completeness_score"`
	ProsodyScore      float64   `gorm:"type:decimal(5,2)" json:"prosody_score"`
	Feedback          string    `gorm:"type:text" json:"feedback"`
	WordResults       string    `gorm:"type:jsonb" json:"-"`
	CreatedAt         time.Time `json:"created_at"`
}

func (PronunciationAttempt) TableName() string {
	return "pronunciation_attempts"
}
