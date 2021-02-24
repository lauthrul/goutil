package model

import (
	"database/sql"
	"fund/config"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/lauthrul/goutil/log"
	_ "github.com/mattn/go-sqlite3"
)

const (
	tbBasic             = "basic"
	tbHoldingStock      = "holding_stock"
	tbManager           = "manager"
	tbManagerExperience = "manager_experience"
	tbNetValue          = "net_value"
)

var (
	dialect = goqu.Dialect("sqlite")
	db      *sql.DB
)

// 基本信息
type FundBasic struct {
	Code        string  `db:"code"`         // 代码
	Name        string  `db:"name"`         // 简称
	Type        string  `db:"type"`         // 类型
	CreateDate  string  `db:"create_date"`  // 成立日期
	CreateScale float64 `db:"create_scale"` // 成立规模（元）
	LatestScale float64 `db:"latest_scale"` // 最新规模（元）
	UpdateDate  string  `db:"update_date"`  // 更新日期
	CompanyCode string  `db:"company_code"` // 基金公司代码
	CompanyName string  `db:"company_name"` // 基金公司名称
	ManagerID   string  `db:"manager_id"`   // 基金经理ID
	ManagerName string  `db:"manager_name"` // 基金经理名称
	ManageExp   float64 `db:"manage_exp"`   // 管理费率（年）
	TrustExp    float64 `db:"trust_exp"`    // 托管费率（年）
	IsFav       bool    `db:"is_fav"`       // 是否收藏
	SortId      int     `db:"sort_id"`      // 排序id
	Remark      string  `db:"remark"`       // 备注
	Tags        string  `db:"tags"`         // 标签
}

// 基金经理
type Manager struct {
	ID            string  `db:"id"`
	Name          string  `db:"name"`
	StartWorkDate string  `db:"start_work_date"`
	WorkDays      int     `db:"work_days"`
	MaxGrowth     float64 `db:"max_growth"`
	MinGrowth     float64 `db:"min_growth"`
	AveGrowth     float64 `db:"ave_growth"`
	HoldingFunds  int     `db:"holding_funds"`
	Education     string  `db:"education"`
	Resume        string  `db:"resume"`
}

// 基金经理从业经历
type ManagerExperience struct {
	ManagerID   string  `db:"manager_id"`
	ManagerName string  `db:"manager_name"`
	FromDate    string  `db:"from_date"`
	ToDate      string  `db:"to_date"`
	FundCode    string  `db:"fund_code"`
	FundName    string  `db:"fund_name"`
	Growth      float64 `db:"growth"`
}

// 基金持仓
type FundHoldingStock struct {
	FundCode     string  `db:"fund_code"`     // 基金代码
	FundName     string  `db:"fund_name"`     // 基金名称
	Season       string  `db:"season"`        // 季度
	Date         string  `db:"date"`          // 日期
	StockCode    string  `db:"stock_code"`    // 股票代码
	StockName    string  `db:"stock_name"`    // 股票名称
	StockPercent float64 `db:"stock_percent"` // 持仓占比
	StockAmount  float64 `db:"stock_amount"`  // 持仓股数，万股
	StockValue   float64 `db:"stock_value"`   // 持仓市值，万元
}

// 基金净值
type FundNetValue struct {
	Code          string  `db:"code"`
	Date          string  `db:"date"`
	NetValue      float64 `db:"net_value"`
	TotalNetValue float64 `db:"total_net_value"`
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

func GetDB() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", config.Conf.DbFile)
		if err != nil {
			panic(err)
		}
	}
	return db
}

func SaveFundBasic(fund ...FundBasic) error {
	stat, _, _ := dialect.Insert(tbBasic).Rows(fund).ToSQL()
	_, err := GetDB().Exec(stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return err
}

func SaveManager(manager ...Manager) error {
	return nil
}

func SaveManagerExperience(experiences ...ManagerExperience) error {
	return nil
}

func SaveFundNetValue(nets ...FundNetValue) error {
	return nil
}

func SaveFundStockHoldings(holdings ...FundHoldingStock) error {
	return nil
}
