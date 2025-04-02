-- name: GetRoleById :one
SELECT id, code, role_name,role_left_value,role_right_value,role_max_number,created_by,create_at,update_at
FROM `role`
WHERE id = ? AND is_deleted = false;


-- name: GetAllRole :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number, created_by, create_at, update_at
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
SET code = ?, role_name = ?,role_left_value = ?,role_right_value = ?,role_max_number = ?,created_by = ?,update_at = NOW()
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
  role_max_number, created_by, create_at, update_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetMaxRightValue :one
SELECT COALESCE(MAX(role_right_value), 0) as max_right_value
FROM `role`
WHERE is_deleted = false;

-- name: GetRoleWithChildren :many
SELECT r.id, r.code, r.role_name, r.role_left_value, r.role_right_value, r.role_max_number, r.created_by, r.create_at, r.update_at
FROM `role` r
WHERE r.role_left_value >= (SELECT r2.role_left_value FROM `role` r2 WHERE r2.id = ? AND r2.is_deleted = false)
AND r.role_right_value <= (SELECT r3.role_right_value FROM `role` r3 WHERE r3.id = ? AND r3.is_deleted = false)
AND r.is_deleted = false
ORDER BY r.role_left_value ASC;

-- name: GetChildRolesByParentId :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number, created_by, create_at, update_at
FROM `role`
WHERE created_by = ? AND is_deleted = false
ORDER BY role_left_value ASC;

-- name: GetTotalRoles :one
SELECT COUNT(*) FROM `role` WHERE is_deleted = false;

-- name: GetRolesWithPagination :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number, created_by, create_at, update_at
FROM `role`
WHERE is_deleted = false
ORDER BY role_left_value ASC LIMIT ? OFFSET ?;

-- name: SoftDeleteRolesByRange :exec
UPDATE `role`
SET is_deleted = true, update_at = ?
WHERE role_left_value >= ? AND role_right_value <= ? AND is_deleted = false;

-- name: GetAllPermissions :many
SELECT a.id,a.name, r.role_name, m.menu_group_name,rm.list_method as Method FROM account a
JOIN role_account ra ON ra.account_id = a.id
JOIN role r ON r.id = ra.role_id
JOIN roles_menu rm ON rm.role_id = r.id
JOIN menu m ON m.id = rm.menu_id
WHERE a.is_deleted = false AND r.is_deleted = false AND m.is_deleted = false;

-- name: GetAllPermissionsByAccountId :many
SELECT a.id,a.name, r.role_name, m.menu_group_name,rm.list_method as Method FROM account a
JOIN role_account ra ON ra.account_id = a.id
JOIN role r ON r.id = ra.role_id
JOIN roles_menu rm ON rm.role_id = r.id
JOIN menu m ON m.id = rm.menu_id
WHERE a.is_deleted = false AND r.is_deleted = false AND m.is_deleted = false AND a.id = ?;

-- name: GetTotalAccounts :one
SELECT 
    SUM(CASE WHEN r.role_max_number REGEXP '^[0-9]+$' 
             THEN CAST(r.role_max_number AS UNSIGNED) 
             ELSE 0 
        END) AS TotalAccount,
    CASE 
        WHEN COUNT(CASE WHEN r.role_max_number = 'MAX' THEN 1 END) > 0 THEN 1
        ELSE 0
    END AS is_max
FROM `role` AS r
WHERE r.created_by = ?;


