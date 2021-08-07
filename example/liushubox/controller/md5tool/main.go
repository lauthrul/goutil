package md5tool

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/valyala/fasthttp"
	"liushubox/util"
)

type md5Param struct {
	Data string `name:"d" min:"1" max:"1024"`
	Bit  int    `name:"b" values:"16,32" empty:"1" default:"32"`
}

type md5FileParam struct {
	Bit int `name:"b" values:"16,32" empty:"1" default:"32"`
}

func fnMd5(data []byte, bit int) string {
	h := md5.New()
	h.Write(data)
	checksum := hex.EncodeToString(h.Sum(nil))
	if bit == 16 {
		checksum = checksum[8:24]
	}
	return checksum
}

func Md5(ctx *fasthttp.RequestCtx) {
	var param md5Param
	if err := util.Bind(ctx, &param); err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}

	checksum := fnMd5([]byte(param.Data), param.Bit)
	util.EchoSuccess(ctx, checksum)
}

func Md5File(ctx *fasthttp.RequestCtx) {
	var param md5FileParam
	if err := util.Bind(ctx, &param); err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}

	var data []byte
	f, err := file.Open()
	if err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}

	_, err = f.Read(data)
	if err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}
	_ = f.Close()

	checksum := fnMd5(data, param.Bit)
	util.EchoSuccess(ctx, checksum)
}
