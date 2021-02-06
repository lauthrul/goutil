package main

import (
	"flag"
	"fmt"
	"github.com/lauthrul/goutil/log"
	"github.com/lauthrul/goutil/util"
	"github.com/olekukonko/tablewriter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	url = "http://hq.sinajs.cn/list="
)

var (
	proxy   string
	indexes []*Index
	stocks  []*Stock
)

func httpGet(uri string) (string, error) {
	cli := &fasthttp.Client{}
	if proxy != "" {
		cli.Dial = fasthttpproxy.FasthttpHTTPDialer(proxy)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)

	resp := fasthttp.AcquireResponse()
	err := cli.Do(req, resp)
	if err != nil {
		log.Error(req.URI(), err)
		return "", err
	}

	bytes := resp.Body()
	data := util.Convert(bytes, util.GB18030)
	//log.Debug(resp)
	return data, nil
}

var (
	reg = regexp.MustCompile(`_str_(.*?)="(.*)"`)
)

func parseIndexes(str string) {
	indexes = nil
	matches := reg.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if len(match) >= 3 {
			var index Index
			if index.Parse(strings.Trim(match[2], ",")) == nil {
				index.Code = match[1]
				indexes = append(indexes, &index)
			}
		}
	}
}

func parseStocks(str string) {
	stocks = nil
	matches := reg.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if len(match) >= 3 {
			var stock Stock
			if stock.Parse(strings.Trim(match[2], ",")) == nil {
				stock.Code = match[1]
				stocks = append(stocks, &stock)
			}
		}
	}
}

func display() {
	util.ClearConsole()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(util.EnumNames(Stock{}))
	table.SetAlignment(tablewriter.ALIGN_RIGHT)

	for _, v := range stocks {
		table.Append(util.EnumValues(*v))
	}
	table.Render()

	for _, v := range indexes {
		fmt.Println(v)
	}
}

func main() {
	flag.StringVar(&proxy, "proxy", "", "proxy host:port")
	flag.Parse()

	log.Init("")
	log.SetLevel(log.LevelDebug)

	config := Config{FilePath: "config.json"}
	config.Load()

	do := func() {
		req := url
		for _, v := range config.Indexes {
			req += v + ","
		}
		resp, err := httpGet(req)
		if err != nil {
			log.Error("httpGet err=", err)
			return
		}
		parseIndexes(resp)

		req = url
		for _, v := range config.Stocks {
			req += v + ","
		}
		resp, err = httpGet(req)
		if err != nil {
			log.Error("httpGet err=", err)
			return
		}
		parseStocks(resp)

		display()
	}

	for {
		select {
		case <-time.Tick(3 * time.Second):
			do()
		}
	}
}
