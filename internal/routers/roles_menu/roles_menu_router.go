package rolesmenu

import (
	"go-backend-api/internal/controller/roles_menu"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type RolesMenuRouter struct{}

func (ar *RolesMenuRouter) InitRolesMenuRouter(Router *gin.RouterGroup) {
	rolesMenuRouterPrivate := Router.Group("/funcpackagesmenu")
	rolesMenuRouterPrivate.Use(middlewares.AuthenMiddleware())
	rolesMenuRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
	rolesMenuRouterPrivate.Use(middlewares.LicenseMiddleware())
	{
		rolesMenuRouterPrivate.POST("/create_funcpackage_menu", roles_menu.RolesMenus.CreateRolesMenu)
		rolesMenuRouterPrivate.POST("/create_funcpackage_menu_multiple", roles_menu.RolesMenus.CreateMultipleRoleMenus)
		rolesMenuRouterPrivate.GET("/get_funcpackage_menu_by_funcpackage_id/:id", roles_menu.RolesMenus.GetRoleMenuByRoleId)
		rolesMenuRouterPrivate.PUT("/update_funcpackage_menu/:id", roles_menu.RolesMenus.UpdateRolesMenu)
		rolesMenuRouterPrivate.DELETE("/delete_funcpackage_menu/:id", roles_menu.RolesMenus.DeleteRolesMenu)
	}
}
