package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	apiTimeoutMsg = `{"code":-1, "msg":"timeout"}`
	apiTimeout    = time.Second * 30
)

// get is a shortcut for router.GET(path string, handle fasthttp.RequestHandler)
func get(router *fasthttprouter.Router, path string, handle fasthttp.RequestHandler) {
	router.GET(path, fasthttp.TimeoutHandler(handle, apiTimeout, apiTimeoutMsg))
}

// head is a shortcut for router.HEAD(path string, handle fasthttp.RequestHandler)
func head(router *fasthttprouter.Router, path string, handle fasthttp.RequestHandler) {
	router.HEAD(path, fasthttp.TimeoutHandler(handle, apiTimeout, apiTimeoutMsg))
}

// options is a shortcut for router.OPTIONS(path string, handle fasthttp.RequestHandler)
func options(router *fasthttprouter.Router, path string, handle fasthttp.RequestHandler) {
	router.OPTIONS(path, fasthttp.TimeoutHandler(handle, apiTimeout, apiTimeoutMsg))
}

// post is a shortcut for router.POST(path string, handle fasthttp.RequestHandler)
func post(router *fasthttprouter.Router, path string, handle fasthttp.RequestHandler) {
	router.POST(path, fasthttp.TimeoutHandler(handle, apiTimeout, apiTimeoutMsg))
}

// put is a shortcut for router.PUT(path string, handle fasthttp.RequestHandler)
func put(router *fasthttprouter.Router, path string, handle fasthttp.RequestHandler) {
	router.PUT(path, fasthttp.TimeoutHandler(handle, apiTimeout, apiTimeoutMsg))
}

// patch is a shortcut for router.PATCH(path string, handle fasthttp.RequestHandler)
func patch(router *fasthttprouter.Router, path string, handle fasthttp.RequestHandler) {
	router.PATCH(path, fasthttp.TimeoutHandler(handle, apiTimeout, apiTimeoutMsg))
}

// delete is a shortcut for router.DELETE(path string, handle fasthttp.RequestHandler)
func delete(router *fasthttprouter.Router, path string, handle fasthttp.RequestHandler) {
	router.DELETE(path, fasthttp.TimeoutHandler(handle, apiTimeout, apiTimeoutMsg))
}
