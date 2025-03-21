package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IRole interface {
		CreateRole(ctx context.Context, in *model.RoleInput) (codeResult int, out model.RoleOutput, err error)
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
