package model

import (
	"fund/common"
	"fund/lang"
	"github.com/lauthrul/goutil/log"
	"github.com/lauthrul/goutil/util"
	"regexp"
	"strconv"
	"strings"
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

type EastMoneyFundEstimate struct {
	Code         string  `json:"fundcode"` // 基金代码 	519983
	Name         string  `json:"name"`     // 基金简称 	长信量化先锋混合A
	NetDate      string  `json:"jzrq"`     // 日期 		2018-09-21
	NetValue     float32 `json:"dwjz"`     // 单位净值 	1.2440
	EstimateNet  string  `json:"gsz"`      // 估算净值 	1.2388
	LastDayNet   string  `json:"gszzl"`    // 估算增长率	-0.42
	EstimateDate string  `json:"gztime"`   // 估值日期 	2018-09-25 15:00
}

type EastMoneyApi struct {
	urlFundSearch string
	urlRank       string
	urlEstimate   string
	referer       string
}

func NewEastMoneyApi() *EastMoneyApi {
	return &EastMoneyApi{
		urlFundSearch: "http://fund.eastmoney.com/js/fundcode_search.js?v=%s", // http://fund.eastmoney.com/js/fundcode_search.js?v=20210204133444
		urlRank:       "http://fund.eastmoney.com/data/rankhandler.aspx",      // http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=all&rs=&gs=0&sc=rzdf&st=desc&sd=2020-02-18&ed=2021-02-18&qdii=&tabSubtype=,,,,,&pi=1&pn=10&dx=1&v=0.29261668485886694
		urlEstimate:   "http://fundgz.1234567.com.cn/js/%s.js?rt=%d",          // http://fundgz.1234567.com.cn/js/007047.js?rt=1612108800
		referer:       "http://fundf10.eastmoney.com/",
	}
}

func (api *EastMoneyApi) GetTHReference() []THReference {
	return EastMoneyFund{}.GetTHReference()
}

func (api *EastMoneyApi) GetFundRank(arg FundRankArg) (FundList, error) {
	var funds FundList
	url := api.urlRank + "?" + arg.String()
	resp, err := common.Client.Get(url, map[string]string{"Referer": api.referer})
	if err != nil {
		log.Error(err)
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
				log.Error(err)
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

func (api *EastMoneyApi) GetFundEstimate(fundCode string) (EastMoneyFundEstimate, error) {
	var estimate EastMoneyFundEstimate
	return estimate, nil
}
