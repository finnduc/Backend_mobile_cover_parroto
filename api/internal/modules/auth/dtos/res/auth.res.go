package res

import "go-cover-parroto/internal/core/enums"

type AuthUserRes struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	UserRole  enums.UserRole `json:"user_role"`
	AvatarURL string         `json:"avatar_url"`
	CreatedAt string         `json:"created_at"`
}
