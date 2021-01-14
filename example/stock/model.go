package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Index struct {
	Name          string          // 0："上证指数", 指数名称
	Last          decimal.Decimal // 1："3565.9046", 当前点数
	Change        decimal.Decimal // 2："-32.7472", 当前价格
	ChangePercent decimal.Decimal // 3："-0.91", 涨跌率
	Volume        decimal.Decimal // 4："3476668", 成交量（手）
	Amount        decimal.Decimal // 5："51132510", 成交额（万元）
}

func (i Index) String() string {
	return fmt.Sprintf("【%s】 指数：%s, 涨幅：(%s, %s%%), 成交量（万手）：%s, 成交额（亿元）：%s",
		i.Name, i.Last, i.Change, i.ChangePercent,
		i.Volume.Div(decimal.NewFromInt(1000000)), i.Amount.Div(decimal.NewFromInt(10000)))
}

func (i *Index) Parse(str string) error {
	arr := strings.Split(str, ",")
	if len(arr) >= 6 {
		return fmt.Errorf("invalid index:%s", str)
	}

	i.Name = arr[0]
	i.Last = Str2Decimal(arr[1])
	i.Change = Str2Decimal(arr[2])
	i.ChangePercent = Str2Decimal(arr[3])
	i.Volume = Str2Decimal(arr[4])
	i.Amount = Str2Decimal(arr[5])

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Stock struct {
	Name           string          // 0：”大秦铁路”，股票名字；
	TodayOpen      decimal.Decimal // 1：”27.55″，今日开盘价；
	YesterdayClose decimal.Decimal // 2：”27.25″，昨日收盘价；
	Current        decimal.Decimal // 3：”26.91″，当前价格；
	TodayHighest   decimal.Decimal // 4：”27.55″，今日最高价；
	TodayLowest    decimal.Decimal // 5：”26.20″，今日最低价；
	CurrentBuy     decimal.Decimal // 6：”26.91″，竞买价，即“买一”报价；
	CurrentSell    decimal.Decimal // 7：”26.92″，竞卖价，即“卖一”报价；
	Volume         decimal.Decimal // 8：”22114263″，成交的股票数，由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百；
	Amount         decimal.Decimal // 9：”589824680″，成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万；
	Buy1Amount     decimal.Decimal // 10：”4695″，“买一”申请4695股，即47手；
	Buy1           decimal.Decimal // 11：”26.91″，“买一”报价；
	Buy2Amount     decimal.Decimal // 12：”57590″，“买二”
	Buy2           decimal.Decimal // 13：”26.90″，“买二”
	Buy3Amount     decimal.Decimal // 14：”14700″，“买三”
	Buy3           decimal.Decimal // 15：”26.89″，“买三”
	Buy4Amount     decimal.Decimal // 16：”14300″，“买四”
	Buy4           decimal.Decimal // 17：”26.88″，“买四”
	Buy5Amount     decimal.Decimal // 18：”15100″，“买五”
	Buy5           decimal.Decimal // 19：”26.87″，“买五”
	Sell1Amount    decimal.Decimal // 20：”3100″，“卖一”申报3100股，即31手；
	Sell1          decimal.Decimal // 21：”26.92″，“卖一”报价
	Sell2Amount    decimal.Decimal // (22, 23), (24, 25), (26,27), (28, 29)分别为“卖二”至“卖五的情况”
	Sell2          decimal.Decimal
	Sell3Amount    decimal.Decimal
	Sell3          decimal.Decimal
	Sell4Amount    decimal.Decimal
	Sell4          decimal.Decimal
	Sell5Amount    decimal.Decimal
	Sell5          decimal.Decimal
	Date           string // 30：”2008-01-11″，日期
	Time           string // 31：”15:05:32″，时间
	Change         decimal.Decimal
	ChangePercent  decimal.Decimal
}

func (s Stock) String() string {
	return fmt.Sprintf("【%s】 当前：%s, 涨幅：(%s, %s%%), 昨收：%s, 今开：%s, 最高：%s, 最低：%s, 成交量（万手）：%s, 成交额（万元）：%s",
		s.Name, s.Current, s.Change, s.ChangePercent.StringFixed(2),
		s.YesterdayClose, s.TodayOpen, s.TodayHighest, s.TodayLowest,
		s.Volume.Div(decimal.NewFromInt(10000)), s.Amount.Div(decimal.NewFromInt(10000)))
}

func (s *Stock) Parse(str string) error {
	arr := strings.Split(str, ",")
	if len(arr) < 32 {
		return fmt.Errorf("invalid stock:%s", s)
	}

	s.Name = arr[0]
	s.TodayOpen = Str2Decimal(arr[1])
	s.YesterdayClose = Str2Decimal(arr[2])
	s.Current = Str2Decimal(arr[3])
	s.TodayHighest = Str2Decimal(arr[4])
	s.TodayLowest = Str2Decimal(arr[5])
	s.CurrentBuy = Str2Decimal(arr[6])
	s.CurrentSell = Str2Decimal(arr[7])
	s.Volume = Str2Decimal(arr[8])
	s.Amount = Str2Decimal(arr[9])
	s.Buy1Amount = Str2Decimal(arr[10])
	s.Buy1 = Str2Decimal(arr[11])
	s.Buy2Amount = Str2Decimal(arr[12])
	s.Buy2 = Str2Decimal(arr[13])
	s.Buy3Amount = Str2Decimal(arr[14])
	s.Buy3 = Str2Decimal(arr[15])
	s.Buy4Amount = Str2Decimal(arr[16])
	s.Buy4 = Str2Decimal(arr[17])
	s.Buy5Amount = Str2Decimal(arr[18])
	s.Buy5 = Str2Decimal(arr[19])
	s.Sell1Amount = Str2Decimal(arr[20])
	s.Sell1 = Str2Decimal(arr[21])
	s.Sell2Amount = Str2Decimal(arr[22])
	s.Sell2 = Str2Decimal(arr[23])
	s.Sell3Amount = Str2Decimal(arr[24])
	s.Sell3 = Str2Decimal(arr[25])
	s.Sell4Amount = Str2Decimal(arr[26])
	s.Sell4 = Str2Decimal(arr[27])
	s.Sell5Amount = Str2Decimal(arr[28])
	s.Sell5 = Str2Decimal(arr[29])
	s.Date = arr[30]
	s.Time = arr[31]
	s.Change = s.Current.Sub(s.YesterdayClose)
	s.ChangePercent = s.Change.Div(s.YesterdayClose).Mul(decimal.NewFromInt(100))

	return nil
}
