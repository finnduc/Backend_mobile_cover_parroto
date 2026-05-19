package res

type VocabularyItemRes struct {
	ID               uint   `json:"id"`
	DeckID           uint   `json:"deck_id"`
	LessonID         *uint  `json:"lesson_id"`
	TranscriptID     *uint  `json:"transcript_id"`
	Phrase           string `json:"phrase"`
	NormalizedPhrase string `json:"normalized_phrase"`
	Meaning          string `json:"meaning"`
	ExampleSentence  string `json:"example_sentence"`
	Note             string `json:"note"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
