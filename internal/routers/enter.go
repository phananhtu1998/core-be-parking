package routers

import (
	"go-backend-api/internal/routers/login"
	"go-backend-api/internal/routers/manage"
	"go-backend-api/internal/routers/menu"
)

type RouterGroup struct {
	Manage manage.ManageRouterGoup
	Login  login.LogimRouterGroup
	Menu   menu.MenuRouterGoup
}

var RouterGroupApp = new(RouterGroup)
