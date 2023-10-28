package config

// Conf对外使用的配置文件
// 通过这个变量访问
var Conf = new(Chat)

type Chat struct {
	App AppConf `mapstructure:"app"`
}

type AppConf struct {
	Port string `mapstructure:"port"`
}

type MysqlConf struct {
	User   string `mapstructure:"user"`
	Passwd string `mapstructure:"passwd"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	DBname string `mapstructure:"dbname"`
}

type RedisConf struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Passwd string `mapstructure:"passwd"`
	DB     int    `mapstructure:"db"`
}
