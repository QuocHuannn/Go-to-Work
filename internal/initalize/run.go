package initalize

import (
	"fmt"

	"github.com/QuocHuannn/Go-to-Work/global"
)

func Run() {
	//Load configuration
	LoadConfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()
	r.Run(":8002")
}
