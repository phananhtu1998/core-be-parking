package initialize

import (
	"fmt"
	"go-backend-api/global"
	"go-backend-api/internal/database"

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
	InitMinio()
	GetServerInfo()
	queries := database.New(global.Mdbc)
	InsertData(ctx, queries)
	// Khởi tạo RBAC
	enforcer, err := InitializeRBAC(global.Mdb)
	if err != nil {
		global.Logger.Error("Failed to initialize RBAC", zap.Error(err))
		return nil
	}
	global.Enforcer = enforcer

	r := InitRouter()
	if r == nil {
		global.Logger.Error("Failed to initialize Router")
		return nil
	}
	return r
}
