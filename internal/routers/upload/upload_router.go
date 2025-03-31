package upload

import (
	"go-backend-api/internal/controller/upload"
	"go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (ar *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouterPrivate := Router.Group("/upload")
	uploadRouterPrivate.Use(middlewares.AuthenMiddleware())
	uploadRouterPrivate.Use(middlewares.RateLimiterPrivateMiddlewareRedis())
	uploadRouterPrivate.Use(middlewares.LicenseMiddleware())
	{
		uploadRouterPrivate.POST("/upload_file", upload.UploadFileHandler)
	}
}
