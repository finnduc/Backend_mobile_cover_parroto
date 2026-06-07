package res

type TranscriptBookmarkRes struct {
	ID           uint   `json:"id"`
	UserID       string `json:"user_id"`
	TranscriptID uint   `json:"transcript_id"`
	Note         string `json:"note"`
	CreatedAt    string `json:"created_at"`
}
