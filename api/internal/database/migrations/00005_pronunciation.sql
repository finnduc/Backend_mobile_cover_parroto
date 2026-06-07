-- +goose Up
-- +goose StatementBegin
CREATE TABLE pronunciation_attempts (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    transcript_id INTEGER NOT NULL REFERENCES transcripts(id) ON DELETE CASCADE,
    reference_text TEXT NOT NULL,
    overall_score DECIMAL(5,2) CHECK (overall_score >= 0 AND overall_score <= 100),
    accuracy_score DECIMAL(5,2) CHECK (accuracy_score >= 0 AND accuracy_score <= 100),
    fluency_score DECIMAL(5,2) CHECK (fluency_score >= 0 AND fluency_score <= 100),
    completeness_score DECIMAL(5,2) CHECK (completeness_score >= 0 AND completeness_score <= 100),
    prosody_score DECIMAL(5,2) CHECK (prosody_score >= 0 AND prosody_score <= 100),
    feedback TEXT,
    word_results JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pronunciation_progress (
    user_id VARCHAR(255) NOT NULL,
    transcript_id INTEGER NOT NULL REFERENCES transcripts(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    best_attempt_id INTEGER REFERENCES pronunciation_attempts(id) ON DELETE SET NULL,
    best_score DECIMAL(5,2) CHECK (best_score >= 0 AND best_score <= 100),
    feedback TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, transcript_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pronunciation_progress;
DROP TABLE IF EXISTS pronunciation_attempts;
-- +goose StatementEnd
