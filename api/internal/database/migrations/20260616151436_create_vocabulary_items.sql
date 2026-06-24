-- +goose Up
-- +goose StatementBegin
CREATE TABLE vocabulary_items (
    id SERIAL PRIMARY KEY,
    deck_id INTEGER NOT NULL REFERENCES vocabulary_decks(id) ON DELETE CASCADE,
    lesson_id INTEGER REFERENCES lessons(id) ON DELETE SET NULL,
    transcript_id INTEGER REFERENCES transcripts(id) ON DELETE SET NULL,
    phrase VARCHAR(500) NOT NULL,
    normalized_phrase VARCHAR(500) NOT NULL,
    meaning TEXT NOT NULL,
    example_sentence TEXT,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(deck_id, normalized_phrase)
);
CREATE INDEX idx_vocabulary_items_deck_id ON vocabulary_items(deck_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vocabulary_items;
-- +goose StatementEnd
