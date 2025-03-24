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
ORDER BY role_left_value DESC
LIMIT ? OFFSET ?;

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

-- name: GetRoleWithChildren :many
SELECT r.id, r.code, r.role_name, r.role_left_value, r.role_right_value, r.role_max_number,
r.is_licensed, r.created_by, r.create_at, r.update_at
FROM `role` r
WHERE r.role_left_value >= (SELECT r2.role_left_value FROM `role` r2 WHERE r2.id = ? AND r2.is_deleted = false)
AND r.role_right_value <= (SELECT r3.role_right_value FROM `role` r3 WHERE r3.id = ? AND r3.is_deleted = false)
AND r.is_deleted = false
ORDER BY r.role_left_value ASC;

-- name: GetChildRolesByParentId :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number,
is_licensed, created_by, create_at, update_at
FROM `role`
WHERE created_by = ? AND is_deleted = false
ORDER BY role_left_value ASC;

-- name: GetTotalRoles :one
SELECT COUNT(*) FROM `role` WHERE is_deleted = false;

-- name: GetRolesWithPagination :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number,
is_licensed, created_by, create_at, update_at
FROM `role`
WHERE is_deleted = false
ORDER BY role_left_value ASC LIMIT ? OFFSET ?;

-- name: SoftDeleteRolesByRange :exec
UPDATE `role`
SET is_deleted = true, update_at = ?
WHERE role_left_value >= ? AND role_right_value <= ? AND is_deleted = false;
