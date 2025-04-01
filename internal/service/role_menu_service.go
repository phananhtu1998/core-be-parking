package service

import (
	"context"
	"go-backend-api/internal/model"
)

type (
	IRolesMenu interface {
		CreateRolesMenu(ctx context.Context, in *model.RolesMenu) (codeResult int, out model.RolesMenu, err error)
		CreateMultipleRoleMenus(ctx context.Context, inputs []model.RolesMenu) (int, []model.RolesMenu, error)
		GetRoleMenuByRoleId(ctx context.Context, roleId, search string) (int, []model.RoleMenuOutput, error)
		UpdateRolesMenu(ctx context.Context, id string, in *model.RolesMenu) (int, model.RolesMenu, error)
		DeleteRolesMenu(ctx context.Context, id string) (int, error)
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
