-- +goose Up
-- +goose StatementBegin
CREATE TABLE `role_account`(
    `id` CHAR(36) NOT NULL,
    `account_id` CHAR(36) NOT NULL,
    `role_id` CHAR(36) NOT NULL,
    `license_id` CHAR(36) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `role_account`;
-- +goose StatementEnd
