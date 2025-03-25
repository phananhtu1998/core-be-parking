package rolesmenu

import (
	"go-backend-api/internal/controller/rolesmenu"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type RolesMenuRouter struct{}

func (ar *RolesMenuRouter) InitRolesMenuRouter(Router *gin.RouterGroup) {
	rolesMenuRouterPrivate := Router.Group("/rolesmenu")
	rolesMenuRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		rolesMenuRouterPrivate.POST("/create_roles_menu", rolesmenu.RolesMenus.CreateRolesMenu)
		rolesMenuRouterPrivate.GET("/get_role_menu_by_role_id/:id", rolesmenu.RolesMenus.GetRoleMenuByRoleId)
		rolesMenuRouterPrivate.PUT("/update_roles_menu/:id", rolesmenu.RolesMenus.UpdateRolesMenu)
		rolesMenuRouterPrivate.DELETE("/delete_roles_menu/:id", rolesmenu.RolesMenus.DeleteRolesMenu)
	}
}
