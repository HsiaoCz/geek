package config

var Conf = new(Lottery)

type Lottery struct {
	App   AppConf   `mapstructure:"app"`
	Mysql MysqlConf `mapstructure:"mysql"`
	Redis RedisConf `mapstructure:"redis"`
}

type AppConf struct{}

type MysqlConf struct{}

type RedisConf struct{}
