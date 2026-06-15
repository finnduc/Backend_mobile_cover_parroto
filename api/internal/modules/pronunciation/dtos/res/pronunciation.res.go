package res

type PronunciationScores struct {
	Accuracy     float64 `json:"accuracy"`
	Fluency      float64 `json:"fluency"`
	Completeness float64 `json:"completeness"`
	Prosody      float64 `json:"prosody"`
}

type PhonemeResult struct {
	Phoneme string  `json:"phoneme"`
	Score   float64 `json:"score"`
}

type WordResult struct {
	Word         string          `json:"word"`
	Score        float64         `json:"score"`
	Feedback     string          `json:"feedback"`
	WeakPhonemes []PhonemeResult `json:"weak_phonemes"`
}

type AttemptInfo struct {
	ID           uint   `json:"id"`
	UserID       string `json:"user_id"`
	LessonID     uint   `json:"lesson_id"`
	TranscriptID uint   `json:"transcript_id"`
	CreatedAt    string `json:"created_at"`
}

type PronunciationRes struct {
	Text         string              `json:"text"`
	OverallScore float64             `json:"overall_score"`
	Scores       PronunciationScores `json:"scores"`
	Feedback     string              `json:"feedback"`
	Words        []WordResult        `json:"words"`
	Attempt      AttemptInfo         `json:"attempt"`
}

type PronunciationProgressRes struct {
	UserID        string   `json:"user_id"`
	LessonID      uint     `json:"lesson_id"`
	TranscriptID  uint     `json:"transcript_id"`
	BestAttemptID *uint    `json:"best_attempt_id"`
	BestScore     *float64 `json:"best_score"`
	Feedback      string   `json:"feedback"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type PronunciationProgressDetailRes struct {
	UserID        string              `json:"user_id"`
	LessonID      uint                `json:"lesson_id"`
	TranscriptID  uint                `json:"transcript_id"`
	BestAttemptID *uint               `json:"best_attempt_id"`
	BestScore     *float64            `json:"best_score"`
	Feedback      string              `json:"feedback"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
	OverallScore  float64             `json:"overall_score"`
	Scores        PronunciationScores `json:"scores"`
	Words         []WordResult        `json:"words"`
}
