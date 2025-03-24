-- name: GetAllRolesMenu :many
SELECT id, menu_id, role_id, list_method
FROM `roles_menu`
WHERE is_deleted = false;

-- name: GetRolesMenuByRoleId :many
SELECT id, menu_id, role_id, list_method
FROM `roles_menu`
WHERE role_id = ? AND is_deleted = false;

-- name: CreateRolesMenu :exec
INSERT INTO `roles_menu` (id, menu_id, role_id, list_method, is_deleted, create_at, update_at)
VALUES (?, ?, ?, ?, false, NOW(), NOW());

-- name: UpdateRolesMenu :exec
UPDATE `roles_menu`
SET menu_id = ?, role_id = ?, list_method = ?
WHERE id = ?;

-- name: DeleteRolesMenu :exec
UPDATE `roles_menu`
SET is_deleted = true, update_at = ?
WHERE id = ?;




