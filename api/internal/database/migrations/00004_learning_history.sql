-- +goose Up
-- +goose StatementBegin
CREATE TABLE learning_history (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    completed_dictation BOOLEAN DEFAULT FALSE,
    completed_pronunciation BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, lesson_id)
);
CREATE INDEX idx_learning_history_user_id ON learning_history(user_id);
CREATE INDEX idx_learning_history_lesson_id ON learning_history(lesson_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS learning_history;
-- +goose StatementEnd
