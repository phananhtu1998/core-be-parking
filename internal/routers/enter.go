package routers

import (
	"go-backend-api/internal/routers/login"
	"go-backend-api/internal/routers/manage"
	"go-backend-api/internal/routers/menu"
	"go-backend-api/internal/routers/role"
)

type RouterGroup struct {
	Manage manage.ManageRouterGoup
	Login  login.LogimRouterGroup
	Menu   menu.MenuRouterGoup
	Role   role.RoleRouterGroup
}

var RouterGroupApp = new(RouterGroup)
