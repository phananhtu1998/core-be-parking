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


-- name: GetRoleMenuByRoleId :many
SELECT 
    m.id, 
    m.menu_name, 
    m.menu_url, 
    m.menu_icon, 
    rm.menu_id,
    rm.role_id,
    m.menu_group_name, 
    r.code, 
    r.role_name, 
    rm.list_method 
FROM roles_menu rm
JOIN menu m ON m.id = rm.menu_id AND m.is_deleted = FALSE
JOIN role r ON r.id = rm.role_id AND r.is_deleted = FALSE
WHERE r.id = ?
AND (
    ? = '' OR MATCH(r.role_name) AGAINST (? IN NATURAL LANGUAGE MODE)
);






