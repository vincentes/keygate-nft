-- +goose Up
-- +goose StatementBegin
CREATE TABLE `Question` (
    `ID` VARCHAR(255) NOT NULL PRIMARY KEY,
    `QuizID` VARCHAR(255) NOT NULL,
    `Content` VARCHAR(255) NOT NULL,
    `CreatedAt` datetime NOT NULL default(current_timestamp),
    `UpdatedAt` datetime NOT NULL default(current_timestamp) on update current_timestamp,
    FOREIGN KEY (`QuizID`) REFERENCES `Quiz` (`ID`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `Question`;
-- +goose StatementEnd
