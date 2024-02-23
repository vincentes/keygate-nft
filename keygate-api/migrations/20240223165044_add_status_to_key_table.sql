-- +goose Up
-- +goose StatementBegin
ALTER TABLE `Key` ADD COLUMN `status` VARCHAR(255) NOT NULL DEFAULT 'inactive';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `Key` DROP COLUMN `status`;
-- +goose StatementEnd
