-- name: GetRoleById :one
SELECT id, code, role_name,role_left_value,role_right_value,role_max_number,
is_licensed,created_by,create_at,update_at
FROM `role`
WHERE id = ? AND is_deleted = false;

-- name: GetAllRole :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number,
is_licensed, created_by, create_at, update_at
FROM `role`
WHERE is_deleted = false
ORDER BY role_left_value ASC;

-- name: GetParentRoleInfo :one
SELECT role_left_value, role_right_value
FROM `role`
WHERE id = ? AND is_deleted = false;

-- name: UpdateRoleTree :execresult
UPDATE `role` SET role_right_value = role_right_value + 2 WHERE role_right_value >= ? AND is_deleted = false;
UPDATE `role` SET role_left_value = role_left_value + 2 WHERE role_left_value > ? AND is_deleted = false;

-- name: UpdateRole :exec
UPDATE `role`
SET code = ?, role_name = ?,role_left_value = ?,role_right_value = ?,role_max_number = ?,
is_licensed = ?,created_by = ?,update_at = NOW()
WHERE id = ?;

-- name: DeleteRole :exec
UPDATE `role`
SET is_deleted = true, update_at = ?
WHERE id = ?;

-- name: UpdateRightValuesForInsert :exec
UPDATE `role` 
SET role_right_value = role_right_value + 2 
WHERE role_right_value >= ? AND is_deleted = false;

-- name: UpdateLeftValuesForInsert :exec
UPDATE `role` 
SET role_left_value = role_left_value + 2 
WHERE role_left_value > ? AND is_deleted = false;

-- name: CreateRole :execresult
INSERT INTO `role` (
  id, code, role_name, role_left_value, role_right_value, 
  role_max_number, is_licensed, created_by, create_at, update_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetMaxRightValue :one
SELECT COALESCE(MAX(role_right_value), 0) as max_right_value
FROM `role`
WHERE is_deleted = false;

