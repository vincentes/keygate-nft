-- +goose Up
-- +goose StatementBegin
-- rename Collectible table to Collection

ALTER TABLE Collectible RENAME TO `Collection`;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE `Collection` RENAME TO Collectible;

-- +goose StatementEnd
