package license

import (
	"go-backend-api/internal/controller/license"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type LicenseRouter struct{}

func (ar *LicenseRouter) InitLicenseRouter(Router *gin.RouterGroup) {
	licenseRouterPrivate := Router.Group("/license")
	licenseRouterPrivate.Use(middlewares.AuthenMiddleware())
	licenseRouterPrivate.Use(middlewares.LicenseMiddleware())
	licenseRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
	{
		licenseRouterPrivate.POST("/create_license", license.Licenses.CreateAccount)
	}
}
