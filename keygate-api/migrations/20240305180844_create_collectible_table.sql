-- +goose Up
-- +goose StatementBegin
CREATE TABLE `Collectible` (
    `ID` varchar(255) NOT NULL,
    `Name` varchar(255) NOT NULL,
    `Description` text NOT NULL,
    `Image` varchar(255) NOT NULL,
    `CreatedAt` datetime NOT NULL default(current_timestamp),
    `UpdatedAt` datetime NOT NULL default(current_timestamp) on update current_timestamp,
    PRIMARY KEY (`ID`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `Collectible`;
-- +goose StatementEnd
