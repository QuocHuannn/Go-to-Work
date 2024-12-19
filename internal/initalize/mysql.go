package initalize

import (
	"fmt"
	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Failed to connect database")
	global.Logger.Info("Connect database success", zap.String("host", m.Host))
	global.Mdb = db

	// Set pool
	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	checkErrorPanic(err, "Failed to set pool")

	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnectTimeout))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	}
}
