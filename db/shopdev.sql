CREATE TABLE `account` (
    `id` CHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `salt` VARCHAR(255) NOT NULL,
    `status` BOOLEAN NOT NULL COMMENT '[active,inactive]',
    `images` VARCHAR(255) NOT NULL,
    `is_deleted` BOOLEAN NOT NULL DEFAULT '0',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='account';

-- bảng lưu trữ refresh token theo id account
CREATE TABLE `keytoken` (
    `id` CHAR(36) NOT NULL,
    `account_id` CHAR(36) NOT NULL,
    `refresh_token` VARCHAR(255) NOT NULL,
    `refresh_tokens_used` JSON DEFAULT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX idx_account_id (`account_id`) -- Thêm index cho account_id
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='keytoken';

