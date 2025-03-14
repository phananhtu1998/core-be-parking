-- +goose Up
-- +goose StatementBegin
CREATE TABLE `keytoken` (
    `id` CHAR(36) NOT NULL,
    `account_id` CHAR(36) NOT NULL,
    `refresh_token` TEXT NOT NULL,
    `refresh_tokens_used` JSON DEFAULT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX idx_account_id (`account_id`) -- Thêm index cho account_id
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='keytoken';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `keytoken`;
-- +goose StatementEnd
