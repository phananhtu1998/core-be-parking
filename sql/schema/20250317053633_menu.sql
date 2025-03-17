-- +goose Up
-- +goose StatementBegin
CREATE TABLE `menu`(
    `id` CHAR(36) NOT NULL,
    `menu_name` VARCHAR(255) NOT NULL,
    `menu_icon` VARCHAR(255) NOT NULL,
    `menu_url` VARCHAR(255) NOT NULL,
    `menu_parent_Id` CHAR(36) DEFAULT NULL,
    `menu_level` INT NOT NULL,
    `menu_number_order` FLOAT(53) NOT NULL,
    `menu_group_name` VARCHAR(255) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL,
    `create_at` TIMESTAMP NOT NULL,
    `update_at` TIMESTAMP NOT NULL,
    PRIMARY KEY(`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `menu`;
-- +goose StatementEnd
