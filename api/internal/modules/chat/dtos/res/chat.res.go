package res

import "time"

type MessageRes struct {
	ID        uint64    `json:"id"`
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name,omitempty"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type HistoryRes struct {
	Messages []MessageRes `json:"messages"`
	HasMore  bool         `json:"has_more"`
	NextID   *uint64      `json:"next_id,omitempty"`
}

type WSEvent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data,omitempty"`
}
