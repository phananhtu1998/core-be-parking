package initialize

import (
	"fmt"
	"go-backend-api/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	// Load config
	Loadconfig()
	fmt.Println("username: ", global.Config.Mysql.Username)
	// Write logs
	InitLogger()
	global.Logger.Info("config Log ok!", zap.String("ok", "success"))
	// connect mysql and redis
	InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	r := InitRouter()
	if r == nil {
		global.Logger.Error("Failed to initialize Router")
		return nil
	}
	return r
}
