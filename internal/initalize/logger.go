package initalize

import "github.com/QuocHuannn/Go-to-Work/global"
import "github.com/QuocHuannn/Go-to-Work/pkg/logger"

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)

}
