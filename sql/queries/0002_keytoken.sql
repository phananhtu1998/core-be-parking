-- name: InsertKey :exec
INSERT INTO `keytoken` (id, account_id, refresh_token, refresh_tokens_used) 
VALUES (?, ?, ?, COALESCE(?, JSON_ARRAY())); -- Nếu NULL, đặt giá trị mặc định là JSON_ARRAY()

-- name: GetKeyByAccountID :one
SELECT id, account_id, refresh_token, refresh_tokens_used 
FROM `keytoken` 
WHERE account_id = ?;

-- name: UpdateRefreshToken :exec
UPDATE `keytoken` 
SET refresh_token = ?, update_at = CURRENT_TIMESTAMP
WHERE account_id = ?;

-- name: AddUsedToken :exec
UPDATE `keytoken`
SET refresh_tokens_used = CASE
    WHEN refresh_tokens_used IS NULL THEN JSON_ARRAY(?)
    ELSE JSON_ARRAY_APPEND(refresh_tokens_used, '$', ?)
END,
update_at = CURRENT_TIMESTAMP
WHERE account_id = ?;

-- name: UpdateRefreshTokenAndUsedTokens :exec
UPDATE `keytoken`
SET refresh_token = ?, 
    refresh_tokens_used = CASE
        WHEN refresh_tokens_used IS NULL THEN JSON_ARRAY(?)
        ELSE JSON_ARRAY_APPEND(refresh_tokens_used, '$', ?)
    END,
    update_at = CURRENT_TIMESTAMP
WHERE account_id = ?;

-- name: DeleteKey :exec
DELETE FROM `keytoken` WHERE account_id = ?;

-- name: CountByTokenAndAccount :one
SELECT COUNT(*) AS total_count
FROM keytoken
WHERE JSON_CONTAINS(refresh_tokens_used, ?, '$') 
AND account_id = ?;

-- name: CountByAccount :one
SELECT COUNT(*) AS total_count
FROM keytoken
WHERE account_id = ?;




