package rbac

import (
	"context"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
)

// RBAC lấy tất cả các vai trò và nhóm quyền theo tài khoản
func GetFullPermisionByAccount(ctx context.Context, accountId string) (out []model.RolePermission, err error) {
	// Lấy danh sách vai trò và quyền theo tài khoản
	roles, err := service.RoleItem().GetAllPermissionByAccountId(ctx, accountId)
	if err != nil {
		return roles, nil
	}
	return roles, nil
}
