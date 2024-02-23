-- +goose Up
-- +goose StatementBegin
CREATE TABLE `KeyPermission` (
  `KeyID` VARCHAR(255),
  `PermissionID` VARCHAR(255),
  PRIMARY KEY (`KeyID`, `PermissionID`),
  FOREIGN KEY (`KeyID`) REFERENCES `Key` (`ID`) ON DELETE CASCADE,
  FOREIGN KEY (`PermissionID`) REFERENCES `Permission` (`ID`) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `KeyPermission`;
-- +goose StatementEnd
