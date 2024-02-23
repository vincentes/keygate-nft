-- +goose Up
-- +goose StatementBegin

CREATE TABLE `User` (
  ID VARCHAR(255) NOT NULL,
  ExternalID VARCHAR(255) NOT NULL,
  PRIMARY KEY (ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `User`;
-- +goose StatementEnd
