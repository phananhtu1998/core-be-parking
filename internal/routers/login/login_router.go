package login

import (
	"go-backend-api/internal/controller/login"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct{}

func (ar *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/auth")
	{
		adminRouterPublic.POST("/login", login.Logins.Login)
	}
	adminRouterPrivate := Router.Group("/auth")
	adminRouterPrivate.Use(middlewares.AuthenMiddleware())
	adminRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
	//adminRouterPublic.Use(middlewares.LicenseMiddleware())
	//adminRouterPrivate.Use(middlewares.PermissionMiddleware(global.Enforcer))
	{
		adminRouterPrivate.POST("/logout", login.Logins.Logout)
		adminRouterPrivate.POST("/change_password", login.Logins.ChangePassword)
	}
	adminRouterRefreshToken := Router.Group("/auth")
	adminRouterRefreshToken.Use(middlewares.AuthenMiddlewareV2())
	adminRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
	adminRouterPublic.Use(middlewares.LicenseMiddleware())
	{
		adminRouterRefreshToken.POST("/refresh-token", login.Logins.RefreshTokens)
	}
}
