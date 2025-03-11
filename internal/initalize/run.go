package initalize

import (
	"fmt"
	"strconv"

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
	port := ":" + strconv.Itoa(global.Config.Server.Port)
	if err := r.Run(port); err != nil {
		global.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
