package model

import (
	"fmt"
	"time"
)

type THMeta struct {
	Text        string
	EnableOrder bool
	OrderFiled  string
}

type Fund interface {
	GetMetas() []THMeta
	GetValues() []string
}

// 基金列表
type FundList struct {
	List       []Fund
	PageIndex  int
	PageSize   int
	TotalPage  int
	TotalCount int
}

// 基金排行请求参数
type FundRankArg struct {
	FundType    string `name:"ft"` // all-所有，gp-股票型，hh-混合型，zq-债券型，zs-指数型，qdii-QDII型，lof-LOF型，fof-FOF型
	FundCompany string `name:"gs"` // 基金公司代码
	SortCode    string `name:"sc"` // dm-基金代码，jc-基金简称，dwjz-单位净值，ljjz-累计净值，rzdf-日增长率，zzf-近1周， 1yzf-近1月，3yzf-近3月，6yzf-近6月，1nzf-近1年，2nzf-近2年，3nzf-近3年，jnzf-今年以来，lnzf-成立以来，qjzf-自定义（查询时间段）
	SortType    string `name:"st"` // asc-升序 desc-降序
	StartDate   string `name:"sd"` // 开始日期 2020-02-18
	EndDate     string `name:"ed"` // 结束日期 2021-02-18
	PageIndex   int    `name:"pi"` // 页码
	PageNumber  int    `name:"pn"` // 页大小
}

func (f *FundRankArg) String() string {
	return fmt.Sprintf(
		"op=ph"+
			"&dt=kf"+
			"&ft=%s"+
			"&rs="+
			"&gs=%s"+
			"&sc=%s"+
			"&st=%s"+
			"&sd=%s"+
			"&ed=%s"+
			"&qdii="+
			"&tabSubtype=,,,,,"+
			"&pi=%d"+
			"&pn=%d"+
			"&dx=1"+
			"&v=%s",
		f.FundType,
		f.FundCompany,
		f.SortCode,
		f.SortType,
		f.StartDate,
		f.EndDate,
		f.PageIndex,
		f.PageNumber,
		fmt.Sprintf("0.%d", time.Now().Unix()),
	)
}

type FundFavArg struct {
	Group     string // 分组
	IsFav     int    // -1-未指定 0-否 1-是
	SortCode  string // 排序字段
	SortType  string // asc-升序 desc-降序
	PageIndex int    // 页码
	PageSize  int    // 页大小
}

// api接口
type Api interface {
	GetFundRankMeta() []THMeta
	GetFundFavMeta() []THMeta
	GetFundRank(arg FundRankArg) (FundList, error)
	GetFundFav(arg FundFavArg) (FundList, error)
	GetFundBasic(fundCode string) (FundBasic, error)
	GetFundManager(fundCode string) ([]Manager, []ManagerExperience, error)
	GetFundHoldingStock(fundCode string, year int) ([]FundHoldingStock, error)
	GetFundNetValue(fundCode string, start, end string, page, pageSize int) ([]FundNetValue, int, error)
	GetFundEstimate(fundCode string) (FundEstimate, error)
}
