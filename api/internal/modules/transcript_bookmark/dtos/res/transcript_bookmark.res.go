package res

type TranscriptBookmarkRes struct {
	LessonID     uint   `json:"lesson_id"`
	TranscriptID uint   `json:"transcript_id"`
	Note         string `json:"note"`
	CreatedAt    string `json:"created_at"`
}

type TranscriptBookmarkLineRes struct {
	TranscriptID uint   `json:"transcript_id"`
	Content      string `json:"content"`
	Phonetic     string `json:"phonetic"`
	Vietnamese   string `json:"vietnamese"`
	Note         string `json:"note"`
	CreatedAt    string `json:"created_at"`
}

type TranscriptBookmarkGroupRes struct {
	LessonID    uint                        `json:"lesson_id"`
	LessonTitle string                      `json:"lesson_title"`
	Transcripts []TranscriptBookmarkLineRes `json:"transcripts"`
}
