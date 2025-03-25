package routers

import (
	"go-backend-api/internal/routers/login"
	"go-backend-api/internal/routers/manage"
	"go-backend-api/internal/routers/menu"
	"go-backend-api/internal/routers/role"
	rolesmenu "go-backend-api/internal/routers/roles_menu"
)

type RouterGroup struct {
	Manage    manage.ManageRouterGoup
	Login     login.LoginRouterGroup
	Menu      menu.MenuRouterGoup
	Role      role.RoleRouterGroup
	RolesMenu rolesmenu.RolesMenuRouterGroup
}

var RouterGroupApp = new(RouterGroup)
