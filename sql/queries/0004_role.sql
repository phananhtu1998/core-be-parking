-- name: GetRoleById :one
SELECT id, code, role_name,role_left_value,role_right_value,role_max_number,
is_licensed,created_by,create_at,update_at
FROM `role`
WHERE id = ? AND is_deleted = false;

-- name: GetAllRole :many
SELECT id, code, role_name,role_left_value,role_right_value,role_max_number,
is_licensed,created_by,create_at,update_at
FROM `role`
WHERE is_deleted = false;

-- name: CreateRole :execresult
INSERT INTO `role` (code, role_name,role_left_value,role_right_value,role_max_number,
is_licensed,created_by,create_at,update_at)
VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW());

-- name: UpdateRole :exec
UPDATE `role`
SET code = ?, role_name = ?,role_left_value = ?,role_right_value = ?,role_max_number = ?,
is_licensed = ?,created_by = ?,create_at = ?,update_at = NOW()
WHERE id = ?;

-- name: DeleteRole :exec
UPDATE `role`
SET is_deleted = true, update_at = ?
WHERE id = ?;

