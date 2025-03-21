package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IRole interface {
		CreateRole(ctx context.Context, in *model.Role) (codeResult int, out model.Role, err error)
		GetAllRoles(ctx context.Context) (codeResult int, out []model.RoleHierarchyOutput, err error)
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
