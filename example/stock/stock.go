package main

import (
	"fmt"
	"github.com/lauthrul/goutil/log"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	url = "http://hq.sinajs.cn/list="
)

var (
	indexes = map[string]*Index{
		"s_sh000001": nil, // 上证指数
		"s_sz399001": nil, // 深证指数
	}
	stocks = map[string]*Stock{
		"sh600291": nil, // 西水股份
		"sh603393": nil, // 新天然气
	}
)

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := Convert(bytes, GB18030)
	//log.Debug(resp)
	return data, nil
}

var (
	reg = regexp.MustCompile(`_str_(.*?)="(.*)"`)
)

func parseIndexes(str string) {
	matches := reg.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if len(match) >= 3 {
			var index Index
			if index.Parse(strings.Trim(match[2], ",")) == nil {
				indexes[match[1]] = &index
			}
		}
	}
}

func parseStocks(str string) {
	matches := reg.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if len(match) >= 3 {
			var stock Stock
			if stock.Parse(strings.Trim(match[2], ",")) == nil {
				stocks[match[1]] = &stock
			}
		}
	}
}

func display() {
	s := "\r"
	for _, v := range indexes {
		s += v.String() + "\n"
	}
	for _, v := range stocks {
		s += v.String() + "\n"
	}
	fmt.Print(s)
}

func main() {
	log.Init("")
	log.SetLevel(log.LevelDebug)

	do := func() {
		req := url
		for k, _ := range indexes {
			req += k + ","
		}
		resp, err := httpGet(req)
		if err != nil {
			log.ErrorF("httpGet err=", err)
			return
		}
		parseIndexes(resp)

		req = url
		for k, _ := range stocks {
			req += k + ","
		}
		resp, err = httpGet(req)
		if err != nil {
			log.ErrorF("httpGet err=", err)
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
