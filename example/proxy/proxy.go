package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"github.com/lauthrul/goutil/log"
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"
	"time"
)

const (
	modHttp  = "http"
	modProxy = "proxy"
)

var (
	addr string
	mode string
)

func runModProxy() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	var server *http.Server
	server = &http.Server{
		Addr:    addr,
		Handler: proxy,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("proxy server error:", err)
	}
}

func runModHttp() {
	handler := func(ctx *fasthttp.RequestCtx) {

		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		defer func() {
			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)
		}()

		uri := string(ctx.RequestURI())
		uri = strings.TrimLeft(uri, "/")
		method := string(ctx.Method())
		req.Header.SetMethod(method)
		req.SetRequestURI(uri)
		req.SetBody(ctx.Request.Body())

		err := fasthttp.Do(req, resp)
		if err != nil {
			log.ErrorF("req:%s, err:%v", req.URI().FullURI(), err)
		} else {
			log.InfoF("req:%s, status:%d, bytes:%d", req.URI().FullURI(), resp.StatusCode(), len(resp.Body()))
		}

		ctx.Response.SetStatusCode(resp.StatusCode())
		ctx.Response.SetBody(resp.Body())
	}

	srv := &fasthttp.Server{
		Handler:            handler,
		ReadTimeout:        5 * time.Second,
		WriteTimeout:       10 * time.Second,
		MaxRequestBodySize: 51 * 1024 * 1024,
	}

	if err := srv.ListenAndServe(addr); err != nil {
		log.Error("proxy server error:", err)
	}
}

func main() {

	flag.StringVar(&mode, "m", "http", "proxy mode. [http|proxy]")
	flag.StringVar(&addr, "a", "127.0.0.1:8000", "proxy bind addr")
	flag.Parse()

	usage := func() {
		flag.Usage()
	}
	if mode != modHttp && mode != modProxy {
		usage()
		return
	}
	if addr == "" {
		usage()
		return
	}

	log.Init("")
	log.Info("mode=", mode)
	log.Info("addr=", addr)

	switch mode {
	case modProxy:
		runModProxy()
	case modHttp:
		runModHttp()
	}
}
