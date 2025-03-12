package manage

import (
	"go-backend-api/internal/controller/account"
	"log"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		log.Println("kkkkkkkkkkkkkkkk")
		adminRouterPublic.GET("/get_all_account", account.Accounts.GetAllAccount)
		adminRouterPublic.GET("/get_account_by_id/:id", account.Accounts.GetAccountById)

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
