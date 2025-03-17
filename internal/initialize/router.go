package initialize

import (
	"go-backend-api/global"
	"go-backend-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middleware
	r.Use() //logging
	r.Use() //cross
	r.Use() //limiter global
	manageRouter := routers.RouterGroupApp.Manage
	loginRouter := routers.RouterGroupApp.Login
	menuRouter := routers.RouterGroupApp
	MainGroup := r.Group("/v1/2025")
	{
		MainGroup.GET("/checkstatus") //tracking monitor
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		loginRouter.InitLoginRouter(MainGroup)
		menuRouter.Menu.InitAdminRouter(MainGroup)
	}
	return r
}
