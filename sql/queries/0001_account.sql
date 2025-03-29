-- name: GetAccountById :one
SELECT id, number, name,username, email, status,images,salt, password
FROM `account`
WHERE id = ? AND is_deleted = false;

-- name: GetOneAccountInfoAdmin :one
SELECT id, number, name, email,username, password,salt,status,create_at,update_at, images
FROM `account`
WHERE username = ? AND is_deleted = false;


-- name: GetAllAccounts :many
SELECT id, number, name, email,username, status, images
FROM `account`
WHERE is_deleted = false;

-- name: InsertAccount :execresult
INSERT INTO `account` (
    id,
    number,
    username,
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
VALUES(?,?,?,?,?,?,?,?,?,false,NOW(),NOW());

-- name: EditAccountById :exec
UPDATE account 
SET
    name = ?,
    username = ?,
    email = ?,
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

-- name: CheckAccountUserNameExists :one
SELECT COUNT(*)
FROM `account`
WHERE username = ?;