package initalize

import (
	"strconv"

	"go.uber.org/zap"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/config"
)

func Run() {
	// Configuration is already loaded in main.go,
	// we can now use config.Cfg for all settings

	// Initialize components
	InitLogger()
	global.Logger.Info("InitLogger success", zap.String("oke", "success"))
	InitMysql()
	InitRedis()

	// Initialize router and start HTTP server
	r := InitRouter()

	// Use server port from configuration
	port := ":" + strconv.Itoa(config.Cfg.Server.Port)
	global.Logger.Info("Starting server", zap.Int("port", config.Cfg.Server.Port))

	if err := r.Run(port); err != nil {
		global.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
