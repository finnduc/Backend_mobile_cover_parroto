-- +goose Up
-- +goose StatementBegin
SELECT setval(pg_get_serial_sequence('categories', 'id'), COALESCE(MAX(id), 1)) FROM categories;
SELECT setval(pg_get_serial_sequence('lessons', 'id'), COALESCE(MAX(id), 1)) FROM lessons;
SELECT setval(pg_get_serial_sequence('transcripts', 'id'), COALESCE(MAX(id), 1)) FROM transcripts;
SELECT setval(pg_get_serial_sequence('vocabulary_categories', 'id'), COALESCE(MAX(id), 1)) FROM vocabulary_categories;
SELECT setval(pg_get_serial_sequence('vocabulary_decks', 'id'), COALESCE(MAX(id), 1)) FROM vocabulary_decks;
SELECT setval(pg_get_serial_sequence('vocabulary_items', 'id'), COALESCE(MAX(id), 1)) FROM vocabulary_items;
SELECT setval(pg_get_serial_sequence('lesson_bookmarks', 'id'), COALESCE(MAX(id), 1)) FROM lesson_bookmarks;
SELECT setval(pg_get_serial_sequence('transcript_bookmarks', 'id'), COALESCE(MAX(id), 1)) FROM transcript_bookmarks;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- No rollback action needed for sequence synchronization
-- +goose StatementEnd
