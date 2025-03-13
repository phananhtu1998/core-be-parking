package login

import (
	"go-backend-api/internal/controller/login"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct{}

func (ar *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/login")
	{
		adminRouterPublic.POST("/login", login.Logins.Login)
	}
}
