package timetool

import (
	"github.com/lauthrul/goutil/http"
	"github.com/lauthrul/goutil/time"
	"github.com/valyala/fasthttp"
	"liushubox/common"
	"liushubox/util"
)

type timeAddParam struct {
	Time   string `name:"t" rule:"datetime"`
	Offset int    `name:"o"`
}

type timeDiffParam struct {
	T1 string `name:"t1" rule:"datetime"`
	T2 string `name:"t2" rule:"datetime"`
}

func Add(ctx *fasthttp.RequestCtx) {
	var param timeAddParam
	if err := util.Bind(ctx, &param); err != nil {
		http.Echo(ctx, common.CodeFail, err.Error(), nil)
		return
	}
	tm, layout, _ := time.StrToTime(param.Time)
	tm = tm.AddDate(0, 0, param.Offset)
	http.Echo(ctx, common.CodeSuccess, "", tm.Format(layout))
}

func Diff(ctx *fasthttp.RequestCtx) {
	var param timeDiffParam
	if err := util.Bind(ctx, &param); err != nil {
		http.Echo(ctx, common.CodeFail, err.Error(), nil)
		return
	}
	t1, _, _ := time.StrToTime(param.T1)
	t2, _, _ := time.StrToTime(param.T2)
	t := t2.Sub(t1)

	d := int(t.Hours()) / 24
	h := int(t.Hours()) - d*24
	m := int(t.Minutes()) - d*24*60 - h*60
	s := int(t.Seconds()) - d*24*60*60 - h*60*60 - m*60

	resp := map[string]int{
		"d": d,
		"h": h,
		"m": m,
		"s": s,
	}
	http.Echo(ctx, common.CodeSuccess, "", resp)
}
