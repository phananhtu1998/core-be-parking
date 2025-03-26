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
		roleAccountRouterPrivate.GET("/get_role_account_by_role_id/:id", roleaccount.RoleAccounts.GetAllRoleAccountByRoleId)
		roleAccountRouterPrivate.GET("/get_role_account_by_account_id/:id", roleaccount.RoleAccounts.GetAllRoleAccountByAccountId)
		roleAccountRouterPrivate.DELETE("/delete_multiple_role_account", roleaccount.RoleAccounts.DeleteRoleAccount)
	}
}
