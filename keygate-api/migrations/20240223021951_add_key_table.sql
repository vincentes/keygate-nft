-- +goose Up
-- +goose StatementBegin
CREATE TABLE `Key` (
  `ID` VARCHAR(255) PRIMARY KEY,
  `Name` VARCHAR(255)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE `Key`;
-- +goose StatementEnd
