-- +goose Up
-- +goose StatementBegin
CREATE TABLE `role`(
    `id` CHAR(36) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `role_name` VARCHAR(255) NOT NULL,
    `role_left_value` INT NOT NULL,
    `role_right_value` INT NOT NULL,
    `role_max_number` INT NOT NULL,
    `created_by` VARCHAR(255) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL DEFAULT '0',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `role`;
-- +goose StatementEnd
