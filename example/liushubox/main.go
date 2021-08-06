package main

import (
	"github.com/lauthrul/goutil/log"
	"github.com/valyala/fasthttp"
	"liushubox/common"
	"liushubox/config"
	"liushubox/middleware"
	"liushubox/router"
	"time"
)

func main() {
	log.Init("")

	cfg, err := config.Load()
	if err != nil {
		log.Error("load config fail:", err)
	}

	rt := router.Init()
	srv := &fasthttp.Server{
		Handler:            middleware.Handle(rt.Handler),
		ReadTimeout:        5 * time.Second,
		WriteTimeout:       10 * time.Second,
		Name:               common.ProjectName,
		MaxRequestBodySize: 51 * 1024 * 1024,
	}

	log.Info(common.ProjectName, "is running at", cfg.Addr)
	if err := srv.ListenAndServe(cfg.Addr); err != nil {
		log.Fatal("run server error:", err)
	}
}
