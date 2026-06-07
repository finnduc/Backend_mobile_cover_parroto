-- +goose Up
-- +goose StatementBegin
CREATE TABLE transcript_progress (
    user_id VARCHAR(255) NOT NULL,
    transcript_id INTEGER NOT NULL REFERENCES transcripts(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, transcript_id)
);
CREATE INDEX idx_transcript_progress_user_lesson ON transcript_progress(user_id, lesson_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transcript_progress;
-- +goose StatementEnd
