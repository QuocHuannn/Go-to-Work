package setting

type Config struct{
	Mysql MySQLSetting `mapstructure:"mysql"`
}

type MySQLSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	ConnectTimeout int `mapstructure:"connect_timeout"`
}