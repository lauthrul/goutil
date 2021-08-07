package timetool

import (
	gtime "github.com/lauthrul/goutil/time"
	"github.com/valyala/fasthttp"
	"liushubox/util"
	"strconv"
	"time"
)

type timeAddParam struct {
	Time   string `name:"t" rule:"datetime"`
	Offset int    `name:"o"`
}

type timeDiffParam struct {
	T1 string `name:"t1" rule:"datetime"`
	T2 string `name:"t2" rule:"datetime"`
}

const (
	time2stamp = 0
	stamp2time = 1
)

type covertParam struct {
	Data string `name:"d"`
	NSec bool   `name:"n"`              // mill-second
	Type uint8  `name:"t" values:"0,1"` // 0:time->stamp, 1:stamp->time
}

func Add(ctx *fasthttp.RequestCtx) {
	var param timeAddParam
	if err := util.Bind(ctx, &param); err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}
	tm, layout, _ := gtime.StrToTime(param.Time)
	tm = tm.AddDate(0, 0, param.Offset)
	util.EchoSuccess(ctx, tm.Format(layout))
}

func Diff(ctx *fasthttp.RequestCtx) {
	var param timeDiffParam
	if err := util.Bind(ctx, &param); err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}
	t1, _, _ := gtime.StrToTime(param.T1)
	t2, _, _ := gtime.StrToTime(param.T2)
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
	util.EchoSuccess(ctx, resp)
}

func Covert(ctx *fasthttp.RequestCtx) {
	var param covertParam
	if err := util.Bind(ctx, &param); err != nil {
		util.EchoFail(ctx, err.Error())
		return
	}
	if param.Type == time2stamp { // time -> stamp
		t, _, err := gtime.StrToTime(param.Data)
		if err != nil {
			util.EchoFail(ctx, err.Error())
		}
		if param.NSec {
			util.EchoSuccess(ctx, t.UnixNano()/int64(time.Millisecond))
		} else {
			util.EchoSuccess(ctx, t.Unix())
		}
	} else { // stamp -> time
		n, err := strconv.ParseInt(param.Data, 10, 64)
		if err != nil {
			util.EchoFail(ctx, err.Error())
		}
		s := n
		ns := int64(0)
		if param.NSec {
			s = n / 1000
			ns = n % 1000 * int64(time.Millisecond)
		}
		t := time.Unix(s, ns)
		if param.NSec {
			util.EchoSuccess(ctx, t.Format("2006-01-02 15:04:05.000"))
		} else {
			util.EchoSuccess(ctx, t.Format("2006-01-02 15:04:05"))
		}
	}
}
