package models

import (
	"go-cover-parroto/internal/core/enums"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	UserRole  enums.UserRole `gorm:"type:varchar(255);not null" json:"-"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	AvatarURL string         `gorm:"type:varchar(500)" json:"avatar_url"`
	CreatedAt time.Time      `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
