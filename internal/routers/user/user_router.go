package user

import (
	"go-backend-api/internal/controller/user"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouterGroup struct {
	UserRouter
}

func (ar *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	userRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
	userRouterPrivate.Use(middlewares.LicenseMiddleware())
	{
		userRouterPrivate.POST("/create_user", user.Users.CreateUser)
	}
}
