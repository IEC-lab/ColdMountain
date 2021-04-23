package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

var globalConfig = &Config{}

func init() {
	if _, err := toml.DecodeFile("./conf/config.toml", globalConfig); err != nil {
		log.Panicf("global config init error: %+v", err)
	}
}

type Config struct {
	ColdRepo ColdRepo `toml:"cold_repo"`
	ColdQ    ColdQ    `toml:"cold_q"`
}

type ColdRepo struct {
	Address  string
	User     string
	Password string
	Port     string
}

type ColdQ struct {
	Address string
	Port    string
}

func GetGlobalConfig() *Config {
	return globalConfig
}
