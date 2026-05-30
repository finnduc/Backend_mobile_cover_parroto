-- +goose Up
-- +goose StatementBegin

CREATE TABLE global_chat_messages (
    id BIGSERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_global_chat_messages_created_at ON global_chat_messages(created_at DESC);
CREATE INDEX idx_global_chat_messages_user_id ON global_chat_messages(user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS global_chat_messages;
-- +goose StatementEnd
