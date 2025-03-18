-- name: GetAccountById :one
SELECT id, number, name, email, status,images
FROM `account`
WHERE id = ? AND is_deleted = false;

-- name: GetOneAccountInfoAdmin :one
SELECT id, number, name, email, password,salt,status,create_at,update_at, images
FROM `account`
WHERE email = ? AND is_deleted = false;


-- name: GetAllAccounts :many
SELECT id, number, name, email, status, images
FROM `account`
WHERE is_deleted = false;

-- name: InsertAccount :execresult
INSERT INTO `account` (
    id,
    number,
    name,
    email,
    password,
    salt,
    status,
    images,
    is_deleted,
    create_at,
    update_at
)
VALUES(?,?,?,?,?,?,?,?,false,NOW(),NOW());

-- name: EditAccountById :exec
UPDATE account 
SET
    name = ?,
    email = ?,
    password = ?,
    status = ?,
    images = ?,
    update_at = NOW()
WHERE id = ?;

-- name: ChangPasswordById :exec
UPDATE account 
SET
    password = ?,
    update_at = NOW()
WHERE id = ?;

-- name: DeleteAccountById :exec
UPDATE account 
SET
    is_deleted = true,
    update_at = NOW()
WHERE id = ?;


-- name: CheckAccountBaseExists :one
SELECT COUNT(*)
FROM `account`
WHERE email = ?;