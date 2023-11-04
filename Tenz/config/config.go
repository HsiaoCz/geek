package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(App)

type App struct {
	AC AppConf   `mapstructure:"app"`
	MC MysqlConf `mapstructure:"mysql"`
	RC RedisConf `mapstructure:"redis"`
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

func InitConf() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/home/hsiaocz/go/src/geek/Tenz")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err := viper.Unmarshal(Conf); err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			return
		}
	})
	return err
}
