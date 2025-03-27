package login

import (
	"go-backend-api/global"
	"go-backend-api/internal/controller/login"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct{}

func (ar *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/auth")
	adminRouterPublic.Use(middlewares.PermissionMiddleware(global.Enforcer))
	{
		adminRouterPublic.POST("/login", login.Logins.Login)
	}
	adminRouterPrivate := Router.Group("/auth")
	adminRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		adminRouterPrivate.POST("/logout", login.Logins.Logout)
		adminRouterPrivate.POST("/change_password", login.Logins.ChangePassword)
	}
	adminRouterRefreshToken := Router.Group("/auth")
	adminRouterRefreshToken.Use(middlewares.AuthenMiddlewareV2())
	{
		adminRouterRefreshToken.POST("/refreshtoken", login.Logins.RefreshTokens)
	}
}
