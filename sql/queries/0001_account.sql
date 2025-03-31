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

-- name: GetLicenseByAccountId :one
SELECT a.id,r.role_name,r.is_licensed,l.license
FROM account as a
JOIN role_account ra ON a.id = ra.account_id
JOIN role r ON ra.role_id = r.id
JOIN license l ON l.id = ra.license_id
WHERE a.id = ? AND a.is_deleted = false AND ra.is_deleted = false AND r.is_deleted = false AND l.is_deleted = false;