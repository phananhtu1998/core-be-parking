package manage

import (
	"go-backend-api/internal/controller/account"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	adminRouterPublic.Use(middlewares.AuthenMiddleware())
	{
		adminRouterPublic.GET("/get_all_account", account.Accounts.GetAllAccount)
		adminRouterPublic.GET("/get_account_by_id/:id", account.Accounts.GetAccountById)
		adminRouterPublic.PUT("/update_account/:id", account.Accounts.UpdateAccount)
		adminRouterPublic.DELETE("/delete_account/:id", account.Accounts.DeleteAccount)
		adminRouterPublic.POST("/create_account/", account.Accounts.CreateAccount)

	}
	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
