package models

import "time"

type LearningHistory struct {
	ID                     uint      `gorm:"primaryKey" json:"id"`
	UserID                 string    `gorm:"not null;index;uniqueIndex:idx_learning_history_user_lesson" json:"user_id"`
	LessonID               uint      `gorm:"not null;index;uniqueIndex:idx_learning_history_user_lesson" json:"lesson_id"`
	CompletedDictation     bool      `gorm:"not null;default:false" json:"completed_dictation"`
	CompletedPronunciation bool      `gorm:"not null;default:false" json:"completed_pronunciation"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	Lesson                 *Lesson   `gorm:"foreignKey:LessonID" json:"-"`
}

func (LearningHistory) TableName() string {
	return "learning_history"
}
