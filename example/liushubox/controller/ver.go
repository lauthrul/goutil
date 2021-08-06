package controller

import (
	"github.com/lauthrul/goutil/http"
	"github.com/valyala/fasthttp"
	"liushubox/common"
	"liushubox/config"
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
	http.Echo(ctx, common.CodeSuccess, "", data)
}
