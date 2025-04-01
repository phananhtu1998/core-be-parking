package funcpackage

import (
	"go-backend-api/internal/controller/funcpackage"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type FuncpackageRouter struct{}

func (ar *FuncpackageRouter) InitFuncpackageRouter(Router *gin.RouterGroup) {
	funcpackageRouterPrivate := Router.Group("/funcpackage")
	funcpackageRouterPrivate.Use(middlewares.AuthenMiddleware())
	funcpackageRouterPrivate.Use(middlewares.LicenseMiddleware())
	funcpackageRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
	{
		funcpackageRouterPrivate.POST("/create_func_package", funcpackage.Funcpackages.CreateFuncPackage)
	}
}
