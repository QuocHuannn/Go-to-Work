package initalize

import (
	"fmt"

	"github.com/QuocHuannn/Go-to-Work/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./config/") // path to config file
	v.SetConfigName("local")     // name of config file
	v.SetConfigType("yaml")

	// Enable environment variables to override config
	v.AutomaticEnv()
	v.SetEnvPrefix("") // No prefix for env vars

	// Map environment variables to config keys
	v.BindEnv("mysql.host", "MYSQL_HOST")
	v.BindEnv("mysql.port", "MYSQL_PORT")

	// read configuration
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s", err))
	}

	// Debug output to verify config
	fmt.Printf("MySQL Host: %s, Port: %d\n", global.Config.Mysql.Host, global.Config.Mysql.Port)
}
