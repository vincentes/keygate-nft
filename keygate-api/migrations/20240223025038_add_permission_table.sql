-- +goose Up
-- +goose StatementBegin
CREATE TABLE `Permission` (
  `ID` VARCHAR(255) PRIMARY KEY,
  `Name` VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `Permission`;
-- +goose StatementEnd
