package initialize

import (
	"fmt"
	"go-backend-api/global"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false, // tắt transaction tăng hiệu năng cho mysql nhưng lại tính nhất quán dữ liệu không cao
	})
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db
	SetPool()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error: %s:: ", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)                      // số lượng kết nối đến server nếu vượt quá sẽ đợi để kết nối
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) //
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(global.Mdb) // reuse your gorm db
	g.Execute()
}

// func migrateTables() {
// 	err := global.Mdb.AutoMigrate(
// 		&po.User{},
// 		&po.Role{},
// 	)
// 	if err != nil {
// 		fmt.Println("Migrating tables error: ", err)
// 	}
// }
