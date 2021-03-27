package cmd

import (
	"fmt"
	"fund/config"
	"fund/model"
	"github.com/lauthrul/goutil/log"
	"github.com/scylladb/go-set/strset"
	"github.com/spf13/cobra"
	"time"
)

func InitCmd() *cobra.Command {
	var (
		configFile string
		types      []string
		fav        bool
		group      []string
		code       []string
		name       string
	)

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init fund info",
		Long:  "Init fund info([basic,manager,net_value,holding,stock])",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)

			listArg := model.ListFundGroupArg{IsFav: -1, Group: group, Code: append(code, args...), Name: name}
			if cmd.Flags().Changed("fav") {
				listArg.IsFav = 0
				if fav {
					listArg.IsFav = 1
				}
			}
			funds, _ := model.ListFundGroup(listArg)

			codes := strset.New()
			for _, f := range funds {
				codes.Add(f.Code)
			}
			if codes.Size() == 0 {
				fmt.Println("please specific fund(s) code")
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

			log.DebugF("fund init type:%s, funds:%s", types, codes)

			api := model.NewEastMoneyApi()

			for _, code := range codes.List() {
				if code == "" {
					continue
				}
				log.Debug("init fund", code)

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
					if date, err := time.Parse(model.DATEFORMAT, basic.CreateDate); err == nil {
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
					date, _ := model.GetNextFundNetValueDate(code)
					if date == "" {
						date = createDate
					}
					var netValues []model.FundNetValue
					pageSize := 100
					for page := 1; page != 0; {
						values, nextPage, err := api.GetFundNetValue(
							code,
							date,
							time.Now().Format(model.DATEFORMAT),
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
	cmd.Flags().StringSliceVarP(&types, "type", "t", []string{"all"}, `init type, can be one or all of [basic,manager,net_value,holdings,all], split by ","`)
	cmd.Flags().BoolVarP(&fav, "fav", "F", false, `choose fund in fav`)
	cmd.Flags().StringSliceVarP(&group, "group", "G", nil, `choose fund in given group(s)`)
	cmd.Flags().StringSliceVarP(&code, "code", "C", nil, `choose fund in given code(s)`)
	cmd.Flags().StringVarP(&name, "name", "N", "", `choose fund with name(fuzzy match)`)
	return cmd
}
