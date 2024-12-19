package global

import (
	"github.com/QuocHuannn/Go-to-Work/pkg/logger"
	"github.com/QuocHuannn/Go-to-Work/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)

/*
Config
Redis
Mysql
...
*/
