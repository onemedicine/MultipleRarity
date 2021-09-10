package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Rarity Summoner `yaml:"summoner"`
}

type Summoner struct {
	Id []uint64 `yaml:"id"`
	Private string `yaml:"private"`
}

var Conf *Config

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("config/")
	v.SetConfigType("yaml")

	if e := v.ReadInConfig(); e != nil {
		panic(e)
	}

	if e := v.Unmarshal(&Conf); e != nil {
		panic(e)
	}
}
