-- name: GetAccountById :one
SELECT id, number, name,username, email, status,images,salt,created_by, password
FROM `account`
WHERE id = ? AND is_deleted = false;

-- name: GetOneAccountInfoAdmin :one
SELECT id, number, name, email,username, password,salt,status,created_by,create_at,update_at, images
FROM `account`
WHERE username = ? AND is_deleted = false;


-- name: GetAllAccounts :many
SELECT id, number, name, email,username, status, images,created_by
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
    created_by,
    is_deleted,
    create_at,
    update_at
)
VALUES(?,?,?,?,?,?,?,?,?,?,false,NOW(),NOW());

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
WHERE email = ? || username = ?;

-- name: CheckAccountUserNameExists :one
SELECT COUNT(*)
FROM `account`
WHERE username = ?;

-- name: CheckAccountExists :one
SELECT COUNT(*) FROM `account`;

-- name: GetLicenseByAccountId :one
SELECT a.id,r.role_name,l.license
FROM account as a
JOIN role_account ra ON a.id = ra.account_id
JOIN role r ON ra.role_id = r.id
JOIN license l ON l.id = r.license_id
WHERE a.id = ? AND a.is_deleted = false AND ra.is_deleted = false AND r.is_deleted = false AND l.is_deleted = false;


-- name: UpdateRoleAccountByAccountId :exec
UPDATE `role_account`
SET role_id = ?
WHERE account_id = ? AND is_deleted = false;


-- name: DeleteRoleAccountByAccountId :exec
UPDATE `role_account`
SET is_deleted = true, update_at = NOW()
WHERE account_id = ?;

-- name: GetAllAccountByCreatedBy :many
SELECT id, number, name,username, email, status,images,salt,created_by, password
FROM `account`
WHERE created_by = ? AND is_deleted = false;