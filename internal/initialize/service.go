package initialize

import (
	"go-backend-api/global"
	"go-backend-api/internal/database"
	"go-backend-api/internal/service"
	"go-backend-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	tx, err := global.Mdbc.Begin()
	if err != nil {
		// handle error
		return
	}
	// user service interface
	service.InitAccountItem(impl.NewAccountImpl(queries, tx, global.Mdbc))
	service.InitLoginItem(impl.NewLoginImpl(queries))
	service.InitMenuItem(impl.NewMenuImpl(queries, tx, global.Mdbc))
	service.InitRoleItem(impl.NewRoleImpl(queries, tx, global.Mdbc))
	service.InitRolesMenuItem(impl.NewRolesMenuImpl(queries, tx, global.Mdbc))
	service.InitRoleAccountItem(impl.NewRoleAccountImpl(queries, tx, global.Mdbc))
	service.InitLicenseItem(impl.NewLicenseImpl(queries, tx, global.Mdbc))
	service.InitUploadItem(impl.NewUploadImpl())
	service.InitFuncpackageItem(impl.NewFuncpackageImpl(queries, tx, global.Mdbc))
}
