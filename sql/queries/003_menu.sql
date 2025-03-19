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
    m1.id, m1.menu_name, m1.menu_icon, m1.menu_url, m1.menu_parent_id, 
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
LEFT JOIN menu m2 ON m1.id = m2.menu_parent_id AND m2.is_deleted = false
WHERE (m1.menu_parent_id IS NULL OR m1.menu_parent_id = '') AND m1.is_deleted = false
GROUP BY m1.id 
ORDER BY m1.menu_number_order ASC;

-- name: InsertMenu :execresult
INSERT INTO menu (
    id, menu_name, menu_icon, menu_url, menu_parent_Id, menu_level, 
    menu_number_order, menu_group_name, is_deleted, create_at, update_at
) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, false, NOW(), NOW());

-- name: UpdateMenu :exec
UPDATE menu
SET 
    menu_name = sqlc.arg(menu_name),
    menu_icon = sqlc.arg(menu_icon),
    menu_url = sqlc.arg(menu_url),
    menu_number_order = sqlc.arg(menu_number_order),
    menu_group_name = sqlc.arg(menu_group_name),
    menu_level = sqlc.arg(menu_level),
    menu_parent_id = sqlc.arg(menu_parent_id),
    update_at = NOW()
WHERE id = sqlc.arg(id);

-- name: GetChildMenus :many
WITH RECURSIVE submenus AS (
    SELECT 
        m.id, 
        m.menu_name, 
        m.menu_icon, 
        m.menu_url, 
        m.menu_number_order, 
        m.menu_group_name, 
        m.menu_level, 
        m.menu_parent_id,  -- Chỉ định rõ từ `menu`
        m.update_at
    FROM menu m
    WHERE m.menu_parent_id = sqlc.arg(id)

    UNION ALL

    SELECT 
        m.id, 
        m.menu_name, 
        m.menu_icon, 
        m.menu_url, 
        m.menu_number_order, 
        m.menu_group_name, 
        m.menu_level, 
        m.menu_parent_id,  -- Chỉ định rõ từ `menu`
        m.update_at
    FROM menu m
    JOIN submenus s ON m.menu_parent_id = s.id  -- Chỉ định rõ từ `submenus`
)
SELECT 
    s.id, 
    s.menu_name, 
    s.menu_icon, 
    s.menu_url, 
    s.menu_number_order, 
    s.menu_group_name, 
    s.menu_level, 
    s.menu_parent_id,  -- Chỉ định rõ từ `submenus`
    s.update_at
FROM submenus s;


-- name: CountMenuByURL :one
SELECT COUNT(*) AS total_count
FROM menu
WHERE menu_url = ?;








