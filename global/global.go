package global

import (
	"github.com/QuocHuannn/Go-to-Work/pkg/logger"
	"github.com/QuocHuannn/Go-to-Work/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	// Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
...
*/
