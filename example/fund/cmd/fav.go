package cmd

import (
	"fund/config"
	"fund/model"
	"github.com/lauthrul/goutil/log"
	"github.com/spf13/cobra"
	"strings"
)

func FavCmd() *cobra.Command {
	var (
		configFile string
		funds      string
	)

	rootCmd := &cobra.Command{
		Use:   "fav",
		Short: "add favorites",
		Long:  "add favorites by specified fund codes, split by `,`",
		Run: func(cmd *cobra.Command, args []string) {
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
				basic, err := api.GetFundBasic(code)
				if err != nil {
					log.ErrorF("get fund[%s] basic fail: %s", code, err.Error())
					continue
				}
				managers, err := api.GetFundManager(code)
				if err != nil {
					log.ErrorF("get fund[%s] manager fail: %s", code, err.Error())
					continue
				}
				netValues, err := api.GetFundNetValue(code, "2020-01-01", "2021-01-01")
				if err != nil {
					log.ErrorF("get fund[%s] net values fail: %s", code, err.Error())
					continue
				}
				holdings, err := api.GetFundHoldingStock(code, 2020)
				if err != nil {
					log.ErrorF("get fund[%s] holdings fail: %s", code, err.Error())
					continue
				}
				_ = model.SaveFundBasic(basic)
				_ = model.SaveFundManager(managers...)
				_ = model.SaveFundStockHoldings(holdings...)
				_ = model.SaveFundNetValue(netValues...)
			}
		},
	}

	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	rootCmd.Flags().StringVarP(&funds, "funds", "f", "", "fund codes, split by `,`")

	return rootCmd
}
