CREATE TABLE `account`(
    `id` CHAR(36) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `status` BOOLEAN NOT NULL COMMENT '[active,inactive]',
    `images` VARCHAR(255) NOT NULL,
    `created_by` VARCHAR(255) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL DEFAULT '0',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
CREATE TABLE `role`(
    `id` CHAR(36) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `role_name` VARCHAR(255) NOT NULL,
    `role_left_value` INT NOT NULL,
    `role_right_value` INT NOT NULL,
    `role_max_number` BIGINT NOT NULL,
    `is_licensed` BOOLEAN NOT NULL,
    `created_by` VARCHAR(255) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL DEFAULT '0',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
CREATE TABLE `menu`(
    `id` CHAR(36) NOT NULL,
    `menu_name` VARCHAR(255) NOT NULL,
    `menu_icon` VARCHAR(255) NOT NULL,
    `menu_url` VARCHAR(255) NOT NULL,
    `menu_parent_Id` CHAR(36) NOT NULL,
    `menu_level` INT NOT NULL,
    `menu_number_order` FLOAT(53) NOT NULL,
    `menu_group_name` VARCHAR(255) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL,
    `create_at` TIMESTAMP NOT NULL,
    `update_at` TIMESTAMP NOT NULL,
    PRIMARY KEY(`id`)
);
CREATE TABLE `RolesMenu`(
    `id` CHAR(36) NOT NULL,
    `menu_id` CHAR(36) NOT NULL,
    `role_id` CHAR(36) NOT NULL,
    `list_method` JSON NOT NULL,
    PRIMARY KEY(`id`)
);
CREATE TABLE `RoleAccount`(
    `id` CHAR(36) NOT NULL,
    `account_id` CHAR(36) NOT NULL,
    `role_id` CHAR(36) NOT NULL,
    `license_id` CHAR(36) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL,
    PRIMARY KEY(`id`)
);
CREATE TABLE `license`(
    `id` CHAR(36) NOT NULL,
    `license` VARCHAR(255) NOT NULL,
    `date_start` DATETIME NOT NULL,
    `date_end` DATETIME NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `update_at` TIMESTAMP NOT NULL,
    PRIMARY KEY(`id`)
);
ALTER TABLE
    `RolesMenu` ADD CONSTRAINT `rolesmenu_role_id_foreign` FOREIGN KEY(`role_id`) REFERENCES `role`(`id`);
ALTER TABLE
    `RoleAccount` ADD CONSTRAINT `roleaccount_role_id_foreign` FOREIGN KEY(`role_id`) REFERENCES `role`(`id`);
ALTER TABLE
    `RolesMenu` ADD CONSTRAINT `rolesmenu_menu_id_foreign` FOREIGN KEY(`menu_id`) REFERENCES `menu`(`id`);
ALTER TABLE
    `RoleAccount` ADD CONSTRAINT `roleaccount_account_id_foreign` FOREIGN KEY(`account_id`) REFERENCES `account`(`id`);
ALTER TABLE
    `RoleAccount` ADD CONSTRAINT `roleaccount_license_id_foreign` FOREIGN KEY(`license_id`) REFERENCES `license`(`id`);