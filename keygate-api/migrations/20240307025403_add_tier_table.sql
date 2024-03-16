-- +goose Up
-- +goose StatementBegin
CREATE TABLE `Tier` (
    `ID` VARCHAR(255) PRIMARY KEY,
    `Name` VARCHAR(255) NOT NULL,
    `Description` TEXT,
    `MinPoints` INT NOT NULL,
    `MaxPoints` INT NOT NULL,
    `Benefits` TEXT,
    `CreatedAt` DATETIME NOT NULL,
    `UpdatedAt` DATETIME NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `Tier`;
-- +goose StatementEnd