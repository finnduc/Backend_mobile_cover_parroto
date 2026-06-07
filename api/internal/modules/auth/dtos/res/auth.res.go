package res

type UserRes struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	UserRole  string `json:"user_role"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
}
