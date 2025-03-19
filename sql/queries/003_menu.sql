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
WHERE m1.menu_parent_Id = '' AND m1.is_deleted = false
GROUP BY m1.id 
ORDER BY m1.menu_number_order ASC;

-- name: InsertMenu :execresult
INSERT INTO menu (
    id, menu_name, menu_icon, menu_url, menu_parent_Id, menu_level, 
    menu_number_order, menu_group_name, is_deleted, create_at, update_at
) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, false, NOW(), NOW());


-- name: EditMenuById :exec
WITH old_values AS (
    SELECT menu_parent_Id, menu_number_order 
    FROM menu 
    WHERE menu.id = ?
),
new_order AS (
    SELECT COALESCE(MAX(menu.menu_number_order), 0) + 1 AS max_order 
    FROM menu 
    WHERE menu.menu_parent_Id = ?
)
UPDATE menu
SET 
    menu.menu_name = ?, 
    menu.menu_icon = ?, 
    menu.menu_url = ?, 
    menu.menu_parent_Id = ?, 
    menu.menu_number_order = (SELECT max_order FROM new_order),
    menu.menu_group_name = ?,
    menu.update_at = NOW()
WHERE menu.id = ?;

-- name: UpdateMenu :exec
WITH updated_parent AS (
    SELECT id, menu_level
    FROM menu
    WHERE id = ?
),
updated_children AS (
    SELECT id
    FROM menu
    WHERE menu_parent_id = (SELECT id FROM updated_parent)
)
UPDATE menu
SET 
    menu_name = CASE WHEN ? THEN ? ELSE menu_name END,
    menu_icon = CASE WHEN ? THEN ? ELSE menu_icon END,
    menu_url = CASE WHEN ? THEN ? ELSE menu_url END,
    menu_group_name = CASE WHEN ? THEN ? ELSE menu_group_name END,
    menu_level = CASE WHEN ? THEN ? ELSE menu_level END,
    update_at = NOW()
WHERE id = ?;

UPDATE menu
SET 
    menu_level = (SELECT menu_level FROM updated_parent WHERE id = ?),
    update_at = NOW()
WHERE menu_parent_id = ?;

UPDATE menu
SET 
    menu_number_order = ROW_NUMBER() OVER (ORDER BY menu_number_order ASC),
    update_at = NOW()
WHERE menu_parent_id IN (SELECT id FROM updated_children);

-- name: DeleteMenuById :exec 

WITH RECURSIVE to_delete AS (
    SELECT m.id FROM menu m WHERE m.id = ? 
    UNION ALL
    SELECT m.id FROM menu m
    INNER JOIN to_delete td ON m.menu_parent_Id = td.id
)
-- Đánh dấu xóa menu cha + con
UPDATE menu 
SET is_deleted = true, update_at = NOW()
WHERE id IN (SELECT td.id FROM to_delete td);

-- Xóa menu cha nếu tất cả con đã bị xóa
UPDATE menu m_parent
SET is_deleted = true, update_at = NOW()
WHERE m_parent.id IN (
    SELECT DISTINCT m.menu_parent_Id 
    FROM menu m 
    WHERE m.menu_parent_Id IS NOT NULL 
    GROUP BY m.menu_parent_Id
    HAVING SUM(IF(m.is_deleted = false, 1, 0)) = 0
);

-- name: CountMenuByURL :one
SELECT COUNT(*) AS total_count
FROM menu
WHERE menu_url = ?;








