-- +goose Up
-- +goose StatementBegin
CREATE TABLE lesson_bookmarks (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, lesson_id)
);
CREATE INDEX idx_lesson_bookmarks_user_id ON lesson_bookmarks(user_id);
CREATE INDEX idx_lesson_bookmarks_lesson_id ON lesson_bookmarks(lesson_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lesson_bookmarks;
-- +goose StatementEnd
