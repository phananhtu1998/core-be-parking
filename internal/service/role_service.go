package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IRole interface {
		CreateRole(ctx context.Context, in *model.Role) (codeResult int, out model.Role, err error)
		GetAllRoles(ctx context.Context, page, pageSize int) (codeResult int, out []model.RoleHierarchyOutput, total int64, err error)
		GetAllPermission(ctx context.Context) (codeResult int, out []model.RolePermission, err error)
		GetAllPermissionByAccountId(ctx context.Context, Account_Id string) (out []model.RolePermission, err error)
		GetRoleById(ctx context.Context, parentId string) (codeResult int, out []model.RoleHierarchyOutput, err error)
		DeleteRole(ctx context.Context, id string) (codeResult int, err error)
	}
)

var (
	localRole IRole
)

func RoleItem() IRole {
	if localRole == nil {
		panic("implement localRole not found for interface IRole")
	}
	return localRole
}

func InitRoleItem(i IRole) {
	localRole = i
}
