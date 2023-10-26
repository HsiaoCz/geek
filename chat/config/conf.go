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
