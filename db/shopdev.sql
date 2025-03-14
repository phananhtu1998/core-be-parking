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
    `refresh_token` TEXT NOT NULL,
    `refresh_tokens_used` JSON DEFAULT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_account_id` (`account_id`) -- Đảm bảo account_id là duy nhất
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='keytoken';



-- trigger cho bảng keytoken 
CREATE TRIGGER before_insert_keytoken
BEFORE INSERT ON keytoken
FOR EACH ROW
BEGIN
    DECLARE account_count INT DEFAULT 0;

    -- Kiểm tra số lượng account_id đã tồn tại
    SELECT COUNT(*) INTO account_count FROM keytoken WHERE account_id = NEW.account_id;

    IF account_count > 0 THEN
        -- Cập nhật refresh_token và refresh_tokens_used
        UPDATE keytoken
        SET refresh_tokens_used = IFNULL(JSON_ARRAY_APPEND(refresh_tokens_used, '$', NEW.refresh_token), JSON_ARRAY(NEW.refresh_token)),
            refresh_token = NEW.refresh_token,
            update_at = NOW()
        WHERE account_id = NEW.account_id;

        -- Hủy INSERT bằng cách tạo lỗi giả lập (vì SIGNAL không dùng được)
        SET NEW.id = UUID(); -- Nếu id là NOT NULL, gán một giá trị UUID để MySQL hiểu là lỗi
    END IF;
END;
