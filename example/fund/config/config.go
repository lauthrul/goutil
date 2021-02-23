package config

import (
	"fmt"
	"github.com/lauthrul/goutil/util"
	"io/ioutil"
)

type Config struct {
	EnableProxy bool   `json:"enable_proxy"` // enable proxy
	Proxy       string `json:"proxy"`        // http proxy [host:port]
	LogFile     string `json:"log_file"`     // log file, default: log.txt
	DbFile      string `json:"db_file"`      // sqlite3 database file, default: fund.db
	CacheFile   string `json:"cache_file"`   // cache file, default: fund.cache
	Verbose     bool   `json:"verbose"`      // log details as much as possible
}

func Load(filePath string) (config Config) {
	defer func() {
		if config.LogFile == "" {
			config.LogFile = "log.txt"
		}
		if config.DbFile == "" {
			config.DbFile = "fund.db"
		}
		if config.CacheFile == "" {
			config.CacheFile = "fund.cache"
		}
	}()
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read config file[%s] fail: %s\n", filePath, err.Error())
		return config
	}
	err = util.Json.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Printf("read config file[%s] fail: %s\n", filePath, err.Error())
	}
	return config
}
