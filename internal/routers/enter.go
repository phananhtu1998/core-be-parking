package routers

import (
	"go-backend-api/internal/routers/funcpackage"
	"go-backend-api/internal/routers/license"
	"go-backend-api/internal/routers/login"
	"go-backend-api/internal/routers/manage"
	"go-backend-api/internal/routers/menu"
	rolesmenu "go-backend-api/internal/routers/roles_menu"
	"go-backend-api/internal/routers/upload"
)

type RouterGroup struct {
	Manage manage.ManageRouterGoup
	Login  login.LoginRouterGroup
	Menu   menu.MenuRouterGoup
	// Role        role.RoleRouterGroup
	RolesMenu rolesmenu.RolesMenuRouterGroup
	//RoleAccount roleaccount.RoleAccountsRouterGroup
	License     license.LicenseRouterGroup
	Upload      upload.UploadRouterGroup
	Funcpackage funcpackage.FuncpackageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
