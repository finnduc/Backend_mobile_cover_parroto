-- +goose Up
-- +goose StatementBegin
CREATE TABLE dictation_status (
    user_id VARCHAR(255) NOT NULL,
    transcript_id INTEGER NOT NULL REFERENCES transcripts(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, transcript_id)
);
CREATE INDEX idx_dictation_status_user_lesson ON dictation_status(user_id, lesson_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dictation_status;
-- +goose StatementEnd
