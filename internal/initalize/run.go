package initalize

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/QuocHuannn/Go-to-Work/global"
)

func Run() {
	//Load configuration
	LoadConfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("InitLogger success", zap.String("oke", "success"))
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()
	if err := r.Run(":8002"); err != nil {
		global.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
