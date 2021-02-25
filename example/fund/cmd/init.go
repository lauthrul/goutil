package cmd

import (
	"fund/config"
	"fund/model"
	"github.com/lauthrul/goutil/log"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

func InitCmd() *cobra.Command {
	var (
		configFile string
		funds      string
		typ        string
	)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "init funds info",
		Long:  "init funds info([basic,manager,net_value,holding,stock]) by specified fund codes, split by `,`",
		Run: func(cmd *cobra.Command, args []string) {
			types := strings.Split(typ, ",")
			if len(types) == 0 || funds == "" {
				cmd.Usage()
				return
			}

			fnType := func(s string) bool {
				for _, t := range types {
					if t == "all" || t == s {
						return true
					}
				}
				return false
			}

			conf := config.Load(configFile)
			Init(conf)
			codes := strings.Split(funds, ",")
			if len(codes) == 0 {
				return
			}
			api := model.NewEastMoneyApi()

			for _, code := range codes {
				if code == "" {
					continue
				}

				tx, err := model.GetDB().Begin()
				if err != nil {
					log.ErrorF("db transaction fail: %s", err)
					return
				}

				// get fund basic
				var createDate string
				var years []int
				if fnType("basic") {
					basic, err := api.GetFundBasic(code)
					if err != nil {
						log.ErrorF("get fund[%s] basic fail: %s", code, err.Error())
						_ = tx.Commit()
						continue
					}
					err = model.SaveFundBasic(tx, basic)
					if err != nil {
						log.ErrorF("SaveFundBasic fail: %s", err)
					}

					createDate = basic.CreateDate
					if date, err := time.Parse("2006-01-02", basic.CreateDate); err == nil {
						years = append(years, date.Year())
						for y := date.Year() + 1; y <= time.Now().Year(); y++ {
							years = append(years, y)
						}
					}
				}

				// get fund manager
				if fnType("manager") {
					managers, experiences, err := api.GetFundManager(code)
					if err != nil {
						log.ErrorF("get fund[%s] manager fail: %s", code, err.Error())
					}
					err = model.SaveManager(tx, managers...)
					if err != nil {
						log.ErrorF("SaveManager fail: %s", err)
					}
					err = model.SaveManagerExperience(tx, experiences...)
					if err != nil {
						log.ErrorF("SaveManagerExperience fail: %s", err)
					}
				}

				// holding stocks
				if fnType("holdings") {
					var holdings []model.FundHoldingStock
					for _, y := range years {
						holds, err := api.GetFundHoldingStock(code, y)
						if err != nil {
							log.ErrorF("get fund[%s] holdings fail: %s", code, err.Error())
						} else {
							holdings = append(holdings, holds...)
						}
					}
					err = model.SaveFundStockHoldings(tx, holdings...)
					if err != nil {
						log.ErrorF("SaveFundStockHoldings fail: %s", err)
					}
				}

				// get  and net values
				if fnType("net_value") {
					var netValues []model.FundNetValue
					pageSize := 100
					for page := 1; page != 0; {
						values, nextPage, err := api.GetFundNetValue(
							code,
							createDate,
							time.Now().Format("2006-01-02"),
							page,
							pageSize,
						)
						if err != nil {
							log.ErrorF("get fund[%s] net values fail: %s", code, err.Error())
						} else {
							netValues = append(netValues, values...)
						}
						page = nextPage
					}
					err = model.SaveFundNetValue(tx, netValues...)
					if err != nil {
						log.ErrorF("SaveFundNetValue fail: %s", err)
					}
				}

				// commit to db
				if err = tx.Commit(); err != nil {
					log.ErrorF("db commit fail: %s", err)
				}
			}
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	cmd.Flags().StringVarP(&funds, "funds", "f", "", "fund codes, split by `,`")
	cmd.Flags().StringVarP(&typ, "type", "t", "", "init type, can be one or all of [basic,manager,net_value,holdings,all]")

	return cmd
}
