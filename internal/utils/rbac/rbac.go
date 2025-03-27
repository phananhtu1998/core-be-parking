package rbac

import (
	"context"
	"go-backend-api/internal/model"
	"go-backend-api/internal/service"
	"go-backend-api/internal/utils/cache"
	"go-backend-api/pkg/response"
)

// RBAC lấy tất cả các vai trò và nhóm quyền
func GetFullPermision(ctx context.Context) (codeResult int, out []model.RolePermission, err error) {
	// Lấy danh sách vai trò và quyền
	codeRole, roles, err := service.RoleItem().GetAllPermission(ctx)
	if err != nil {
		return response.ErrCodeRoleError, roles, err
	}
	return codeRole, out, err
}

// RBAC lấy tất cả các vai trò và nhóm quyền theo tài khoản
func GetFullPermisionByAccount(ctx context.Context, subtoken string) (out []model.RolePermission, err error) {
	var infoUser model.GetCacheTokenForChangePassword
	if err := cache.GetCache(ctx, subtoken, &infoUser); err != nil {
		return out, nil
	}
	// Lấy danh sách vai trò và quyền theo tài khoản
	roles, err := service.RoleItem().GetAllPermissionByAccountId(ctx, infoUser.ID)
	if err != nil {
		return roles, nil
	}
	return out, nil
}
