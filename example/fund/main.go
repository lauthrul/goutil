package main

import (
	"flag"
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

	log.Init("")
	if verbose {
		log.SetLevel(log.LevelDebug)
	}

	log.Debug("proxy =", proxy)
	log.Debug("cacheFile =", cacheFile)

	client = &http.Client{Proxy: proxy, Timeout: 10 * time.Second}
	client.Init()

	//cache, err = LoadCache(cacheFile)
	//if err != nil {
	//	return
	//}
	//
	//ok, err := FundListUpdateCheck(cache)
	//if err != nil {
	//	return
	//}
	//if ok {
	//	_ = SaveCache(cache, cacheFile)
	//}

	app := App{
		updateInterval: 10 * time.Second,
		updateFunc: func(app *App) {
			funds, err := FundMarketList()
			if err != nil {
				return
			}
			app.Update(funds)
		},
	}
	app.Init()
	app.Run()
}
