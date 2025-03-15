package initalize

import (
	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/QuocHuannn/Go-to-Work/internal/config"
	"github.com/QuocHuannn/Go-to-Work/pkg/logger"
	"github.com/QuocHuannn/Go-to-Work/pkg/setting"
)

func InitLogger() {
	// Convert from new config structure to old logger settings format
	loggerSettings := setting.LoggerSetting{
		Log_level:     config.Cfg.App.LogLevel,
		File_log_name: "logs/app.log",
		Max_size:      10,
		Max_backups:   5,
		Max_age:       30,
		Compress:      true,
	}

	global.Logger = logger.NewLogger(loggerSettings)
}
