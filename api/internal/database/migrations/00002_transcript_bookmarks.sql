-- +goose Up
-- +goose StatementBegin
CREATE TABLE transcript_bookmarks (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    transcript_id INTEGER NOT NULL REFERENCES transcripts(id) ON DELETE CASCADE,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, transcript_id)
);

CREATE INDEX idx_transcript_bookmarks_user_id ON transcript_bookmarks(user_id);
CREATE INDEX idx_transcript_bookmarks_transcript_id ON transcript_bookmarks(transcript_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transcript_bookmarks;
-- +goose StatementEnd
