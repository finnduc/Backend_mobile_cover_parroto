-- +goose Up
-- +goose StatementBegin
CREATE TABLE learning_history (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    duration_watched DOUBLE PRECISION,
    completed BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_learning_history_user_id ON learning_history(user_id);
CREATE INDEX idx_learning_history_lesson_id ON learning_history(lesson_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS learning_history;
-- +goose StatementEnd
