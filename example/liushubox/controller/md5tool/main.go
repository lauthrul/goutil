package md5tool

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/lauthrul/goutil/http"
	"github.com/valyala/fasthttp"
	"liushubox/common"
	"liushubox/util"
)

type md5Req struct {
	Data string `name:"data" min:"1" max:"1024"`
	Bit  int    `name:"bit" values:"16,32"`
}

func CheckSum(ctx *fasthttp.RequestCtx) {
	var param md5Req
	if err := util.Bind(ctx, &param); err != nil {
		http.Echo(ctx, common.CodeFail, err.Error(), nil)
		return
	}

	h := md5.New()
	h.Write([]byte(param.Data))
	checksum := hex.EncodeToString(h.Sum(nil))
	if param.Bit == 16 {
		checksum = checksum[8:24]
	}

	http.Echo(ctx, common.CodeSuccess, "", checksum)
}
