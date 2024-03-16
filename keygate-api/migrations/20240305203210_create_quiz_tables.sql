-- +goose Up
-- +goose StatementBegin
CREATE TABLE `Quiz` (
    `ID` VARCHAR(255) NOT NULL,
    `Name` VARCHAR(255) NOT NULL,
    `Description` VARCHAR(255) NOT NULL,
    `Image` VARCHAR(255),
    `CreatedAt` datetime NOT NULL default(current_timestamp),
    `UpdatedAt` datetime NOT NULL default(current_timestamp) on update current_timestamp,
    PRIMARY KEY (`ID`)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `Quiz`;
-- DROP TABLE `Question`;
-- DROP TABLE `Answer`;
-- +goose StatementEnd



