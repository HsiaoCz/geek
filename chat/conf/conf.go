package conf

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AllConf)

type AllConf struct {
	MySql MysqlConf `mapstructure:"mysql"`
	Redis RedisConf `mapstructure:"redis"`
	Mongo MongoConf `mapstructure:"mongo"`
	App   AppConf   `mapstructure:"app"`
}

type MysqlConf struct {
	MsName     string `mapstructure:"ms_name"`
	MsUser     string `mapstructure:"ms_user"`
	MsPassword string `mapstructure:"ms_passwd"`
	MsHost     string `mapstructure:"ms_host"`
	MsPort     string `mapstructure:"ms_port"`
}

type RedisConf struct {
	RdHost   string `mapstructure:"rd_host"`
	RdPort   string `mapstructure:"rd_port"`
	RdPasswd string `mapstructure:"rd_passwd"`
	RdName   int    `mapstructure:"rd_name"`
}

type MongoConf struct {
	MoHost   string `mapstructure:"mo_host"`
	MoPort   string `mapstructure:"mo_port"`
	MoUser   string `mapstructure:"mo_user"`
	MoPasswd string `mapstructure:"mo_passwd"`
	MoDbname string `mapstructure:"mo_dbname"`
}

type AppConf struct {
	AppAddr string `mapstructure:"app_addr"`
}

func Init() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/home/hsiaocz/go/src/geek/chat/conf")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err := viper.Unmarshal(Conf); err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(_ fsnotify.Event) {
		log.Println("配置文件被修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			return
		}
	})
	return err
}
