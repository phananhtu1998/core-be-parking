package initialize

import (
	"go-backend-api/global"
	consts "go-backend-api/internal/const"
	"go-backend-api/internal/middlewares"
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
	r.Use(middlewares.RateLimiterGlobalMiddlewareRedis())
	manageRouter := routers.RouterGroupApp.Manage
	loginRouter := routers.RouterGroupApp.Login
	menuRouter := routers.RouterGroupApp.Menu
	roleRouter := routers.RouterGroupApp.Role
	rolesMenuRouter := routers.RouterGroupApp.RolesMenu
	roleAccountRouter := routers.RouterGroupApp.RoleAccount
	licenseRouter := routers.RouterGroupApp.License
	MainGroup := r.Group(consts.HOST_PREFIX)
	{
		MainGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "pong",
			})
		})
		MainGroup.GET("/checkstatus") //tracking monitor
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		loginRouter.InitLoginRouter(MainGroup)
		menuRouter.InitAdminRouter(MainGroup)
		roleRouter.InitRoleRouter(MainGroup)
		rolesMenuRouter.InitRolesMenuRouter(MainGroup)
		roleAccountRouter.InitRoleAccountRouter(MainGroup)
		licenseRouter.InitLicenseRouter(MainGroup)
	}
	return r
}
