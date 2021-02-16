package model

import (
	"fmt"
	"fund/common"
	"fund/lang"
	"github.com/lauthrul/goutil/log"
	"github.com/lauthrul/goutil/util"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type Fund struct {
	Id                      string `json:"id"`
	FundCode                string `json:"fund_code" name:"代码"`
	FundName                string `json:"fund_name" name:"基金"`
	CompanyId               string `json:"company_id"`
	FundShortName           string `json:"fund_short_name"`
	FundLegalName           string `json:"fund_legal_name"`
	FundPinyin              string `json:"fund_pinyin"`
	FundRiskLevel           string `json:"fund_risk_level"`
	ShareType               string `json:"share_type"`
	FundType                string `json:"fund_type"`
	Type1Id                 string `json:"type1_id"`
	Type2Id                 string `json:"type2_id"`
	Type3Id                 string `json:"type3_id"`
	HelpType4Id             string `json:"help_type4_id"`
	IsMoneyFund             string `json:"is_money_fund"`
	HfIncomeratio           string `json:"hf_incomeratio"`
	Income                  string `json:"income"`
	Incomeratio             string `json:"incomeratio"`
	IncomeratioT001         string `json:"incomeratio_t001"`
	IncomeratioS010         string `json:"incomeratio_s010"`
	Netvalue                string `json:"netvalue" name:"净值"`
	TotalNetvalue           string `json:"total_netvalue" name:"累计净值"`
	Navdate                 string `json:"navdate" name:"更新日期"`
	FundState               string `json:"fund_state"`
	MinShare                string `json:"min_share"`
	DayIncratio             string `json:"day_incratio" name:"日涨幅"`
	MonthIncratio           string `json:"month_incratio" name:"月涨幅"`
	QuarterIncratio         string `json:"quarter_incratio" name:"近3月涨幅"`
	YearIncratio            string `json:"year_incratio" name:"近1年涨幅"`
	ThisYearIncratio        string `json:"this_year_incratio" name:"今年以来涨幅"`
	HalfYearIncratio        string `json:"half_year_incratio" name:"近半年涨幅"`
	ThreeYearIncratio       string `json:"three_year_incratio" name:"近3年涨幅"`
	Feerate1                string `json:"feerate1"`
	SinaFeerate1            string `json:"sina_feerate1"`
	Feerate2                string `json:"feerate2"`
	SinaFeerate2            string `json:"sina_feerate2"`
	Discountrate2           string `json:"discountrate2"`
	Feerate3                string `json:"feerate3"`
	SinaFeerate3            string `json:"sina_feerate3"`
	Dayinc                  string `json:"dayinc"`
	SubscribeState          string `json:"subscribe_state"`
	DeclareState            string `json:"declare_state"`
	WithdrawState           string `json:"withdraw_state"`
	ValuagrState            string `json:"valuagr_state"`
	TrendState              string `json:"trend_state"`
	HopedeclareState        string `json:"hopedeclare_state"`
	DeclareendDay           string `json:"declareend_day"`
	RedeemendDay            string `json:"redeemend_day"`
	OpenDay                 string `json:"open_day"`
	IsForbidbuybyredeem     string `json:"is_forbidbuybyredeem"`
	TransFlag               string `json:"trans_flag"`
	FundsubType             string `json:"fundsub_type"`
	SubscribeBeigin         string `json:"subscribe_beigin"`
	SinaSubscribeBeigin     string `json:"sina_subscribe_beigin"`
	HelpSinaSubscribeBeigin string `json:"help_sina_subscribe_beigin"`
	SinaSubscribeEnd        string `json:"sina_subscribe_end"`
	SubscribeEnd            string `json:"subscribe_end"`
	Zjzfe                   string `json:"zjzfe"`
	FundScale               string `json:"fund_scale" name:"规模"`
	FundManager             []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"fund_manager"`
	UTime             string `json:"u_time"`
	CTime             string `json:"c_time" name:"成立时间"`
	ToThisDayIncratio string `json:"to_this_day_incratio" name:"成立以来涨幅"`
	LDay              string `json:"l_day"`
	BonusDate         string `json:"bonus_date"`
	Forbidbonustype   string `json:"forbidbonustype"`
	Status            string `json:"status"`
	QdiiConfirmDay    string `json:"qdii_confirm_day"`
	TaCode            string `json:"ta_code"`
	ManageRatio       string `json:"manage_ratio"`
	FarePayRate       string `json:"fare_pay_rate"`
	MinValue1         string `json:"min_value1"`
	MinValue2         string `json:"min_value2"`
	OpenStatus        string `json:"open_status"`
	CanBuy            struct { //在售状态
		Show   string `json:"show"`
		Canbuy int    `json:"canbuy"`
	} `json:"can_buy"`
	Type2IdDesc string `json:"type2_id_desc" name:"类型"`
	JzAndSy     string `json:"jz_and_sy"`
	MinValue    string `json:"min_value"` // 起购金额
}

func (f Fund) GetTitles() []string {
	return []string{
		lang.Text(common.Lan, "thCode"),
		lang.Text(common.Lan, "thName"),
		lang.Text(common.Lan, "thType"),
		lang.Text(common.Lan, "thScale"),
		lang.Text(common.Lan, "thFoundTime"),
		lang.Text(common.Lan, "thManager"),
		lang.Text(common.Lan, "thNet"),
		lang.Text(common.Lan, "thTotalNet"),
		lang.Text(common.Lan, "thDate"),
		lang.Text(common.Lan, "thDateRate"),
		lang.Text(common.Lan, "thThisYearRate"),
		lang.Text(common.Lan, "thMonthRate"),
		lang.Text(common.Lan, "th3MonthRate"),
		lang.Text(common.Lan, "th6MonthRate"),
		lang.Text(common.Lan, "thYearRate"),
		lang.Text(common.Lan, "th3YearRate"),
		lang.Text(common.Lan, "thSinceFoundRate"),
	}
}

func formatIncratio(incratio string) string {
	value := util.Str2Decimal(incratio)
	s := value.Mul(decimal.NewFromInt(100)).StringFixed(2) + "%"
	if value.IsPositive() {
		return "+" + s
	} else {
		return s
	}
}

func (f Fund) GetValues() []string {
	var manager string
	for _, m := range f.FundManager {
		manager += m.Name + ", "
	}
	manager = strings.TrimRight(manager, ", ")
	return []string{
		f.FundCode, f.FundName, f.Type2IdDesc, f.FundScale, f.CTime, manager, f.Netvalue, f.TotalNetvalue, f.Navdate,
		formatIncratio(f.DayIncratio),
		formatIncratio(f.ThisYearIncratio),
		formatIncratio(f.MonthIncratio),
		formatIncratio(f.QuarterIncratio),
		formatIncratio(f.HalfYearIncratio),
		formatIncratio(f.YearIncratio),
		formatIncratio(f.ThreeYearIncratio),
		formatIncratio(f.ToThisDayIncratio),
	}
}

type Args struct {
	Tab       int    `json:"tab"`        // [1:在售基金 2:新发基金]
	PageSize  int    `json:"page_size"`  // 页大小
	PageNo    int    `json:"page_no"`    // 页码
	FundCode  string `json:"fund_code"`  // 基金代码
	Type      string `json:"type"`       // 基金类型 [x2001:混合型 x2002:股票型 x2003:债券型 x2005:货币型 x2006:QDII型]
	Company   string `json:"company"`    // 基金公司 [80000220:南方 80000223:嘉实 80000224:国泰 80000226:博时 80000229:易方达 80053708:汇添富 80055334:华泰博瑞 80280038:前海开源 80329448:中融]
	Scale     string `json:"scale"`      // 规模 [1:0~10亿 2:10~20亿 3:20~50亿 4:50~100亿 5:>100亿]
	Status    string `json:"status"`     // 状态 [0:暂停 1:在售]
	S         string `json:"s"`          // 搜索关键字
	Order     string `json:"order"`      // 排序字段 [fund_scale day_incratio month_incratio quarter_incratio half_year_incratio year_incratio this_year_incratio min_value]
	OrderType string `json:"order_type"` // 排序方式 [asc:升序 desc:降序]
	Time      int64  `json:"_"`          // 时间戳
}

func (a Args) String() string {
	return fmt.Sprintf(
		"tab=%v&"+
			"page_size=%v&"+
			"page_no=%v&"+
			"fund_code=%v&"+
			"type=%v&"+
			"company=%v&"+
			"scale=%v&"+
			"status=%v&"+
			"s=%v&"+
			"order=%v&"+
			"order_type=%v&"+
			"_=%v",
		a.Tab,
		a.PageSize,
		a.PageNo,
		a.FundCode,
		a.Type,
		a.Company,
		a.Scale,
		a.Status,
		a.S,
		a.Order,
		a.OrderType,
		a.Time,
	)
}

func FundMarketList() ([]Fund, error) {
	arg := Args{
		Tab:       1,
		PageSize:  20,
		PageNo:    0,
		FundCode:  "",
		Type:      "",
		Company:   "",
		Scale:     "",
		Status:    "1",
		S:         "",
		Order:     "day_incratio",
		OrderType: "desc",
		Time:      time.Now().Unix(),
	}
	url := fundUrl + "?" + arg.String()
	resp, err := common.Client.Get(url)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Items     []Fund `json:"items"`
			PageSize  int    `json:"page_size"`
			PageNo    int    `json:"page_no"`
			PageTotal string `json:"page_total"`
		} `json:"data"`
	}
	body := string(resp.Body())
	body = strings.ReplaceAll(body, `"fund_manager":"",`, `"fund_manager":[],`)
	err = util.Json.UnmarshalFromString(body, &result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return result.Data.Items, err
}
