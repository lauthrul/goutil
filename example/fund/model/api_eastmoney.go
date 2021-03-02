package model

import (
	"fmt"
	"fund/common"
	"fund/lang"
	"github.com/antchfx/htmlquery"
	"github.com/lauthrul/goutil/log"
	"github.com/lauthrul/goutil/util"
	"golang.org/x/net/html"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type EastMoneyFund struct {
	Code            string // 基金代码
	Name            string // 基金简称
	NetDate         string // 日期
	NetValue        string // 单位净值
	TotalNetValue   string // 累计净值
	DayRate         string // 日增长率
	WeekRate        string // 近1周
	MonthRate       string // 近1月
	ThreeMonthRate  string // 近3月
	SixMonthRate    string // 近6月
	YearRate        string // 近1年
	TwoYearRate     string // 近2年
	ThreeYearRate   string // 近3年
	ThisYearRate    string // 今年来
	SinceCreateRate string // 成立来
	CreateDate      string // 成立日期
	CustomRate      string // 自定义（查询时间段）
	FeeRate         string // 手续费
}

type EastMoneyFundBasic struct {
	Code        string `json:"FCODE"`     // 代码
	Name        string `json:"SHORTNAME"` // 简称
	Type        string `json:"FTYPE"`     // 类型
	CreateDate  string `json:"ESTABDATE"` // 成立日期
	CreateScale string `json:"NETNAV"`    // 成立规模
	LatestScale string `json:"ENDNAV"`    // 最新规模
	UpdateDate  string `json:"FEGMRQ"`    // 更新日期
	CompanyCode string `json:"JJGSID"`    // 基金公司代码
	CompanyName string `json:"JJGS"`      // 基金公司名称
	ManagerID   string `json:"JJJLID"`    // 基金经理id
	ManagerName string `json:"JJJL"`      // 基金经理
	ManageExp   string `json:"MGREXP"`    // 管理费率（年）
	TrustExp    string `json:"TRUSTEXP"`  // 托管费率（年）
}

func (e EastMoneyFund) GetTHReference() []THReference {
	return []THReference{
		{lang.Text(common.Lan, "thCode"), true, "dm"},
		{lang.Text(common.Lan, "thName"), true, "jc"},
		{lang.Text(common.Lan, "thDate"), false, ""},
		{lang.Text(common.Lan, "thNet"), true, "dwjz"},
		{lang.Text(common.Lan, "thTotalNet"), true, "ljjz"},
		{lang.Text(common.Lan, "thDateRate"), true, "rzdf"},
		{lang.Text(common.Lan, "thWeekRate"), true, "zzf"},
		{lang.Text(common.Lan, "thMonthRate"), true, "1yzf"},
		{lang.Text(common.Lan, "th3MonthRate"), true, "3yzf"},
		{lang.Text(common.Lan, "th6MonthRate"), true, "6yzf"},
		{lang.Text(common.Lan, "thYearRate"), true, "1nzf"},
		{lang.Text(common.Lan, "th2YearRate"), true, "2nzf"},
		{lang.Text(common.Lan, "th3YearRate"), true, "3nzf"},
		{lang.Text(common.Lan, "thThisYearRate"), true, "jnzf"},
		{lang.Text(common.Lan, "thSinceFoundRate"), true, "lnzf"},
		{lang.Text(common.Lan, "thCustomRate"), true, "qjzf"},
		{lang.Text(common.Lan, "thFeeRate"), false, ""},
	}
}

func (e EastMoneyFund) GetValues() []string {
	name := []rune(e.Name)
	if len(name) > 10 {
		name = name[:10]
		e.Name = string(name) + "..."
	}
	return []string{
		e.Code,
		e.Name,
		e.NetDate,
		e.NetValue,
		e.TotalNetValue,
		e.DayRate,
		e.WeekRate,
		e.MonthRate,
		e.ThreeMonthRate,
		e.SixMonthRate,
		e.YearRate,
		e.TwoYearRate,
		e.ThreeYearRate,
		e.ThisYearRate,
		e.SinceCreateRate,
		e.CustomRate,
		e.FeeRate,
	}
}

type EastMoneyManager struct {
	ID        string  `json:"MGRID"`
	Name      string  `json:"MGRNAME"`
	FromDate  string  `json:"FEMPDATE"`
	Days      float64 `json:"DAYS"`
	Education string  `json:"EDUCATION"`
	Growth    float64 `json:"GROWTH"`
	Resume    string  `json:"RESUME"`
	FundList  []struct {
		FromDate string  `json:"FEMPDATE"`
		ToDate   string  `json:"LEMPDATE"`
		FundCode string  `json:"FCODE"`
		FundName string  `json:"SHORTNAME"`
		Growth   float64 `json:"PENAVGROWTH"`
	}
}

type EastMoneyFundNetValue struct {
	Code          string
	Date          string `json:"FSRQ"`
	NetValue      string `json:"DWJZ"`
	TotalNetValue string `json:"LJJZ"`
	Growth        string `json:"JZZZL"`
}

type EastMoneyFundNetValueList struct {
	Datas struct {
		List []EastMoneyFundNetValue `json:"LSJZList"`
	} `json:"Data"`
	TotalCount int `json:"TotalCount"`
	PageSize   int `json:"PageSize"`
	PageIndex  int `json:"PageIndex"`
}

// 基金估值
type EastMoneyFundEstimate struct {
	Code         string `json:"fundcode"` // 基金代码 	519983
	Name         string `json:"name"`     // 基金简称 	长信量化先锋混合A
	NetDate      string `json:"jzrq"`     // 日期 		2018-09-21
	NetValue     string `json:"dwjz"`     // 单位净值 	1.2440
	EstimateNet  string `json:"gsz"`      // 估算净值 	1.2388
	EstimateRate string `json:"gszzl"`    // 估算增长率	-0.42
	EstimateDate string `json:"gztime"`   // 估值日期 	2018-09-25 15:00
}

type EastMoneyApi struct {
	urlFundSearch   string
	urlRank         string
	urlBasic        string
	urlManager      string
	urlEstimate     string
	urlHoldingStock string
	urlNetValue     string
	referer         string
}

func NewEastMoneyApi() *EastMoneyApi {
	return &EastMoneyApi{
		urlFundSearch:   "http://fund.eastmoney.com/js/fundcode_search.js?v=%s",                                                                                                       // http://fund.eastmoney.com/js/fundcode_search.js?v=20210204133444
		urlRank:         "http://fund.eastmoney.com/data/rankhandler.aspx",                                                                                                            // http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=all&rs=&gs=0&sc=rzdf&st=desc&sd=2020-02-18&ed=2021-02-18&qdii=&tabSubtype=,,,,,&pi=1&pn=10&dx=1&v=0.29261668485886694
		urlBasic:        "http://fundmobapi.eastmoney.com/FundMApi/FundDetailInformation.ashx?callback=jQuery%d&FCODE=%s&deviceid=Wap&plat=Wap&product=EFund&version=2.0.0&Uid=&_=%d", // http://fundmobapi.eastmoney.com/FundMApi/FundDetailInformation.ashx?callback=jQuery3110907213304748691_1614149729392&FCODE=161725&deviceid=Wap&plat=Wap&product=EFund&version=2.0.0&Uid=&_=1614149729394
		urlManager:      "http://fundmobapi.eastmoney.com/FundMApi/FundMangerDetail.ashx?callback=jQuery%d&FCODE=%s&deviceid=Wap&plat=Wap&product=EFund&version=2.0.0&Uid=&_=%d",      // http://fundmobapi.eastmoney.com/FundMApi/FundMangerDetail.ashx?callback=jQuery31108228140360881444_1614155764008&FCODE=004707&deviceid=Wap&plat=Wap&product=EFund&version=2.0.0&Uid=&_=1614155764009
		urlEstimate:     "http://fundgz.1234567.com.cn/js/%s.js?rt=%s",                                                                                                                // http://fundgz.1234567.com.cn/js/007047.js?rt=1612108800
		urlHoldingStock: "http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=%s&topline=%d&year=%d&month=&rt=%s",                                                      // http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=007047&topline=10&year=2020&month=1&rt=0.6304561761132715
		urlNetValue:     "http://api.fund.eastmoney.com/f10/lsjz?callback=jQuery%d&fundCode=%s&pageIndex=%d&pageSize=%d&startDate=%s&endDate=%s&_=%d",                                 // http://api.fund.eastmoney.com/f10/lsjz?callback=jQuery18305743557371800165_1613632567355&fundCode=003494&pageIndex=1&pageSize=20&startDate=2021-02-17&endDate=2021-02-17&_=1613632577360
		referer:         "http://fundf10.eastmoney.com/",
	}
}

func (api *EastMoneyApi) GetRandom() string {
	return fmt.Sprintf("0.%d", time.Now().Unix())
}

func (api *EastMoneyApi) GetTHReference() []THReference {
	return EastMoneyFund{}.GetTHReference()
}

func (api *EastMoneyApi) GetFundRank(arg FundRankArg) (FundList, error) {
	var funds FundList
	url := api.urlRank + "?" + arg.String()
	resp, err := common.Client.Get(url, map[string]string{"Referer": api.referer})
	if err != nil {
		log.ErrorF("%s: %s", err, url)
		return funds, err
	}

	body := string(resp.Body())
	reg := regexp.MustCompile(`(\[.*?\]).*pageIndex:(\d+).*allPages:(\d+).*allNum:(\d+)`)
	results := reg.FindAllStringSubmatch(body, -1)
	if len(results) == 1 {
		matches := results[0]
		if len(matches) == 5 {
			var items []string
			err = util.Json.UnmarshalFromString(matches[1], &items)
			if err != nil {
				log.ErrorF("[url] %s: %s", url, err, matches[1])
				return funds, err
			}
			for _, item := range items {
				values := strings.Split(item, ",")
				if len(values) == 25 {
					/*
						0	165520,
						1	信诚中证800有色指数(LOF),
						2	XCZZ800YSZSLOF,
						3	2021-02-18,
						4	1.6230,
						5	1.6160,
						6	7.13,
						7	7.13,
						8	15.27,
						9	38.55,
						10	50.65,
						11	67.58,
						12	108.32,
						13	54.08,
						14	25.13,
						15	76.8607,
						16	2013-08-30,
						17	1,
						18	67.5755,
						19	1.20%,
						20	0.12%,
						21	1,
						22	0.12%,
						23	1,
						24	106.12
					*/
					fund := EastMoneyFund{
						Code:            values[0],
						Name:            values[1],
						NetDate:         values[3],
						NetValue:        values[4],
						TotalNetValue:   values[5],
						DayRate:         values[6],
						WeekRate:        values[7],
						MonthRate:       values[8],
						ThreeMonthRate:  values[9],
						SixMonthRate:    values[10],
						YearRate:        values[11],
						TwoYearRate:     values[12],
						ThreeYearRate:   values[13],
						ThisYearRate:    values[14],
						SinceCreateRate: values[15],
						CreateDate:      values[16],
						CustomRate:      values[18],
						FeeRate:         values[20],
					}
					funds.List = append(funds.List, fund)
				}
			}
			funds.PageIndex, _ = strconv.Atoi(matches[2])
			funds.TotalPage, _ = strconv.Atoi(matches[3])
			funds.TotalCount, _ = strconv.Atoi(matches[4])
		}
	}

	return funds, nil
}

func (api *EastMoneyApi) GetFundBasic(fundCode string) (FundBasic, error) {
	var fundBasic FundBasic
	stamp := time.Now().Unix()
	url := fmt.Sprintf(api.urlBasic, stamp, fundCode, stamp)
	resp, err := common.Client.Get(url, map[string]string{"Referer": api.referer})
	if err != nil {
		log.ErrorF("%s: %s", err, url)
		return fundBasic, err
	}

	data := string(resp.Body())
	data = data[strings.Index(data, "(")+1 : strings.LastIndex(data, ")")]
	var result struct {
		Datas EastMoneyFundBasic `json:"Datas"`
	}
	err = util.Json.UnmarshalFromString(data, &result)
	if err != nil {
		log.ErrorF("[%s] %s: %s", url, err, data)
		return fundBasic, err
	}
	fundBasic.Code = result.Datas.Code
	fundBasic.Name = result.Datas.Name
	fundBasic.Type = result.Datas.Type
	fundBasic.CreateDate = result.Datas.CreateDate
	fundBasic.CreateScale, _ = strconv.ParseFloat(result.Datas.CreateScale, 64)
	fundBasic.LatestScale, _ = strconv.ParseFloat(result.Datas.LatestScale, 64)
	fundBasic.UpdateDate = result.Datas.UpdateDate
	fundBasic.CompanyCode = result.Datas.CompanyCode
	fundBasic.CompanyName = result.Datas.CompanyName
	fundBasic.ManagerName = result.Datas.ManagerName
	fundBasic.ManageExp, _ = strconv.ParseFloat(strings.TrimRight(result.Datas.ManageExp, "%"), 64)
	fundBasic.TrustExp, _ = strconv.ParseFloat(strings.TrimRight(result.Datas.TrustExp, "%"), 64)
	return fundBasic, nil
}

func (api *EastMoneyApi) GetFundManager(fundCode string) ([]Manager, []ManagerExperience, error) {
	var (
		managers    []Manager
		experiences []ManagerExperience
	)
	stamp := time.Now().Unix()
	url := fmt.Sprintf(api.urlManager, stamp, fundCode, stamp)
	resp, err := common.Client.Get(url, map[string]string{"Referer": api.referer})
	if err != nil {
		log.ErrorF("%s: %s", err, url)
		return nil, nil, err
	}

	data := string(resp.Body())
	data = data[strings.Index(data, "(")+1 : strings.LastIndex(data, ")")]
	var result struct {
		Datas []EastMoneyManager `json:"Datas"`
	}
	err = util.Json.UnmarshalFromString(data, &result)
	if err != nil {
		log.ErrorF("[%s] %s: %s", url, err, data)
		return nil, nil, err
	}
	fnDays := func(d1, d2 string) int {
		t1, err := time.Parse(DATEFORMAT, d1)
		if err != nil {
			return 0
		}
		t2, err := time.Parse(DATEFORMAT, d2)
		if err != nil {
			t2 = time.Now()
		}
		return int(t2.Sub(t1).Hours() / 24)
	}
	type Range struct {
		Start string
		End   string
	}
	fnMergeRange := func(ranges []Range) []Range {
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].Start < ranges[j].Start
		})
		fnOverlap := func(r1, r2 Range) bool {
			return (r2.Start <= r1.Start && r1.Start <= r2.End) || (r1.Start <= r2.Start && r2.Start <= r1.End)
		}
		fnMerge := func(r1, r2 Range) Range {
			s := r1.Start
			if s > r2.Start {
				s = r2.Start
			}
			e := r1.End
			if e < r2.End {
				e = r2.End
			}
			return Range{s, e}
		}
		var merges []Range
		for i := 0; i < len(ranges); i++ {
			if i == 0 {
				merges = append(merges, ranges[i])
			} else {
				curr := ranges[i]
				prev := merges[len(merges)-1]
				if fnOverlap(curr, prev) {
					merges[len(merges)-1] = fnMerge(curr, prev)
				} else {
					merges = append(merges, curr)
				}
			}
		}
		return merges
	}
	for _, m := range result.Datas {
		startWorkDate := time.Now().Format(DATEFORMAT)
		days := 0
		maxGrowth := float64(-100)
		minGrowth := float64(100)
		totalGrowth := float64(0)
		holdings := 0
		var ranges []Range
		for _, l := range m.FundList {
			experiences = append(experiences, ManagerExperience{
				ManagerID:   m.ID,
				ManagerName: m.Name,
				FromDate:    l.FromDate,
				ToDate:      l.ToDate,
				FundCode:    l.FundCode,
				FundName:    l.FundName,
				Growth:      l.Growth,
			})
			if startWorkDate > l.FromDate {
				startWorkDate = l.FromDate
			}
			if maxGrowth < l.Growth {
				maxGrowth = l.Growth
			}
			if minGrowth > l.Growth {
				minGrowth = l.Growth
			}
			totalGrowth += l.Growth
			if l.ToDate == "" {
				holdings++
			}
			r := Range{l.FromDate, l.ToDate}
			if r.End == "" {
				r.End = time.Now().Format(DATEFORMAT)
			}
			ranges = append(ranges, r)
		}
		merges := fnMergeRange(ranges)
		for _, m := range merges {
			days += fnDays(m.Start, m.End)
		}
		managers = append(managers, Manager{
			ID:            m.ID,
			Name:          m.Name,
			StartWorkDate: startWorkDate,
			WorkDays:      days,
			MaxGrowth:     maxGrowth,
			MinGrowth:     minGrowth,
			AveGrowth:     totalGrowth / float64(len(m.FundList)),
			HoldingFunds:  holdings,
			Education:     m.Education,
			Resume:        m.Resume,
		})
	}
	return managers, experiences, nil
}

func (api *EastMoneyApi) GetFundHoldingStock(fundCode string, year int) ([]FundHoldingStock, error) {
	url := fmt.Sprintf(api.urlHoldingStock, fundCode, 20, year, api.GetRandom())
	resp, err := common.Client.Get(url, map[string]string{"Referer": api.referer})
	if err != nil {
		log.ErrorF("%s: %s", err, url)
		return nil, err
	}

	body := string(resp.Body())
	reg := regexp.MustCompile(`"(.*)",arryear:(\[.*\]),curyear:(\d+)`)
	results := reg.FindAllStringSubmatch(body, -1)
	htmls := ""
	var years []int
	curYear := year
	if len(results) == 1 {
		matches := results[0]
		if len(matches) == 4 {
			htmls = matches[1]
			util.Json.UnmarshalFromString(matches[2], &years)
			curYear, _ = strconv.Atoi(matches[3])
		}
	}
	_ = years
	_ = curYear

	doc, err := htmlquery.Parse(strings.NewReader(htmls))
	if err != nil {
		log.ErrorF("[%s] %s: %s", url, err, htmls)
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	fnGetNodeValue := func(top *html.Node, expr string, index int) string {
		nodes := htmlquery.Find(top, expr)
		if index >= 0 && index < len(nodes) {
			return nodes[index].Data
		}
		return ""
	}

	var data []FundHoldingStock
	nodes := htmlquery.Find(doc, "//div[@class='boxitem w790']")
	for _, node := range nodes {
		name := fnGetNodeValue(node, ".//label[@class='left']/a/text()", 0)
		label := string([]rune(strings.TrimSpace(fnGetNodeValue(node, ".//label[@class='left']/text()", 0)))[0:8])
		date := fnGetNodeValue(node, ".//label[@class='right lab2 xq505']/font/text()", 0)
		trs := htmlquery.Find(node, ".//table[contains(@class,'w782 comm tzxq')]/tbody/tr")
		for _, tr := range trs {
			item := FundHoldingStock{
				FundCode:  fundCode,
				FundName:  strings.TrimSpace(name),
				Season:    strings.TrimSpace(label),
				Date:      strings.TrimSpace(date),
				StockCode: strings.TrimSpace(fnGetNodeValue(tr, "./td[2]/a/text()", 0)), // 股票代码
				StockName: strings.TrimSpace(fnGetNodeValue(tr, "./td[3]/a/text()", 0)), // 股票名称
			}
			s := strings.TrimSpace(fnGetNodeValue(tr, "./td[last()-2]/text()", 0)) // 持仓占净值比例
			item.StockPercent, _ = strconv.ParseFloat(strings.TrimRight(s, "%"), 64)
			s = strings.TrimSpace(fnGetNodeValue(tr, "./td[last()-1]/text()", 0)) // 持仓股数，万股
			item.StockAmount, _ = strconv.ParseFloat(strings.ReplaceAll(s, ",", ""), 64)
			s = strings.TrimSpace(fnGetNodeValue(tr, "./td[last()]/text()", 0)) // 持仓市值，万元
			item.StockValue, _ = strconv.ParseFloat(strings.ReplaceAll(s, ",", ""), 64)
			data = append(data, item)
		}
	}
	return data, nil
}

func (api *EastMoneyApi) GetFundNetValue(fundCode string, start, end string, page, pageSize int) ([]FundNetValue, int, error) {
	stamp := time.Now().Unix()
	url := fmt.Sprintf(api.urlNetValue, stamp, fundCode, page, pageSize, start, end, stamp)
	resp, err := common.Client.Get(url, map[string]string{"Referer": api.referer})
	if err != nil {
		log.ErrorF("%s: %s", err, url)
		return nil, 0, err
	}

	data := string(resp.Body())
	data = data[strings.Index(data, "(")+1 : strings.LastIndex(data, ")")]
	var result EastMoneyFundNetValueList
	err = util.Json.UnmarshalFromString(data, &result)
	if err != nil {
		log.ErrorF("[%s] %s: %s", url, err, data)
		return nil, 0, err
	}
	var values []FundNetValue
	for _, r := range result.Datas.List {
		value := FundNetValue{
			Code: fundCode,
			Date: r.Date,
		}
		value.NetValue, _ = strconv.ParseFloat(r.NetValue, 64)
		value.TotalNetValue, _ = strconv.ParseFloat(r.TotalNetValue, 64)
		value.Growth, _ = strconv.ParseFloat(r.Growth, 64)
		values = append(values, value)
	}

	nextPage := page + 1
	if len(values) == 0 {
		nextPage = 0
	}
	return values, nextPage, nil
}

func (api *EastMoneyApi) GetFundEstimate(fundCode string) (FundEstimate, error) {
	var estimate FundEstimate
	url := fmt.Sprintf(api.urlEstimate, fundCode, api.GetRandom())
	resp, err := common.Client.Get(url, map[string]string{"Referer": api.referer})
	if err != nil {
		log.ErrorF("%s: %s", err, url)
		return estimate, err
	}

	data := string(resp.Body())
	data = data[strings.Index(data, "(")+1 : strings.LastIndex(data, ")")]
	var result EastMoneyFundEstimate
	err = util.Json.UnmarshalFromString(data, &result)
	if err != nil {
		log.ErrorF("[%s] %s: %s", url, err, data)
		return estimate, err
	}

	estimate.Code = result.Code
	estimate.Name = result.Name
	estimate.NetDate = result.NetDate
	estimate.NetValue, _ = strconv.ParseFloat(result.NetValue, 64)
	estimate.EstimateNet, _ = strconv.ParseFloat(result.EstimateNet, 64)
	estimate.EstimateRate, _ = strconv.ParseFloat(result.EstimateRate, 64)
	estimate.EstimateDate = result.EstimateDate

	return estimate, nil
}
