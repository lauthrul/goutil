package main

import (
	"fmt"
	"github.com/lauthrul/goutil/util"
	"github.com/shopspring/decimal"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Index struct {
	Code          string          `name:"-"`
	Name          string          `name:"名称"`      // 0："上证指数", 指数名称
	Last          decimal.Decimal `name:"当前点数"`    // 1："3565.9046", 当前点数
	Change        decimal.Decimal `name:"涨跌"`      // 2："-32.7472", 当前价格
	ChangePercent decimal.Decimal `name:"涨跌率"`     // 3："-0.91", 涨跌率
	Volume        decimal.Decimal `name:"成交量（手）"`  // 4："3476668", 成交量（手）
	Amount        decimal.Decimal `name:"成交额（万元）"` // 5："51132510", 成交额（万元）
}

func (i Index) String() string {
	return fmt.Sprintf("【%s】 指数：%s, 涨幅：(%s, %s%%), 成交量（万手）：%s, 成交额（亿元）：%s",
		i.Name,
		i.Last.StringFixed(3),
		i.Change.StringFixed(3),
		i.ChangePercent.StringFixed(2),
		i.Volume.Div(decimal.NewFromInt(1000000)).StringFixed(3),
		i.Amount.Div(decimal.NewFromInt(10000)).StringFixed(3),
	)
}

func (i *Index) Parse(str string) error {
	arr := strings.Split(str, ",")
	if len(arr) < 6 {
		return fmt.Errorf("invalid index:%s", str)
	}

	i.Name = arr[0]
	i.Last = util.Str2Decimal(arr[1])
	i.Change = util.Str2Decimal(arr[2])
	i.ChangePercent = util.Str2Decimal(arr[3])
	i.Volume = util.Str2Decimal(arr[4])
	i.Amount = util.Str2Decimal(arr[5])

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Stock struct {
	Code           string          `index:"" name:"-"`
	Name           string          `index:"1" name:"名称"`   // 0：”大秦铁路”，股票名字；
	TodayOpen      decimal.Decimal `index:"6" name:"今开"`   // 1：”27.55″，今日开盘价；
	YesterdayClose decimal.Decimal `index:"5" name:"昨收"`   // 2：”27.25″，昨日收盘价；
	Current        decimal.Decimal `index:"2" name:"当前"`   // 3：”26.91″，当前价格；
	TodayHighest   decimal.Decimal `index:"8" name:"最高"`   // 4：”27.55″，今日最高价；
	TodayLowest    decimal.Decimal `index:"7" name:"最低"`   // 5：”26.20″，今日最低价；
	CurrentBuy     decimal.Decimal `index:"" name:"-"`     // 6：”26.91″，竞买价，即“买一”报价；
	CurrentSell    decimal.Decimal `index:"" name:"-"`     // 7：”26.92″，竞卖价，即“卖一”报价；
	Volume         decimal.Decimal `index:"9" name:"成交量"`  // 8：”22114263″，成交的股票数，由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百；
	Amount         decimal.Decimal `index:"10" name:"成交额"` // 9：”589824680″，成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万；
	Buy1Amount     decimal.Decimal `index:"" name:"-"`     // 10：”4695″，“买一”申请4695股，即47手；
	Buy1           decimal.Decimal `index:"" name:"-"`     // 11：”26.91″，“买一”报价；
	Buy2Amount     decimal.Decimal `index:"" name:"-"`     // 12：”57590″，“买二”
	Buy2           decimal.Decimal `index:"" name:"-"`     // 13：”26.90″，“买二”
	Buy3Amount     decimal.Decimal `index:"" name:"-"`     // 14：”14700″，“买三”
	Buy3           decimal.Decimal `index:"" name:"-"`     // 15：”26.89″，“买三”
	Buy4Amount     decimal.Decimal `index:"" name:"-"`     // 16：”14300″，“买四”
	Buy4           decimal.Decimal `index:"" name:"-"`     // 17：”26.88″，“买四”
	Buy5Amount     decimal.Decimal `index:"" name:"-"`     // 18：”15100″，“买五”
	Buy5           decimal.Decimal `index:"" name:"-"`     // 19：”26.87″，“买五”
	Sell1Amount    decimal.Decimal `index:"" name:"-"`     // 20：”3100″，“卖一”申报3100股，即31手；
	Sell1          decimal.Decimal `index:"" name:"-"`     // 21：”26.92″，“卖一”报价
	Sell2Amount    decimal.Decimal `index:"" name:"-"`     // (22, 23), (24, 25), (26,27), (28, 29)分别为“卖二”至“卖五的情况”
	Sell2          decimal.Decimal `index:"" name:"-"`
	Sell3Amount    decimal.Decimal `index:"" name:"-"`
	Sell3          decimal.Decimal `index:"" name:"-"`
	Sell4Amount    decimal.Decimal `index:"" name:"-"`
	Sell4          decimal.Decimal `index:"" name:"-"`
	Sell5Amount    decimal.Decimal `index:"" name:"-"`
	Sell5          decimal.Decimal `index:"" name:"-"`
	Date           string          `index:"" name:"-"` // 30：”2008-01-11″，日期
	Time           string          `index:"" name:"-"` // 31：”15:05:32″，时间
	Change         decimal.Decimal `index:"3" name:"涨跌"`
	ChangePercent  decimal.Decimal `index:"4" name:"涨跌率" format:"%s%%"`
}

func (s Stock) String() string {
	return fmt.Sprintf("【%s】 当前：%s, 涨幅：(%s, %s%%), 昨收：%s, 今开：%s, 最高：%s, 最低：%s, 成交量（万手）：%s, 成交额（万元）：%s",
		s.Name,
		s.Current.StringFixed(3),
		s.Change.StringFixed(3),
		s.ChangePercent.StringFixed(2),
		s.YesterdayClose.StringFixed(3),
		s.TodayOpen.StringFixed(3),
		s.TodayHighest.StringFixed(3),
		s.TodayLowest.StringFixed(3),
		s.Volume.Div(decimal.NewFromInt(10000)).StringFixed(3),
		s.Amount.Div(decimal.NewFromInt(10000)).StringFixed(3))
}

func (s *Stock) Parse(str string) error {
	arr := strings.Split(str, ",")
	if len(arr) < 32 {
		return fmt.Errorf("invalid stock:%s", s)
	}

	s.Name = arr[0]
	s.TodayOpen = util.Str2Decimal(arr[1])
	s.YesterdayClose = util.Str2Decimal(arr[2])
	s.Current = util.Str2Decimal(arr[3])
	s.TodayHighest = util.Str2Decimal(arr[4])
	s.TodayLowest = util.Str2Decimal(arr[5])
	s.CurrentBuy = util.Str2Decimal(arr[6])
	s.CurrentSell = util.Str2Decimal(arr[7])
	s.Volume = util.Str2Decimal(arr[8])
	s.Amount = util.Str2Decimal(arr[9])
	s.Buy1Amount = util.Str2Decimal(arr[10])
	s.Buy1 = util.Str2Decimal(arr[11])
	s.Buy2Amount = util.Str2Decimal(arr[12])
	s.Buy2 = util.Str2Decimal(arr[13])
	s.Buy3Amount = util.Str2Decimal(arr[14])
	s.Buy3 = util.Str2Decimal(arr[15])
	s.Buy4Amount = util.Str2Decimal(arr[16])
	s.Buy4 = util.Str2Decimal(arr[17])
	s.Buy5Amount = util.Str2Decimal(arr[18])
	s.Buy5 = util.Str2Decimal(arr[19])
	s.Sell1Amount = util.Str2Decimal(arr[20])
	s.Sell1 = util.Str2Decimal(arr[21])
	s.Sell2Amount = util.Str2Decimal(arr[22])
	s.Sell2 = util.Str2Decimal(arr[23])
	s.Sell3Amount = util.Str2Decimal(arr[24])
	s.Sell3 = util.Str2Decimal(arr[25])
	s.Sell4Amount = util.Str2Decimal(arr[26])
	s.Sell4 = util.Str2Decimal(arr[27])
	s.Sell5Amount = util.Str2Decimal(arr[28])
	s.Sell5 = util.Str2Decimal(arr[29])
	s.Date = arr[30]
	s.Time = arr[31]
	s.Change = s.Current.Sub(s.YesterdayClose)
	s.ChangePercent = s.Change.Div(s.YesterdayClose).Mul(decimal.NewFromInt(100)).Round(2)

	return nil
}
