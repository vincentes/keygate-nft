-- +goose Up
-- +goose StatementBegin
CREATE TABLE `Answer` (
    `ID` VARCHAR(255) NOT NULL,
    `QuestionID` VARCHAR(255) NOT NULL,
    `Answer` VARCHAR(255) NOT NULL,
    `IsCorrect` tinyint(1) NOT NULL,
    `CreatedAt` datetime NOT NULL default(current_timestamp),
    `UpdatedAt` datetime NOT NULL default(current_timestamp) on update current_timestamp,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`QuestionID`) REFERENCES `Question` (`ID`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `Answer`;
-- +goose StatementEnd
