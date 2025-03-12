package routers

import (
	"go-backend-api/internal/routers/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGoup
}

var RouterGroupApp = new(RouterGroup)
