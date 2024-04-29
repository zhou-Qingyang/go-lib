package viper

type Config struct {
	AppConfig AppConfig
	Mysql     Mysql
	Redis     Redis
}

type Mysql struct {
	User     string
	Password string
	Database string
}

type Redis struct {
	Ip       string
	Port     int
	UserName string
}

type AppConfig struct {
	AppName  string
	LogLevel string
}
