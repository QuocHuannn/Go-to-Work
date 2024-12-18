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

	// read configuration
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s", err))
	}
}
