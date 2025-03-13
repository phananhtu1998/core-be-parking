package routers

import (
	"go-backend-api/internal/routers/login"
	"go-backend-api/internal/routers/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGoup
	Login  login.LogimRouterGroup
}

var RouterGroupApp = new(RouterGroup)
