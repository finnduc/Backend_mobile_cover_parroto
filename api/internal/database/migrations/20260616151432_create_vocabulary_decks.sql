-- +goose Up
-- +goose StatementBegin
CREATE TABLE vocabulary_decks (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255),
    category_id INTEGER REFERENCES vocabulary_categories(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    thumbnail_url VARCHAR(500),
    level VARCHAR(20),
    is_default BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_vocabulary_decks_user_id ON vocabulary_decks(user_id);
CREATE INDEX idx_vocabulary_decks_category_id ON vocabulary_decks(category_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vocabulary_decks;
-- +goose StatementEnd
