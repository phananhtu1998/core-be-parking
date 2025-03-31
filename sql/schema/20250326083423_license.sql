-- +goose Up
-- +goose StatementBegin
CREATE TABLE `license`(
    `id` CHAR(36) NOT NULL,
    `license` TEXT NOT NULL,
    `role_id` CHAR(36) NOT NULL,
    `date_start` DATETIME NOT NULL,
    `date_end` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `update_at` TIMESTAMP NOT NULL,
    `is_deleted` BOOLEAN NOT NULL,
    PRIMARY KEY(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `license`;
-- +goose StatementEnd
