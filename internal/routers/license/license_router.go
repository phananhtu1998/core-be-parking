package license

import (
	"go-backend-api/internal/controller/license"

	"github.com/gin-gonic/gin"
)

type LicenseRouter struct{}

func (ar *LicenseRouter) InitLicenseRouter(Router *gin.RouterGroup) {
	licenseRouterPrivate := Router.Group("/license")
	//licenseRouterPrivate.Use(middlewares.AuthenMiddleware())
	//licenseRouterPrivate.Use(middlewares.PermissionMiddleware(global.Enforcer))
	{
		licenseRouterPrivate.POST("/create_license", license.Licenses.CreateAccount)
	}
}
