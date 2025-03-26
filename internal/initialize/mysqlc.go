package initialize

import (
	"database/sql"
	"fmt"
	"go-backend-api/global"

	// "go-backend-api/internal/po"
	"time"
)

//	func checkErrorPanicC(err error, errString string) {
//		if err != nil {
//			global.Logger.Error(errString, zap.Error(err))
//			panic(err)
//		}
//	}
func InitMysqlC() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	fmt.Println("dsn: ", dsn)
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdbc = db
	SetPool()
	genTableDAO()
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error: %s:: ", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)                      // số lượng kết nối đến server nếu vượt quá sẽ đợi để kết nối
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) //
}

// func genTableDAOC() {
// 	g := gen.NewGenerator(gen.Config{
// 		OutPath: "./internal/models",
// 		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
// 	})
// 	g.UseDB(global.Mdb) // reuse your gorm db
// 	g.GenerateModel("go_crm_user")
// 	g.Execute()
// }

// func migrateTablesC() {
// 	err := global.Mdb.AutoMigrate(
// 		&po.User{},
// 		&po.Role{},
// 	)
// 	if err != nil {
// 		fmt.Println("Migrating tables error: ", err)
// 	}
// }
