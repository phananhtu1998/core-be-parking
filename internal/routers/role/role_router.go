package role

import (
	"go-backend-api/internal/controller/role"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (ar *RoleRouter) InitLoginRouter(Router *gin.RouterGroup) {
	roleRouterPrivate := Router.Group("/role")
	//roleRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		roleRouterPrivate.POST("/create_role", role.Roles.CreateRole)
	}
}
