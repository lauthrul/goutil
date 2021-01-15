package main

import (
	"fmt"
	"github.com/lauthrul/goutil/log"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const (
	url = "http://hq.sinajs.cn/list="
)

var (
	indexes = []map[string]*Index{
		{"s_sh000001": nil}, // 上证指数
		{"s_sz399001": nil}, // 深证指数
	}

	stocks = []map[string]*Stock{
		{"sh600291": nil}, // 西水股份
		{"sh603393": nil}, // 新天然气
		{"sh603551": nil},
		{"sh601872": nil},
		{"sh601519": nil},
		{"sh600968": nil},
		{"sh600900": nil},
		{"sh512290": nil},
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
				for i, _ := range indexes {
					if _, ok := indexes[i][match[1]]; ok {
						indexes[i][match[1]] = &index
					}
				}
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
				for i, _ := range stocks {
					if _, ok := stocks[i][match[1]]; ok {
						stocks[i][match[1]] = &stock
					}
				}
			}
		}
	}
}

func display() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(EnumNames(Stock{}))
	table.SetAlignment(tablewriter.ALIGN_RIGHT)

	for _, v := range stocks {
		for _, vv := range v {
			table.Append(EnumValues(*vv))
		}
	}
	table.Render()

	for _, v := range indexes {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}

func main() {
	log.Init("")
	log.SetLevel(log.LevelDebug)

	do := func() {
		req := url
		for _, v := range indexes {
			for k, _ := range v {
				req += k + ","
			}
		}
		resp, err := httpGet(req)
		if err != nil {
			log.ErrorF("httpGet err=", err)
			return
		}
		parseIndexes(resp)

		req = url
		for _, v := range stocks {
			for k, _ := range v {
				req += k + ","
			}
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
