-- +goose Up
-- +goose StatementBegin

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);
CREATE INDEX idx_categories_name ON categories(name);

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    title VARCHAR(255),
    description TEXT,
    video_url VARCHAR(500) NOT NULL,
    thumbnail_url VARCHAR(500),
    level VARCHAR(20),
    duration DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_lessons_category_id ON lessons(category_id);

CREATE TABLE transcripts (
    id SERIAL PRIMARY KEY,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    sequence INTEGER NOT NULL,
    content TEXT NOT NULL,
    phonetic VARCHAR(500),
    vietnamese TEXT,
    start_timestamp DOUBLE PRECISION,
    end_timestamp DOUBLE PRECISION
);
CREATE INDEX idx_transcripts_lesson_id ON transcripts(lesson_id);

CREATE TABLE bookmarks (
    user_id VARCHAR(255) NOT NULL,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, lesson_id)
);
CREATE INDEX idx_bookmarks_user_id ON bookmarks(user_id);
CREATE INDEX idx_bookmarks_lesson_id ON bookmarks(lesson_id);

CREATE TABLE shadowing_status (
    user_id VARCHAR(255) NOT NULL,
    transcript_id INTEGER NOT NULL REFERENCES transcripts(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, transcript_id)
);
CREATE INDEX idx_shadowing_status_user_lesson ON shadowing_status(user_id, lesson_id);

CREATE TABLE dictation_status (
    user_id VARCHAR(255) NOT NULL,
    transcript_id INTEGER NOT NULL REFERENCES transcripts(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, transcript_id)
);
CREATE INDEX idx_dictation_status_user_lesson ON dictation_status(user_id, lesson_id);

CREATE TABLE vocabulary_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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
DROP TABLE IF EXISTS vocabulary_decks;
DROP TABLE IF EXISTS vocabulary_categories;
DROP TABLE IF EXISTS dictation_status;
DROP TABLE IF EXISTS shadowing_status;
DROP TABLE IF EXISTS bookmarks;
DROP TABLE IF EXISTS transcripts;
DROP TABLE IF EXISTS lessons;
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
