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
	indexes []*Index
	stocks  []*Stock
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
				index.Code = match[1]
				indexes = append(indexes, &index)
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
				stock.Code = match[1]
				stocks = append(stocks, &stock)
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
		table.Append(EnumValues(*v))
	}
	table.Render()

	for _, v := range indexes {
		fmt.Println(v)
	}
}

func main() {
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
			log.ErrorF("httpGet err=", err)
			return
		}
		parseIndexes(resp)

		req = url
		for _, v := range config.Stocks {
			req += v + ","
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
