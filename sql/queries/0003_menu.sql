-- name: GetMenuById :one
SELECT 
    m1.id, m1.menu_name, m1.menu_icon, m1.menu_url, m1.menu_parent_Id, 
    m1.menu_level, m1.menu_number_order, m1.menu_group_name, m1.is_deleted, 
    m1.create_at, m1.update_at,
    COALESCE(
        CONCAT('[', GROUP_CONCAT(
            CASE 
                WHEN m2.id IS NOT NULL THEN 
                    JSON_OBJECT(
                        'id', m2.id, 
                        'menu_name', m2.menu_name, 
                        'menu_icon', m2.menu_icon, 
                        'menu_url', m2.menu_url, 
                        'menu_level', m2.menu_level,
                        'menu_number_order', m2.menu_number_order,
                        'menu_group_name', m2.menu_group_name,
                        'is_deleted', m2.is_deleted
                    )
                ELSE NULL
            END 
            ORDER BY m2.menu_number_order ASC SEPARATOR ','
        ), ']'), '[]'
    ) AS children
FROM menu m1
LEFT JOIN menu m2 ON m1.id = m2.menu_parent_Id AND m2.is_deleted = false
WHERE m1.id = ? AND m1.is_deleted = false
GROUP BY m1.id 
ORDER BY m1.menu_number_order ASC;

-- name: GetAllMenus :many
SELECT 
    m1.id, m1.menu_name, m1.menu_icon, m1.menu_url, m1.menu_parent_Id,
    m1.menu_level, m1.menu_number_order, m1.menu_group_name, m1.is_deleted
FROM menu m1
WHERE m1.is_deleted = false
ORDER BY m1.menu_level ASC, m1.menu_number_order ASC;

-- name: InsertMenu :execresult
INSERT INTO menu (
    id, menu_name, menu_icon, menu_url, menu_parent_Id, menu_level, 
    menu_number_order, menu_group_name, is_deleted, create_at, update_at
) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, false, NOW(), NOW());

-- name: GetMenusByIDs :many
SELECT id, menu_name, menu_icon, menu_url, menu_parent_id, menu_level, menu_number_order, menu_group_name, is_deleted
FROM menu
WHERE FIND_IN_SET(id, $1) AND is_deleted = false;

-- name: CountMenuByURL :one
SELECT COUNT(*) AS total_count
FROM menu
WHERE menu_url = ?;

-- name: UpdateSingleMenu :exec
UPDATE menu
SET 
    menu_name = ?, 
    menu_icon = ?, 
    menu_url = ?, 
    menu_parent_id = ?, 
    menu_level = ?, 
    menu_number_order = ?, 
    menu_group_name = ?
WHERE id = ?;

-- name: DeleteMenu :exec
UPDATE menu SET is_deleted = true WHERE id = ?;

-- name: UpdateMenuDeleted :exec
UPDATE menu SET is_deleted = true WHERE id = ?;


-- name: GetMenuByRoleId :many
SELECT 
    m.id,
    m.menu_name,
    m.menu_icon,
    m.menu_url,
    m.menu_level,
    m.menu_number_order,
    m.menu_parent_Id,
    m.menu_group_name,
    CASE 
        WHEN m.menu_parent_Id IS NULL 
        THEN CAST(m.menu_number_order AS CHAR(20))
        ELSE CONCAT(
            (SELECT CAST(parent.menu_number_order AS CHAR(20))
             FROM menu parent 
             WHERE parent.id = m.menu_parent_Id),
            '.',
            CAST(m.menu_number_order AS CHAR(20))
        )
    END AS stt
FROM `menu` m
JOIN `roles_menu` rm ON rm.menu_id = m.id
JOIN `role` r ON rm.role_id = r.id
WHERE r.id = ? 
  AND m.is_deleted = false 
  AND rm.is_deleted = false 
  AND r.is_deleted = false 
ORDER BY 
    CAST(
        SUBSTRING_INDEX(stt, '.', 1) AS UNSIGNED
    ) ASC,
    LENGTH(stt) ASC,
    stt ASC;







