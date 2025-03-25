package roleaccount

import (
	roleaccount "go-backend-api/internal/controller/role_account"

	"github.com/gin-gonic/gin"
)

type RoleAccountRouterGroup struct {
	RoleAccountRouter
}

func (ar *RoleAccountRouter) InitRoleAccountRouter(Router *gin.RouterGroup) {
	roleAccountRouterPrivate := Router.Group("/roleaccount")
	//roleAccountRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		roleAccountRouterPrivate.POST("/create_roles_account", roleaccount.RoleAccounts.CreateRoleAccount)

	}
}
