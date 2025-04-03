-- name: GetAllRoleAccount :many
SELECT id, account_id, role_id
FROM `role_account`
WHERE is_deleted = false;

-- name: GetRoleAccountById :one
SELECT id, account_id, role_id
FROM `role_account`
WHERE is_deleted = false AND id = ?;

-- name: GetRoleAccountByAccountId :many
SELECT id, account_id, role_id, create_at, update_at
FROM `role_account`
WHERE is_deleted = false AND account_id = ?;

-- name: GetRoleAccountByRoleId :many
SELECT id, account_id, role_id, create_at, update_at
FROM `role_account`
WHERE is_deleted = false AND role_id = ?;

-- name: CreateRoleAccount :exec
INSERT INTO `role_account` (id, account_id, role_id,created_by, is_deleted, create_at, update_at)
VALUES (?, ?, ?, ?, false, NOW(), NOW());

-- name: UpdateRoleAccount :exec
UPDATE `role_account`
SET account_id = ?, role_id = ?
WHERE id = ? AND is_deleted = false;

-- name: DeleteRoleAccount :exec
UPDATE `role_account`
SET is_deleted = true, update_at = NOW()
WHERE id IN (?);


-- name: GetOneRoleAccountByAccountId :one
SELECT id, account_id, role_id, create_at, update_at
FROM `role_account`
WHERE is_deleted = false AND account_id = ?;

-- name: CheckCountRoleId :one
SELECT COUNT(*)
FROM `role_account`
WHERE role_id = ? AND is_deleted = false;