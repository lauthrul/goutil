package main

import (
	"flag"
	"fund/common"
	"fund/ui"
	"github.com/lauthrul/goutil/http"
	"github.com/lauthrul/goutil/log"
	"time"
)

func main() {
	var (
		proxy     string
		cacheFile string
		verbose   bool
	)

	flag.StringVar(&proxy, "proxy", "", "Proxy host:port")
	flag.StringVar(&cacheFile, "cache", "fundlist.cache", "Cache file")
	flag.BoolVar(&verbose, "v", false, "Make the operation more talkative")
	flag.Parse()

	log.Init("log.txt")
	if verbose {
		log.SetLevel(log.LevelDebug)
	}

	log.Debug("proxy =", proxy)
	log.Debug("cacheFile =", cacheFile)

	common.Client = &http.Client{Proxy: proxy, Timeout: 10 * time.Second}
	common.Client.Init()

	//cache, err = LoadCache(cacheFile)
	//if err != nil {
	//	return
	//}

	ui.Run()
}
