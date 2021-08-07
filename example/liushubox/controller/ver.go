package controller

import (
	"github.com/valyala/fasthttp"
	"liushubox/config"
	"liushubox/util"
	"runtime"
	"time"
)

func Ver(ctx *fasthttp.RequestCtx) {
	data := map[string]interface{}{
		"ver":  config.Get().Ver,
		"build": config.Get().Build,
		"go":   runtime.Version(),
		"time": time.Now(),
	}
	util.EchoSuccess(ctx, data)
}
