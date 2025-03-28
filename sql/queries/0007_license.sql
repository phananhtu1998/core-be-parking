-- name: GetLicenseById :one
SELECT id, license, date_start, date_end, created_at, update_at
FROM `license`
WHERE id = ? AND is_deleted = false;

-- name: GetAllLicenses :many
SELECT id, license, date_start, date_end, created_at, update_at
FROM `license`
WHERE is_deleted = false;

-- name: CreateLicense :execresult
INSERT INTO `license` (id,license, date_start, date_end, created_at, update_at, is_deleted)
    VALUES (?,?, ?, ?, NOW(), NOW(), false);

-- name: UpdateLicense :exec
UPDATE license
SET
    license = ?,
    date_start = ?,
    date_end = ?,
    update_at = NOW()
WHERE id = ?;

-- name: DeleteLicense :exec
UPDATE license
SET is_deleted = true,
    update_at = NOW()
WHERE id = ?;