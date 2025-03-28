package routers

import (
	"go-backend-api/internal/routers/license"
	"go-backend-api/internal/routers/login"
	"go-backend-api/internal/routers/manage"
	"go-backend-api/internal/routers/menu"
	"go-backend-api/internal/routers/role"
	roleaccount "go-backend-api/internal/routers/role_account"
	rolesmenu "go-backend-api/internal/routers/roles_menu"
)

type RouterGroup struct {
	Manage      manage.ManageRouterGoup
	Login       login.LoginRouterGroup
	Menu        menu.MenuRouterGoup
	Role        role.RoleRouterGroup
	RolesMenu   rolesmenu.RolesMenuRouterGroup
	RoleAccount roleaccount.RoleAccountRouterGroup
	License     license.LicenseRouterGroup
}

var RouterGroupApp = new(RouterGroup)
