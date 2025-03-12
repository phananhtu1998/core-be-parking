package initialize

import (
	"go-backend-api/global"
	"go-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
