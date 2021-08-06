package middleware

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type handler func(ctx *fasthttp.RequestCtx) error

var handlers = []handler{
	PrivMiddleware,
}

func Handle(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		for _, cb := range handlers {
			if err := cb(ctx); err != nil {
				fmt.Fprint(ctx, err)
				return
			}
		}
		next(ctx)
	})
}
