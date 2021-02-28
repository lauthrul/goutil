package cmd

import (
	"fmt"
	"fund/config"
	"fund/model"
	"github.com/lauthrul/goutil/log"
	"github.com/spf13/cobra"
	"sort"
)

type Similarity struct {
	Code1   string
	Name1   string
	Stocks1 []string
	Code2   string
	Name2   string
	Stocks2 []string
	R1      float64
	R2      float64
}

type MostHolding struct {
	StockName string
	Funds     []string
}

func AnalysisCmd() *cobra.Command {
	var (
		configFile string
		withFav    bool
		group      []string
	)

	cmd := &cobra.Command{
		Use:   "analysis",
		Short: "Analysis funds",
		Long:  "Analysis funds with similarity, most holding stocks or some other aspects.",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)

			codes := args
			if withFav {
				funds, err := model.ListFund(model.ListFundArg{
					IsFav: true,
				})
				if err != nil {
					return
				}
				for _, f := range funds {
					codes = append(codes, f.Code)
				}
			}
			if len(group) > 0 {
				funds, err := model.ListGroupFund(group...)
				if err != nil {
					return
				}
				for _, f := range funds {
					codes = append(codes, f.FundCode)
				}
			}
			if len(codes) == 0 {
				fmt.Println(`no funds to analysis. you can specific fund codes, or use flag "-f,--with-Fav" to analysis favorite funds`)
				return
			}

			log.DebugF("analysis funds:%s", codes)

			holdings, err := model.GetLatestHoldingStock(codes...)
			if err != nil {
				log.Error(err)
				return
			}

			var (
				similarities []Similarity
				mostHoldings []MostHolding
			)
			stocks := map[string][]string{}
			for i := 0; i < len(holdings); i++ {
				h1 := holdings[i]
				n1 := len(h1)
				var s1 []string
				for _, v := range h1 {
					s1 = append(s1, v.StockName)
					stocks[v.StockName] = append(stocks[v.StockName], v.FundName)
				}
				for j := i + 1; j < len(holdings); j++ {
					h2 := holdings[j]
					n2 := len(h2)
					union := append(h1, h2...)
					set := map[string]int{}
					for _, v := range union {
						set[v.StockCode]++
					}
					n := 0
					for _, v := range set {
						if v > 1 {
							n++
						}
					}
					var s2 []string
					for _, v := range h2 {
						s2 = append(s2, v.StockName)
					}
					r1 := float64(n) / float64(n1) * 100
					r2 := float64(n) / float64(n2) * 100
					similarities = append(similarities, Similarity{
						Code1:   h1[0].FundCode,
						Name1:   h1[0].FundName,
						Stocks1: s1,
						Code2:   h2[0].FundCode,
						Name2:   h2[0].FundName,
						Stocks2: s2,
						R1:      r1,
						R2:      r2,
					})
				}
			}

			s := "【1. Fund similarity】"
			fmt.Println(s)
			log.Info(s)
			sort.Slice(similarities, func(i, j int) bool {
				return similarities[i].R1 > similarities[j].R1
			})
			for _, item := range similarities {
				s := fmt.Sprintf("%s(%s) <--%.2f%% : %.2f%%--> %s(%s): \n\t%q\n\t%q",
					item.Name1,
					item.Code1,
					item.R1,
					item.R2,
					item.Name2,
					item.Code2,
					item.Stocks1,
					item.Stocks2,
				)
				fmt.Println(s)
				log.Info(s)
			}

			s = "【2. Stocks holdings】"
			fmt.Println(s)
			log.Info(s)
			for k, v := range stocks {
				mostHoldings = append(mostHoldings, MostHolding{
					StockName: k,
					Funds:     v,
				})
			}
			sort.Slice(mostHoldings, func(i, j int) bool {
				return len(mostHoldings[i].Funds) > len(mostHoldings[j].Funds)
			})
			for _, v := range mostHoldings {
				s := fmt.Sprintf("\t%s(%d):%q", v.StockName, len(v.Funds), v.Funds)
				fmt.Println(s)
				log.Info(s)
			}
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	cmd.Flags().BoolVarP(&withFav, "with-fav", "f", false, `analysis funds with fav`)
	cmd.Flags().StringSliceVarP(&group, "group", "g", nil, `analysis funds in given group name(s)`)

	return cmd
}
