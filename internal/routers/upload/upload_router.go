package upload

import (
	"go-backend-api/internal/controller/upload"

	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (ar *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouterPrivate := Router.Group("/upload")
	//uploadRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		uploadRouterPrivate.POST("/upload_file", upload.UploadFileHandler)
	}
}
