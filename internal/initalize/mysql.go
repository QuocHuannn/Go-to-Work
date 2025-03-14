package initalize

import (
	"fmt"
	"os"
	"time"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/config"
	"github.com/QuocHuannn/Go-to-Work/internal/po"
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

// For non-critical errors that shouldn't cause a panic
func logError(err error, errString string) bool {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		return true
	}
	return false
}

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	return global.Mdb
}

func InitMysql() {
	m := config.Cfg.DB

	// Override host with environment variable if running in Docker
	host := m.Host
	if os.Getenv("MYSQL_HOST") != "" {
		host = os.Getenv("MYSQL_HOST")
		fmt.Printf("Using MySQL host from environment: %s\n", host)
	} else {
		fmt.Printf("Using MySQL host from config: %s\n", host)
	}

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.Username, m.Password, host, m.Port, m.DBName)

	fmt.Printf("MySQL DSN (without password): %s:%s@tcp(%s:%v)/%s\n",
		m.Username, "****", host, m.Port, m.DBName)

	// Thêm cơ chế retry cho kết nối MySQL trong Docker
	var db *gorm.DB
	var err error

	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(s), &gorm.Config{
			SkipDefaultTransaction: false,
		})

		if err == nil {
			break
		}

		fmt.Printf("Failed to connect to MySQL (attempt %d/%d): %v\n", i+1, maxRetries, err)
		if i < maxRetries-1 {
			fmt.Printf("Retrying in %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	checkErrorPanic(err, "Failed to connect database after multiple attempts")
	global.Logger.Info("Connect database success", zap.String("host", host))
	global.Mdb = db

	// Set pool
	SetPool()

	// Create tables first
	migrateTables()

	// Then try to generate DAO, but don't panic if it fails
	safeGenTableDAO()
}

func SetPool() {
	m := config.Cfg.DB
	sqlDB, err := global.Mdb.DB()
	checkErrorPanic(err, "Failed to set pool")

	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// Safe version that doesn't panic
func safeGenTableDAO() {
	defer func() {
		if r := recover(); r != nil {
			global.Logger.Warn("Recovered from panic in genTableDAO", zap.Any("recover", r))
		}
	}()

	// Check if table exists before generating
	var count int64
	err := global.Mdb.Table("information_schema.tables").
		Where("table_schema = ? AND table_name = ?", config.Cfg.DB.DBName, "go_crm_user").
		Count(&count).Error

	if err != nil {
		global.Logger.Warn("Failed to check if table exists", zap.Error(err))
		return
	}

	if count == 0 {
		global.Logger.Warn("Table go_crm_user doesn't exist, skipping DAO generation")
		return
	}

	// Only proceed if the table exists
	genTableDAO()
}

func genTableDAO() {
	// init table
	g := gen.NewGenerator(gen.Config{
		OutPath: "/internal/models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(global.Mdb) // reuse your gorm db

	// Only generate if table exists
	g.GenerateModel("go_crm_user")
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
	} else {
		global.Logger.Info("Tables migrated successfully")
	}
}
