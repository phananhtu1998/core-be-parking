-- name: GetMenuById :one
SELECT id, menu_name, menu_icon, menu_url, menu_parent_Id, menu_level, 
       menu_number_order, menu_group_name, create_at, update_at
FROM menu
WHERE id = ? AND is_deleted = false;

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
-- Bước 1: Lưu giá trị menu_parent_Id và menu_number_order hiện tại
SET @old_parent_Id = (SELECT menu_parent_Id FROM menu WHERE id = ?);
SET @old_order = (SELECT menu_number_order FROM menu WHERE id = ?);

-- Bước 2: Nếu menu_parent_Id thay đổi, cập nhật lại thứ tự của menu cũ
UPDATE menu
SET menu_number_order = menu_number_order - 1
WHERE menu_parent_Id = @old_parent_Id
AND menu_number_order > @old_order;

-- Bước 3: Lấy số thứ tự lớn nhất trong menu cha mới
SET @new_order = (
    SELECT COALESCE(MAX(menu_number_order), 0) + 1 
    FROM menu 
    WHERE menu_parent_Id = ?
);

-- Bước 4: Cập nhật thông tin menu
UPDATE menu 
SET menu_name = ?, 
    menu_icon = ?, 
    menu_url = ?, 
    menu_parent_Id = ?, 
    menu_level = ?, 
    menu_number_order = IF(menu_parent_Id <> @old_parent_Id, @new_order, menu_number_order),
    menu_group_name = ?, 
    update_at = NOW()
WHERE id = ?;


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








