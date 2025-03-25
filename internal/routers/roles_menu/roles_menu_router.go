package rolesmenu

import (
	"go-backend-api/internal/controller/roles_menu"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type RolesMenuRouter struct{}

func (ar *RolesMenuRouter) InitRolesMenuRouter(Router *gin.RouterGroup) {
	rolesMenuRouterPrivate := Router.Group("/rolesmenu")
	rolesMenuRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		rolesMenuRouterPrivate.POST("/create_roles_menu", roles_menu.RolesMenus.CreateRolesMenu)
		rolesMenuRouterPrivate.GET("/get_role_menu_by_role_id/:id", roles_menu.RolesMenus.GetRoleMenuByRoleId)
		rolesMenuRouterPrivate.PUT("/update_roles_menu/:id", roles_menu.RolesMenus.UpdateRolesMenu)
		rolesMenuRouterPrivate.DELETE("/delete_roles_menu/:id", roles_menu.RolesMenus.DeleteRolesMenu)
	}
}
