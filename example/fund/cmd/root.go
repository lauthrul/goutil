package cmd

import (
	"fund/common"
	"fund/config"
	"fund/ui"
	"github.com/lauthrul/goutil/http"
	"github.com/lauthrul/goutil/log"
	"github.com/spf13/cobra"
	"time"
)

func Init(conf config.Config) {
	log.Init(conf.LogFile)
	if conf.Verbose {
		log.SetLevel(log.LevelDebug)
	}
	log.Debug("cacheFile =", conf.CacheFile)

	common.Client = &http.Client{Timeout: 10 * time.Second}
	if conf.EnableProxy {
		log.Debug("enable proxy:", conf.Proxy)
		common.Client.Proxy = conf.Proxy
	}
	common.Client.Init()

	//cache, err = LoadCache(cacheFile)
	//if err != nil {
	//	return
	//}
}

func RootCmd() *cobra.Command {
	var configFile string

	rootCmd := &cobra.Command{
		Use:   "fund",
		Short: "An effective fund manage tool",
		Long:  "An effective fund manage tool",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)
			ui.Run()
		},
	}

	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")

	return rootCmd
}
