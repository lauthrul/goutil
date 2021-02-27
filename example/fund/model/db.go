package model

import (
	"database/sql"
	"fmt"
	"fund/config"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
	"github.com/lauthrul/goutil/log"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"time"
)

const (
	DATEFORMAT = "2006-01-02"

	tbFund                 = "fund"
	tbGroup                = "group"
	tbFundGroup            = "fund_group"
	tbHoldingStock         = "holding_stock"
	tbManager              = "manager"
	tbManagerExperience    = "manager_experience"
	tbNetValue             = "net_value"
	viewLatestHoldingStock = "latest_holding_stock"
)

var (
	dialect = goqu.Dialect("sqlite3")
	db      *sqlx.DB
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

func (f FundBasic) Titles() []string {
	return []string{
		"代码",
		"简称",
		"类型",
		"成立日期",
		"成立规模（元）",
		"最新规模（元）",
		"更新日期",
		"基金公司代码",
		"基金公司名称",
		"基金经理ID",
		"基金经理名称",
		"管理费率（年）",
		"托管费率（年）",
		"是否收藏",
		"排序id",
		"备注",
		"标签",
	}
}

func (f FundBasic) Values() []string {
	return []string{
		f.Code,
		f.Name,
		f.Type,
		f.CreateDate[0:10],
		fmt.Sprintf("%.2f", f.CreateScale),
		fmt.Sprintf("%.2f", f.LatestScale),
		f.UpdateDate[0:10],
		f.CompanyCode,
		f.CompanyName,
		f.ManagerID,
		f.ManagerName,
		fmt.Sprintf("%.2f%%", f.ManageExp),
		fmt.Sprintf("%.2f%%", f.TrustExp),
		fmt.Sprintf("%t", f.IsFav),
		fmt.Sprintf("%d", f.SortId),
		f.Remark,
		f.Tags,
	}
}

// 分组名
type Group struct {
	Name string `db:"name"`
}

// 基金分组
type FundGroup struct {
	FundCode string `db:"fund_code"`
	Group    string `db:"group"`
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
	Growth        float64 `db:"growth"`
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

type ListFundArg struct {
	IsFav bool
	Code  []string
	Name  string
}

func GetDB() *sqlx.DB {
	if db == nil {
		var err error
		db, err = sqlx.Open("sqlite3", config.Conf.DbFile)
		if err != nil {
			panic(err)
		}
	}
	return db
}

func SaveFundBasic(tx *sql.Tx, funds ...FundBasic) error {
	for _, v := range funds {
		stat, _, err := dialect.Insert(tbFund).Rows(v).
			OnConflict(goqu.DoUpdate("code", goqu.Record{
				"latest_scale": v.LatestScale,
				"update_date":  v.UpdateDate,
				"manager_id":   v.ManagerID,
				"manager_name": v.ManagerName,
				"manage_exp":   v.ManageExp,
				"trust_exp":    v.TrustExp,
			})).ToSQL()
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
		// log.Debug(stat)
		_, err = tx.Exec(stat)
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
	}
	return nil
}

func SaveManager(tx *sql.Tx, managers ...Manager) error {
	for _, v := range managers {
		stat, _, err := dialect.Insert(tbManager).Rows(v).
			OnConflict(goqu.DoUpdate("id", goqu.Record{
				"work_days":     v.WorkDays,
				"max_growth":    v.MaxGrowth,
				"min_growth":    v.MinGrowth,
				"ave_growth":    v.AveGrowth,
				"holding_funds": v.HoldingFunds,
				"resume":        v.Resume,
			})).ToSQL()
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
		// log.Debug(stat)
		_, err = tx.Exec(stat)
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
	}
	return nil
}

func SaveManagerExperience(tx *sql.Tx, experiences ...ManagerExperience) error {
	for _, v := range experiences {
		stat, _, err := dialect.Insert(tbManagerExperience).Rows(v).OnConflict(goqu.DoNothing()).ToSQL()
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
		// log.Debug(stat)
		_, err = tx.Exec(stat)
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
	}
	return nil
}

func SaveFundNetValue(tx *sql.Tx, nets ...FundNetValue) error {
	for _, v := range nets {
		stat, _, err := dialect.Insert(tbNetValue).Rows(v).OnConflict(goqu.DoNothing()).ToSQL()
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
		// log.Debug(stat)
		_, err = tx.Exec(stat)
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
	}
	return nil
}

func SaveFundStockHoldings(tx *sql.Tx, holdings ...FundHoldingStock) error {
	for _, v := range holdings {
		stat, _, err := dialect.Insert(tbHoldingStock).Rows(v).
			OnConflict(goqu.DoUpdate("fund_code, season, stock_code", goqu.Record{
				"fund_code":     v.FundCode,
				"fund_name":     v.FundName,
				"stock_percent": v.StockPercent,
				"stock_amount":  v.StockAmount,
				"stock_value":   v.StockValue,
			})).ToSQL()
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
		// log.Debug(stat)
		_, err = tx.Exec(stat)
		if err != nil {
			log.ErrorF("%q: %s", err, stat)
			return err
		}
	}
	return nil
}

func GetNextFundNetValueDate(fundCode string) (string, error) {
	stat, _, err := dialect.Select(goqu.MAX("date")).From(tbNetValue).Where(goqu.Ex{"code": fundCode}).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return "", err
	}
	var date string
	err = GetDB().Get(&date, stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return "", err
	}
	t, err := time.Parse(DATEFORMAT, date)
	if err != nil {
		log.ErrorF("%q: %s", err, date)
		return "", err
	}
	t = t.AddDate(0, 0, 1)
	date = t.Format(DATEFORMAT)
	return date, err
}

func GetLatestHoldingStock(fundCode ...string) ([][]FundHoldingStock, error) {
	s := ""
	for _, f := range fundCode {
		s += fmt.Sprintf(`"%s",`, f)
	}
	s = strings.TrimRight(s, ",")
	stat := fmt.Sprintf(`select t1.* from %s t1, %s t2 
where t1.fund_code = t2.fund_code and t1.date = t2.date and t1.fund_code in (%s)`, tbHoldingStock, viewLatestHoldingStock, s)
	var items []FundHoldingStock
	err := GetDB().Select(&items, stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return nil, err
	}
	data := map[string][]FundHoldingStock{}
	for _, item := range items {
		data[item.FundCode] = append(data[item.FundCode], item)
	}
	var result [][]FundHoldingStock
	for _, v := range data {
		result = append(result, v)
	}
	return result, nil
}

func SetFundFav(flag bool, fundCode ...string) error {
	stat, _, err := dialect.Update(tbFund).Set(goqu.Record{"is_fav": flag}).Where(goqu.Ex{"code": fundCode}).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return err
	}
	_, err = GetDB().Exec(stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return err
}

func AddGroup(group ...string) error {
	var groups []Group
	for _, g := range group {
		groups = append(groups, Group{
			Name: g,
		})
	}
	stat, _, err := dialect.Insert(tbGroup).Rows(groups).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return err
	}
	_, err = GetDB().Exec(stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return err
}

func RemoveGroup(group ...string) error {
	stat, _, err := dialect.Delete(tbGroup).Where(goqu.Ex{"name": group}).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return err
	}
	_, err = GetDB().Exec(stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return err
}

func ListGroup() ([]string, error) {
	var list []string
	stat, _, err := dialect.From(tbGroup).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return list, err
	}
	err = GetDB().Select(&list, stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return list, err
}

func AddFundGroup(group string, fundCode ...string) error {
	var records []FundGroup
	for _, f := range fundCode {
		records = append(records, FundGroup{
			FundCode: f,
			Group:    group,
		})
	}
	stat, _, err := dialect.Insert(tbFundGroup).Rows(records).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return err
	}
	_, err = GetDB().Exec(stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return err
}

func RemoveFundGroup(group string, fundCode ...string) error {
	stat, _, err := dialect.Delete(tbFundGroup).Where(goqu.Ex{"group": group, "fund_code": fundCode}).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return err
	}
	_, err = GetDB().Exec(stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return err
}

func SetFundRemark(remark string, fundCode ...string) error {
	stat, _, err := dialect.Update(tbFund).Set(goqu.Record{"remark": remark}).Where(goqu.Ex{"code": fundCode}).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return err
	}
	_, err = GetDB().Exec(stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return err
}

func ListFund(arg ListFundArg) ([]FundBasic, error) {
	var exs []exp.Expression
	if arg.IsFav {
		exs = append(exs, goqu.Ex{"is_fav": arg.IsFav})
	}
	if len(arg.Code) > 0 {
		exs = append(exs, goqu.Ex{"code": arg.Code})
	}
	if len(arg.Name) > 0 {
		exs = append(exs, goqu.C("name").Like(fmt.Sprintf("%%%s%%", arg.Name)))
	}
	var list []FundBasic
	stat, _, err := dialect.From(tbFund).Where(exs...).ToSQL()
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
		return list, err
	}
	err = GetDB().Select(&list, stat)
	if err != nil {
		log.ErrorF("%q: %s", err, stat)
	}
	return list, err
}
