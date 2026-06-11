package req

import "go-cover-parroto/internal/database"

type ListVocabularyItemQuery struct {
	DeckID *uint `form:"deck_id"`
	Page   int   `form:"page"`
	Limit  int   `form:"limit"`
}

func (q ListVocabularyItemQuery) ToQuery() *database.Query {
	query := database.NewQuery().SetPage(q.Page).SetLimit(q.Limit)
	if q.DeckID != nil {
		query.SetFilter("deck_id", *q.DeckID)
	}
	return query
}

type CreateVocabularyItemReq struct {
	Phrase           string `json:"phrase"`
	NormalizedPhrase string `json:"normalized_phrase"`
	Meaning          string `json:"meaning"`
	ExampleSentence  string `json:"example_sentence"`
	Note             string `json:"note"`
	LessonID         *uint  `json:"lesson_id"`
	TranscriptID     *uint  `json:"transcript_id"`
}

type CreateVocabularyItemFromDeckReq struct {
	DeckID           uint   `uri:"id"`
	Phrase           string `json:"phrase"`
	NormalizedPhrase string `json:"normalized_phrase"`
	Meaning          string `json:"meaning"`
	ExampleSentence  string `json:"example_sentence"`
	Note             string `json:"note"`
	LessonID         *uint  `json:"lesson_id"`
	TranscriptID     *uint  `json:"transcript_id"`
}

type UpdateVocabularyItemReq struct {
	Phrase           *string `json:"phrase"`
	NormalizedPhrase *string `json:"normalized_phrase"`
	Meaning          *string `json:"meaning"`
	ExampleSentence  *string `json:"example_sentence"`
	Note             *string `json:"note"`
}
