package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"database"`
	Security struct {
		JWT struct {
			SecretKey string `mapstructure:"secretKey"`
			ExpiresIn string `mapstructure:"expiresIn"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	v := viper.New()
	v.AddConfigPath("./config/") // path to config file
	v.SetConfigName("local")     // name of config file
	v.SetConfigType("yaml")

	// read configuration
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s", err))
	}

	// print server configuration
	fmt.Printf("Server port is: %v\n", config.Server.Port)

	// print database configurations
	for _, db := range config.Databases {
		fmt.Printf("Database User: %s, Password: %s, Host: %s, DbName: %s\n", db.User, db.Password, db.Host, db.DbName)
	}

	// print security configuration
	fmt.Printf("JWT SecretKey: %s, ExpiresIn: %s\n", config.Security.JWT.SecretKey, config.Security.JWT.ExpiresIn)
}
