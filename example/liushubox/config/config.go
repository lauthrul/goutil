package config

import (
	"github.com/lauthrul/goutil/util"
	"io/ioutil"
	"liushubox/common"
)

type Config struct {
	Ver   string `json:"ver"`
	Build string `json:"build"`
	Key   string `json:"key"`
	Addr  string `json:"addr"`
}

var (
	conf *Config
)

func Load() (*Config, error) {
	conf = &Config{
		Ver:   "0.0.0",
		Build: "",
		Key:   "",
		Addr:  ":80",
	}
	data, err := ioutil.ReadFile(common.ConfigFile)
	if err != nil {
		return conf, err
	}
	if err = util.Json.Unmarshal(data, conf)
		err != nil {
		return conf, err
	}
	return conf, err
}

func Get() *Config {
	if conf == nil {
		Load()
	}
	return conf
}
