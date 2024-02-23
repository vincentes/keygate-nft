-- +goose Up
-- +goose StatementBegin
ALTER TABLE User ADD CONSTRAINT UK_ExternalID UNIQUE (ExternalID);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE User DROP CONSTRAINT UK_ExternalID;
-- +goose StatementEnd
