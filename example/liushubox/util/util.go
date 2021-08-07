package util

import (
	"github.com/lauthrul/goutil/http"
	"github.com/lauthrul/goutil/time"
	"github.com/lauthrul/goutil/validator"
	"github.com/valyala/fasthttp"
	"net/url"
)

const (
	// http responses
	CodeFail    = 0
	CodeSuccess = 1
)

func init() {
	validator.RegisterRule("datetime", CheckDateTime)
}

func CheckDateTime(date string) (bool, error) {
	if _, _, err := time.StrToTime(date); err != nil {
		return false, err
	}
	return true, nil
}

func Bind(ctx *fasthttp.RequestCtx, param interface{}) error {
	var (
		data string
		err  error
	)

	if string(ctx.Method()) == "GET" {
		data = string(ctx.QueryArgs().String())
	} else if string(ctx.Method()) == "POST" {
		data = string(ctx.PostArgs().String())
	}

	data, err = url.QueryUnescape(data)
	if err != nil {
		return err
	}

	return validator.Bind(data, param)
}

func EchoFail(ctx *fasthttp.RequestCtx, msg string) {
	http.Echo(ctx, CodeFail, msg, nil)
}

func EchoSuccess(ctx *fasthttp.RequestCtx, data interface{}) {
	http.Echo(ctx, CodeSuccess, "", data)
}
