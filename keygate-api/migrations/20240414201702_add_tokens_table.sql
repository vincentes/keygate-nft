-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tokens (
    ID VARCHAR(255) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    Address VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Tokens;
-- +goose StatementEnd