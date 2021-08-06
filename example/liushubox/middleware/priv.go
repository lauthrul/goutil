package middleware

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"liushubox/common"
	"liushubox/config"
	"strings"
)

func PrivMiddleware(ctx *fasthttp.RequestCtx) error {
	path := string(ctx.Path())
	allows := map[string]bool{
	}

	if _, ok := allows[path]; ok {
		return nil
	}

	if strings.Index(path, "/debug/") == 0 {
		token := string(ctx.QueryArgs().Peek("t"))
		stamp := string(ctx.QueryArgs().Peek("_"))
		if len(token) == 0 || len(stamp) == 0 {
			return errors.New("invalid request")
		}
		str := common.ProjectName + stamp + config.Get().Key
		sum := fmt.Sprintf("%x", md5.Sum([]byte(str)))
		if token != sum {
			return errors.New("invalid request")
		}
		return nil
	}

	return nil
}
