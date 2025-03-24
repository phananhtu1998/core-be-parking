package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IRolesMenu interface {
		CreateRolesMenu(ctx context.Context, in *model.RolesMenu) (codeResult int, out model.RolesMenu, err error)
		GetRoleMenuByRoleId(ctx context.Context, roleId, search string) (int, []model.RoleMenuOutput, error)
	}
)

var (
	localRolesMenu IRolesMenu
)

func RolesMenuItem() IRolesMenu {
	if localRolesMenu == nil {
		panic("implement localRolesMenu not found for interface IRolesMenu")
	}
	return localRolesMenu
}

func InitRolesMenuItem(i IRolesMenu) {
	localRolesMenu = i
}
