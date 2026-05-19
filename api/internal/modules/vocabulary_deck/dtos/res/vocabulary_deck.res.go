package res

type VocabularyDeckRes struct {
	ID           uint    `json:"id"`
	UserID       *string `json:"user_id"`
	CategoryID   *uint   `json:"category_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ThumbnailURL string  `json:"thumbnail_url"`
	Level        string  `json:"level"`
	IsDefault    bool    `json:"is_default"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
