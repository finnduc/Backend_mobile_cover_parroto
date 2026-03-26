package models

import (
	"time"

	"gorm.io/gorm"
)

// Base embeds into all entities
type Base struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// User entity
type User struct {
	Base
	Email        string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"type:varchar(255);not null"            json:"-"`
	Name         string `gorm:"type:varchar(255)"                     json:"name"`
	AvatarURL    string `gorm:"type:varchar(500)"                     json:"avatar_url"`
	IsActive     bool   `gorm:"default:true"                          json:"is_active"`
	UserRoles    []UserRole `gorm:"foreignKey:UserID"                json:"-"`
}

// Role entity
type Role struct {
	ID          uint       `gorm:"primaryKey"              json:"id"`
	Name        string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	CreatedAt   time.Time  `json:"created_at"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"  json:"-"`
}

// Permission entity
type Permission struct {
	ID        uint      `gorm:"primaryKey"              json:"id"`
	Name      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// UserRole join table
type UserRole struct {
	UserID uint `gorm:"primaryKey;not null;index" json:"user_id"`
	RoleID uint `gorm:"primaryKey;not null"       json:"role_id"`
}

// RolePermission join table
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey;not null;index" json:"role_id"`
	PermissionID uint `gorm:"primaryKey;not null"       json:"permission_id"`
}

// Lesson entity
type Lesson struct {
	Base
	Title        string       `gorm:"type:varchar(255)"             json:"title"`
	Description  string       `gorm:"type:text"                     json:"description"`
	VideoURL     string       `gorm:"type:varchar(500);not null"    json:"video_url"`
	ThumbnailURL string       `gorm:"type:varchar(500)"             json:"thumbnail_url"`
	Level        string       `gorm:"type:varchar(20)"              json:"level"` // easy|medium|hard
	Duration     float64      `                                     json:"duration"`
	Transcripts  []Transcript `gorm:"foreignKey:LessonID"           json:"transcripts,omitempty"`
	Categories   []LessonCategory `gorm:"foreignKey:LessonID"       json:"-"`
}

// Transcript entity
type Transcript struct {
	Base
	LessonID       uint    `gorm:"not null;index"             json:"lesson_id"`
	Sequence       int     `gorm:"not null"                   json:"sequence"`
	Content        string  `gorm:"type:text;not null"         json:"content"`
	Phonetic       string  `gorm:"type:varchar(500)"          json:"phonetic"`
	Vietnamese     string  `gorm:"type:text"                  json:"vietnamese"`
	StartTimestamp float64 `                                   json:"start_timestamp"`
	EndTimestamp   float64 `                                   json:"end_timestamp"`
}

// Attempt entity
type Attempt struct {
	Base
	UserID     uint         `gorm:"not null;index"            json:"user_id"`
	LessonID   uint         `gorm:"not null;index"            json:"lesson_id"`
	TotalScore float64      `                                 json:"total_score"`
	Completed  bool         `gorm:"default:false"             json:"completed"`
	Answers    []UserAnswer `gorm:"foreignKey:AttemptID"      json:"answers,omitempty"`
}

// UserAnswer entity
type UserAnswer struct {
	Base
	AttemptID    uint    `gorm:"not null;index"            json:"attempt_id"`
	TranscriptID uint    `gorm:"not null;uniqueIndex:idx_attempt_transcript" json:"transcript_id"`
	AttemptIDForIdx uint `gorm:"not null;uniqueIndex:idx_attempt_transcript" json:"-"` // paired composite index
	AnswerText   string  `gorm:"type:text"                 json:"answer_text"`
	Score        float64 `                                 json:"score"`
	IsCorrect    bool    `                                 json:"is_correct"`
}

// UserProgress entity
type UserProgress struct {
	ID           uint      `gorm:"primaryKey"                                             json:"id"`
	UserID       uint      `gorm:"not null;uniqueIndex:idx_user_lesson"                   json:"user_id"`
	LessonID     uint      `gorm:"not null;uniqueIndex:idx_user_lesson"                   json:"lesson_id"`
	LastSequence int       `gorm:"default:0"                                              json:"last_sequence"`
	AnswerCount  int       `gorm:"default:0"                                              json:"answer_count"`
	Completed    bool      `gorm:"default:false"                                          json:"completed"`
	ScoreAvg     float64   `gorm:"default:0"                                              json:"score_avg"`
	UpdatedAt    time.Time `                                                              json:"updated_at"`
}

// Category entity
type Category struct {
	ID   uint   `gorm:"primaryKey"                               json:"id"`
	Name string `gorm:"type:varchar(100);uniqueIndex;not null"   json:"name"`
}

// LessonCategory join table
type LessonCategory struct {
	LessonID   uint `gorm:"primaryKey;not null;index" json:"lesson_id"`
	CategoryID uint `gorm:"primaryKey;not null"       json:"category_id"`
}

// Bookmark entity
type Bookmark struct {
	UserID    uint      `gorm:"primaryKey;not null;index" json:"user_id"`
	LessonID  uint      `gorm:"primaryKey;not null"       json:"lesson_id"`
	CreatedAt time.Time `                                 json:"created_at"`
}
