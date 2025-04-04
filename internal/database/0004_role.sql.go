// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 0004_role.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createRole = `-- name: CreateRole :execresult
INSERT INTO ` + "`" + `role` + "`" + ` (
  id, code, role_name, role_left_value, role_right_value,license_id,
  role_max_number, created_by, create_at, update_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()
)
`

type CreateRoleParams struct {
	ID             string
	Code           string
	RoleName       string
	RoleLeftValue  int32
	RoleRightValue int32
	LicenseID      string
	RoleMaxNumber  int32
	CreatedBy      string
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createRole,
		arg.ID,
		arg.Code,
		arg.RoleName,
		arg.RoleLeftValue,
		arg.RoleRightValue,
		arg.LicenseID,
		arg.RoleMaxNumber,
		arg.CreatedBy,
	)
}

const deleteRole = `-- name: DeleteRole :exec
UPDATE ` + "`" + `role` + "`" + `
SET is_deleted = true, update_at = ?
WHERE id = ?
`

type DeleteRoleParams struct {
	UpdateAt time.Time
	ID       string
}

func (q *Queries) DeleteRole(ctx context.Context, arg DeleteRoleParams) error {
	_, err := q.db.ExecContext(ctx, deleteRole, arg.UpdateAt, arg.ID)
	return err
}

const getAccountCreated = `-- name: GetAccountCreated :one
SELECT COUNT(*) 
FROM ` + "`" + `role` + "`" + ` 
WHERE created_by = ? AND is_deleted = false
`

func (q *Queries) GetAccountCreated(ctx context.Context, createdBy string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountCreated, createdBy)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAllFuncPackageByCreatedBy = `-- name: GetAllFuncPackageByCreatedBy :many
SELECT id, code, role_name,role_max_number,create_at
FROM ` + "`" + `role` + "`" + `
WHERE is_deleted = false AND created_by = ? 
ORDER BY role_left_value DESC
`

type GetAllFuncPackageByCreatedByRow struct {
	ID            string
	Code          string
	RoleName      string
	RoleMaxNumber int32
	CreateAt      time.Time
}

func (q *Queries) GetAllFuncPackageByCreatedBy(ctx context.Context, createdBy string) ([]GetAllFuncPackageByCreatedByRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllFuncPackageByCreatedBy, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllFuncPackageByCreatedByRow
	for rows.Next() {
		var i GetAllFuncPackageByCreatedByRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.RoleName,
			&i.RoleMaxNumber,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPermissions = `-- name: GetAllPermissions :many
SELECT a.id,a.name, r.role_name, m.menu_group_name,rm.list_method as Method FROM account a
JOIN role_account ra ON ra.account_id = a.id
JOIN role r ON r.id = ra.role_id
JOIN roles_menu rm ON rm.role_id = r.id
JOIN menu m ON m.id = rm.menu_id
WHERE a.is_deleted = false AND r.is_deleted = false AND m.is_deleted = false
`

type GetAllPermissionsRow struct {
	ID            string
	Name          string
	RoleName      string
	MenuGroupName string
	Method        json.RawMessage
}

func (q *Queries) GetAllPermissions(ctx context.Context) ([]GetAllPermissionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllPermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllPermissionsRow
	for rows.Next() {
		var i GetAllPermissionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.RoleName,
			&i.MenuGroupName,
			&i.Method,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPermissionsByAccountId = `-- name: GetAllPermissionsByAccountId :many
SELECT a.id,a.name, r.role_name, m.menu_group_name,rm.list_method as Method FROM account a
JOIN role_account ra ON ra.account_id = a.id
JOIN role r ON r.id = ra.role_id
JOIN roles_menu rm ON rm.role_id = r.id
JOIN menu m ON m.id = rm.menu_id
WHERE a.is_deleted = false AND r.is_deleted = false AND m.is_deleted = false AND a.id = ?
`

type GetAllPermissionsByAccountIdRow struct {
	ID            string
	Name          string
	RoleName      string
	MenuGroupName string
	Method        json.RawMessage
}

func (q *Queries) GetAllPermissionsByAccountId(ctx context.Context, id string) ([]GetAllPermissionsByAccountIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllPermissionsByAccountId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllPermissionsByAccountIdRow
	for rows.Next() {
		var i GetAllPermissionsByAccountIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.RoleName,
			&i.MenuGroupName,
			&i.Method,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllRole = `-- name: GetAllRole :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number, created_by, create_at, update_at
FROM ` + "`" + `role` + "`" + `
WHERE is_deleted = false
ORDER BY role_left_value DESC
LIMIT ? OFFSET ?
`

type GetAllRoleParams struct {
	Limit  int32
	Offset int32
}

type GetAllRoleRow struct {
	ID             string
	Code           string
	RoleName       string
	RoleLeftValue  int32
	RoleRightValue int32
	RoleMaxNumber  int32
	CreatedBy      string
	CreateAt       time.Time
	UpdateAt       time.Time
}

func (q *Queries) GetAllRole(ctx context.Context, arg GetAllRoleParams) ([]GetAllRoleRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllRole, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllRoleRow
	for rows.Next() {
		var i GetAllRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.RoleName,
			&i.RoleLeftValue,
			&i.RoleRightValue,
			&i.RoleMaxNumber,
			&i.CreatedBy,
			&i.CreateAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChildRolesByParentId = `-- name: GetChildRolesByParentId :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number, created_by, create_at, update_at
FROM ` + "`" + `role` + "`" + `
WHERE created_by = ? AND is_deleted = false
ORDER BY role_left_value ASC
`

type GetChildRolesByParentIdRow struct {
	ID             string
	Code           string
	RoleName       string
	RoleLeftValue  int32
	RoleRightValue int32
	RoleMaxNumber  int32
	CreatedBy      string
	CreateAt       time.Time
	UpdateAt       time.Time
}

func (q *Queries) GetChildRolesByParentId(ctx context.Context, createdBy string) ([]GetChildRolesByParentIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getChildRolesByParentId, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetChildRolesByParentIdRow
	for rows.Next() {
		var i GetChildRolesByParentIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.RoleName,
			&i.RoleLeftValue,
			&i.RoleRightValue,
			&i.RoleMaxNumber,
			&i.CreatedBy,
			&i.CreateAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMaxRightValue = `-- name: GetMaxRightValue :one
SELECT COALESCE(MAX(role_right_value), 0) as max_right_value
FROM ` + "`" + `role` + "`" + `
WHERE is_deleted = false
`

func (q *Queries) GetMaxRightValue(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getMaxRightValue)
	var max_right_value interface{}
	err := row.Scan(&max_right_value)
	return max_right_value, err
}

const getParentRoleInfo = `-- name: GetParentRoleInfo :one
SELECT role_left_value, role_right_value
FROM ` + "`" + `role` + "`" + `
WHERE id = ? AND is_deleted = false
`

type GetParentRoleInfoRow struct {
	RoleLeftValue  int32
	RoleRightValue int32
}

func (q *Queries) GetParentRoleInfo(ctx context.Context, id string) (GetParentRoleInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getParentRoleInfo, id)
	var i GetParentRoleInfoRow
	err := row.Scan(&i.RoleLeftValue, &i.RoleRightValue)
	return i, err
}

const getRoleById = `-- name: GetRoleById :one
SELECT id, code, role_name,role_left_value,role_right_value,role_max_number,created_by,create_at,update_at
FROM ` + "`" + `role` + "`" + `
WHERE id = ? AND is_deleted = false
`

type GetRoleByIdRow struct {
	ID             string
	Code           string
	RoleName       string
	RoleLeftValue  int32
	RoleRightValue int32
	RoleMaxNumber  int32
	CreatedBy      string
	CreateAt       time.Time
	UpdateAt       time.Time
}

func (q *Queries) GetRoleById(ctx context.Context, id string) (GetRoleByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getRoleById, id)
	var i GetRoleByIdRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.RoleName,
		&i.RoleLeftValue,
		&i.RoleRightValue,
		&i.RoleMaxNumber,
		&i.CreatedBy,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getRoleWithChildren = `-- name: GetRoleWithChildren :many
SELECT r.id, r.code, r.role_name, r.role_left_value, r.role_right_value, r.role_max_number, r.created_by, r.create_at, r.update_at
FROM ` + "`" + `role` + "`" + ` r
WHERE r.role_left_value >= (SELECT r2.role_left_value FROM ` + "`" + `role` + "`" + ` r2 WHERE r2.id = ? AND r2.is_deleted = false)
AND r.role_right_value <= (SELECT r3.role_right_value FROM ` + "`" + `role` + "`" + ` r3 WHERE r3.id = ? AND r3.is_deleted = false)
AND r.is_deleted = false
ORDER BY r.role_left_value ASC
`

type GetRoleWithChildrenParams struct {
	ID   string
	ID_2 string
}

type GetRoleWithChildrenRow struct {
	ID             string
	Code           string
	RoleName       string
	RoleLeftValue  int32
	RoleRightValue int32
	RoleMaxNumber  int32
	CreatedBy      string
	CreateAt       time.Time
	UpdateAt       time.Time
}

func (q *Queries) GetRoleWithChildren(ctx context.Context, arg GetRoleWithChildrenParams) ([]GetRoleWithChildrenRow, error) {
	rows, err := q.db.QueryContext(ctx, getRoleWithChildren, arg.ID, arg.ID_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRoleWithChildrenRow
	for rows.Next() {
		var i GetRoleWithChildrenRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.RoleName,
			&i.RoleLeftValue,
			&i.RoleRightValue,
			&i.RoleMaxNumber,
			&i.CreatedBy,
			&i.CreateAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRolesWithPagination = `-- name: GetRolesWithPagination :many
SELECT id, code, role_name, role_left_value, role_right_value, role_max_number, created_by, create_at, update_at
FROM ` + "`" + `role` + "`" + `
WHERE is_deleted = false
ORDER BY role_left_value ASC LIMIT ? OFFSET ?
`

type GetRolesWithPaginationParams struct {
	Limit  int32
	Offset int32
}

type GetRolesWithPaginationRow struct {
	ID             string
	Code           string
	RoleName       string
	RoleLeftValue  int32
	RoleRightValue int32
	RoleMaxNumber  int32
	CreatedBy      string
	CreateAt       time.Time
	UpdateAt       time.Time
}

func (q *Queries) GetRolesWithPagination(ctx context.Context, arg GetRolesWithPaginationParams) ([]GetRolesWithPaginationRow, error) {
	rows, err := q.db.QueryContext(ctx, getRolesWithPagination, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRolesWithPaginationRow
	for rows.Next() {
		var i GetRolesWithPaginationRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.RoleName,
			&i.RoleLeftValue,
			&i.RoleRightValue,
			&i.RoleMaxNumber,
			&i.CreatedBy,
			&i.CreateAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalAccounts = `-- name: GetTotalAccounts :one
SELECT COALESCE(CAST(SUM(role_max_number) AS SIGNED), 0) AS MaxNumberAccount
FROM ` + "`" + `role` + "`" + `
WHERE (created_by = ? OR ? IS NULL OR ? = '')
    AND is_deleted = false
`

type GetTotalAccountsParams struct {
	CreatedBy string
	Column2   interface{}
	Column3   interface{}
}

func (q *Queries) GetTotalAccounts(ctx context.Context, arg GetTotalAccountsParams) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getTotalAccounts, arg.CreatedBy, arg.Column2, arg.Column3)
	var maxnumberaccount interface{}
	err := row.Scan(&maxnumberaccount)
	return maxnumberaccount, err
}

const getTotalRoles = `-- name: GetTotalRoles :one
SELECT COUNT(*) FROM ` + "`" + `role` + "`" + ` WHERE is_deleted = false
`

func (q *Queries) GetTotalRoles(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalRoles)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const softDeleteRolesByRange = `-- name: SoftDeleteRolesByRange :exec
UPDATE ` + "`" + `role` + "`" + `
SET is_deleted = true, update_at = ?
WHERE role_left_value >= ? AND role_right_value <= ? AND is_deleted = false
`

type SoftDeleteRolesByRangeParams struct {
	UpdateAt       time.Time
	RoleLeftValue  int32
	RoleRightValue int32
}

func (q *Queries) SoftDeleteRolesByRange(ctx context.Context, arg SoftDeleteRolesByRangeParams) error {
	_, err := q.db.ExecContext(ctx, softDeleteRolesByRange, arg.UpdateAt, arg.RoleLeftValue, arg.RoleRightValue)
	return err
}

const updateLeftValuesForInsert = `-- name: UpdateLeftValuesForInsert :exec
UPDATE ` + "`" + `role` + "`" + `
SET role_left_value = role_left_value + 2
WHERE role_left_value > ? AND is_deleted = false
`

func (q *Queries) UpdateLeftValuesForInsert(ctx context.Context, roleLeftValue int32) error {
	_, err := q.db.ExecContext(ctx, updateLeftValuesForInsert, roleLeftValue)
	return err
}

const updateLicenseByRoleId = `-- name: UpdateLicenseByRoleId :exec
UPDATE ` + "`" + `role` + "`" + `
SET license_id = ?
WHERE id = ? AND is_deleted = false
`

type UpdateLicenseByRoleIdParams struct {
	LicenseID string
	ID        string
}

func (q *Queries) UpdateLicenseByRoleId(ctx context.Context, arg UpdateLicenseByRoleIdParams) error {
	_, err := q.db.ExecContext(ctx, updateLicenseByRoleId, arg.LicenseID, arg.ID)
	return err
}

const updateRightValuesForInsert = `-- name: UpdateRightValuesForInsert :exec
UPDATE ` + "`" + `role` + "`" + `
SET role_right_value = role_right_value + 2
WHERE role_right_value >= ? AND is_deleted = false
`

func (q *Queries) UpdateRightValuesForInsert(ctx context.Context, roleRightValue int32) error {
	_, err := q.db.ExecContext(ctx, updateRightValuesForInsert, roleRightValue)
	return err
}

const updateRole = `-- name: UpdateRole :exec
UPDATE ` + "`" + `role` + "`" + `
SET code = ?, role_name = ?,role_left_value = ?,role_right_value = ?,role_max_number = ?,created_by = ?,update_at = NOW()
WHERE id = ?
`

type UpdateRoleParams struct {
	Code           string
	RoleName       string
	RoleLeftValue  int32
	RoleRightValue int32
	RoleMaxNumber  int32
	CreatedBy      string
	ID             string
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateRole,
		arg.Code,
		arg.RoleName,
		arg.RoleLeftValue,
		arg.RoleRightValue,
		arg.RoleMaxNumber,
		arg.CreatedBy,
		arg.ID,
	)
	return err
}

const updateRoleTree = `-- name: UpdateRoleTree :execresult
UPDATE ` + "`" + `role` + "`" + ` SET role_right_value = role_right_value + 2 WHERE role_right_value >= ? AND is_deleted = false
`

func (q *Queries) UpdateRoleTree(ctx context.Context, roleRightValue int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateRoleTree, roleRightValue)
}
