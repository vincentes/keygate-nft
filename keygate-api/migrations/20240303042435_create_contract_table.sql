-- +goose Up
-- +goose StatementBegin

CREATE TABLE `Contract` (
  ID VARCHAR(255) PRIMARY KEY,
  `Name` VARCHAR(255) NOT NULL,
  `Description` VARCHAR(255) NOT NULL,
  `Type` VARCHAR(255) NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  `Address` VARCHAR(255) NOT NULL
);

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE `Contract`;

-- +goose StatementEnd
