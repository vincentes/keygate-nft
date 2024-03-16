-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN points INTEGER DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
