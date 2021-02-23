package model

import (
	_ "github.com/mattn/go-sqlite3"
)

// 基本信息
type FundBasic struct {
	Code       string  `db:"code"`        // 代码
	Name       string  `db:"name"`        // 简称
	CreateDate string  `db:"create_date"` // 成立日期
	Scale      float32 `db:"scale"`       // 规模
	Type       uint    `db:"type"`        // 类型
	IsFav      bool    `db:"is_fav"`      // 是否收藏
	SortId     int     `db:"sort_id"`     // 排序id
	Remark     string  `db:"remark"`      // 备注
	Tags       string  `db:"tags"`        // 标签
	UpdateDate string  `db:"update_date"` // 更新日期
}

// 基金经理
type FundManager struct {
	FundCode        string  `db:"fund_code"`
	FromDate        string  `db:"from_date"`
	ToDate          string  `db:"to_date"`
	ManagerName     string  `db:"manager_name"`
	ManagerCompany  string  `db:"manager_company"`
	ManagerWorkDate string  `db:"manager_work_date"`
	Roi             float32 `db:"roi"`
}

// 基金持仓
type FundHoldingStock struct {
	FundCode        string `db:"fund_code"`        // 基金代码
	FundName        string `db:"fund_name"`        // 基金名称
	Season          string `db:"season"`           // 季度
	Date            string `db:"date"`             // 日期
	StockCode       string `db:"stock_code"`       // 股票代码
	StockName       string `db:"stock_name"`       // 股票名称
	StockProportion string `db:"stock_proportion"` // 持仓占净值比例
	StockAmount     string `db:"stock_amount"`     // 持仓股数，万股
	StockValue      string `db:"stock_value"`      // 持仓市值，万元
}

// 基金净值
type FundNetValue struct {
	Code          string  `db:"code"`
	Date          string  `db:"date"`
	NetValue      float32 `db:"net_value"`
	TotalNetValue float32 `db:"total_net_value"`
}

// 基金估值
type FundEstimate struct {
	Code         string `json:"fundcode"` // 基金代码 	519983
	Name         string `json:"name"`     // 基金简称 	长信量化先锋混合A
	NetDate      string `json:"jzrq"`     // 日期 		2018-09-21
	NetValue     string `json:"dwjz"`     // 单位净值 	1.2440
	EstimateNet  string `json:"gsz"`      // 估算净值 	1.2388
	EstimateRate string `json:"gszzl"`    // 估算增长率	-0.42
	EstimateDate string `json:"gztime"`   // 估值日期 	2018-09-25 15:00
}

func SaveFundBasic(fund ...FundBasic) error {
	return nil
}

func SaveFundManager(manager ...FundManager) error {
	return nil
}

func SaveFundNetValue(nets ...FundNetValue) error {
	return nil
}

func SaveFundStockHoldings(holdings ...FundHoldingStock) error {
	return nil
}
