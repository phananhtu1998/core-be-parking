package rolesmenu

import (
	"go-backend-api/internal/controller/rolesmenu"

	"github.com/gin-gonic/gin"
)

type RolesMenuRouter struct{}

func (ar *RolesMenuRouter) InitRolesMenuRouter(Router *gin.RouterGroup) {
	rolesMenuRouterPrivate := Router.Group("/rolesmenu")
	//roleRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		rolesMenuRouterPrivate.POST("/create_roles_menu", rolesmenu.RolesMenus.CreateRolesMenu)
	}
}
