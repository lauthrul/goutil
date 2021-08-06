package http

import (
	"github.com/lauthrul/goutil/util"
	"github.com/valyala/fasthttp"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempy"`
	Data interface{} `json:"data,omitempy"`
}

func Echo(ctx *fasthttp.RequestCtx, code int, msg string, data interface{}) {
	ctx.SetStatusCode(200)
	res := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	bytes, err := util.Json.Marshal(res)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	ctx.SetBody(bytes)
}
