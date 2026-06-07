package models

import "time"

type GlobalChatMessage struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string    `gorm:"type:varchar(255);not null;index" json:"user_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"index:idx_global_chat_messages_created_at,sort:desc" json:"created_at"`
	User      *User     `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (GlobalChatMessage) TableName() string {
	return "global_chat_messages"
}
