-- +goose Up
-- +goose StatementBegin
CREATE TABLE `hello`(
    `id` CHAR(36) NOT NULL,
    `hello` VARCHAR(255) NOT NULL,
    `date_start` DATETIME NOT NULL,
    `date_end` DATETIME NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `update_at` TIMESTAMP NOT NULL,
    `is_deleted` BOOLEAN NOT NULL,
    PRIMARY KEY(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `hello`;
-- +goose StatementEnd
