-- +goose Up
-- +goose StatementBegin
ALTER TABLE shadowing_status ADD COLUMN best_score DOUBLE PRECISION DEFAULT 0.0;
ALTER TABLE shadowing_status ADD COLUMN feedback TEXT DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE shadowing_status DROP COLUMN IF EXISTS best_score;
ALTER TABLE shadowing_status DROP COLUMN IF EXISTS feedback;
-- +goose StatementEnd
