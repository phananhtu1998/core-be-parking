package initialize

import (
	"go-backend-api/global"
	"go-backend-api/internal/database"
	"go-backend-api/internal/service"
	"go-backend-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// user service interface
	service.InitAccountItem(impl.NewAccountImpl(queries))
}
