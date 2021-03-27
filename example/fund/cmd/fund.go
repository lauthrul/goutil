package cmd

import (
	"fmt"
	"fund/config"
	"fund/model"
	"github.com/olekukonko/tablewriter"
	"github.com/scylladb/go-set/strset"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

func FundCmd() *cobra.Command {
	var (
		configFile  string
		setFav      bool
		setRemark   string
		addGroup    []string
		removeGroup []string
		list        bool
		fav         bool
		group       []string
		code        []string
		name        string
		manager     bool
		estimate    bool
		order       string
		holdings    bool
		season      string
	)

	cmd := &cobra.Command{
		Use:   "fund",
		Short: "Fund management",
		Long:  "Fund management",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)

			isSetFav := cmd.Flags().Changed("set-fav")
			isSetRemark := cmd.Flags().Changed("set-remark")
			if !isSetFav && !isSetRemark && len(addGroup) == 0 && len(removeGroup) == 0 &&
				!list && !manager && !estimate && !holdings {
				_ = cmd.Usage()
				return
			}

			listArg := model.ListFundGroupArg{IsFav: -1, Group: group, Code: append(code, args...), Name: name}
			if cmd.Flags().Changed("fav") {
				listArg.IsFav = 0
				if fav {
					listArg.IsFav = 1
				}
			}
			funds, _ := model.ListFundGroup(listArg)

			if list {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(model.ViewFundGroup{}.GetTitles())
				table.SetAlignment(tablewriter.ALIGN_RIGHT)
				for _, v := range funds {
					table.Append(v.GetValues())
				}
				table.Render()
				return
			}

			codes := strset.New()
			for _, f := range funds {
				codes.Add(f.Code)
			}
			if codes.Size() == 0 {
				fmt.Println("please specific fund(s) code")
				return
			}

			if isSetFav {
				_ = model.SetFundFav(setFav, codes.List()...)
			}
			if isSetRemark {
				_ = model.SetFundRemark(setRemark, codes.List()...)
			}
			if len(addGroup) > 0 {
				_ = model.AddGroupFund(codes.List(), addGroup...)
			}
			if len(removeGroup) > 0 {
				_ = model.RemoveGroupFund(codes.List(), addGroup...)
			}
			if manager {
				data, err := model.ListFundManager(codes.List()...)
				if err != nil {
					return
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(model.ViewFundManager{}.GetTitles())
				table.SetAlignment(tablewriter.ALIGN_RIGHT)
				for _, v := range data {
					table.Append(v.GetValues())
				}
				table.Render()
			}
			if estimate {
				api := model.NewEastMoneyApi()
				var data []model.FundEstimate
				for _, code := range codes.List() {
					result, err := api.GetFundEstimate(code)
					if err != nil {
						continue
					}
					data = append(data, result)
				}
				if order == "asc" {
					sort.Slice(data, func(i, j int) bool {
						return data[i].EstimateRate < data[j].EstimateRate
					})
				} else if order == "desc" {
					sort.Slice(data, func(i, j int) bool {
						return data[i].EstimateRate > data[j].EstimateRate
					})
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(model.FundEstimate{}.GetTitles())
				table.SetAlignment(tablewriter.ALIGN_RIGHT)
				for _, v := range data {
					table.Append(v.GetValues())
				}
				table.Render()
			}
			if holdings {
				var data []model.FundHoldingStock
				if season == "" {
					results, _ := model.GetLatestHoldingStock(codes.List()...)
					for _, r := range results {
						data = append(data, r...)
					}
				} else {
					data, _ = model.ListFundHoldings(season, codes.List()...)
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(model.FundHoldingStock{}.GetTitles())
				table.SetAlignment(tablewriter.ALIGN_RIGHT)
				for _, v := range data {
					table.Append(v.GetValues())
				}
				table.Render()
			}
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	cmd.Flags().BoolVarP(&setFav, "set-fav", "f", false, "set fund(s) favorites")
	cmd.Flags().StringVarP(&setRemark, "set-remark", "m", "", "set fund(s) remark")
	cmd.Flags().StringSliceVarP(&addGroup, "add-to-group", "g", nil, "add fund(s) to group(s)")
	cmd.Flags().StringSliceVarP(&removeGroup, "remove-to-group", "d", nil, "remove fund(s) from group(s)")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "list all funds")
	cmd.Flags().BoolVarP(&fav, "fav", "F", false, `choose fund in fav`)
	cmd.Flags().StringSliceVarP(&group, "group", "G", nil, `choose fund in given group(s)`)
	cmd.Flags().StringSliceVarP(&code, "code", "C", nil, `choose fund in given code(s)`)
	cmd.Flags().StringVarP(&name, "name", "N", "", `choose fund with name(fuzzy match)`)
	cmd.Flags().BoolVarP(&manager, "manager", "M", false, `show fund manager info`)
	cmd.Flags().BoolVarP(&estimate, "estimate", "e", false, `show fund estimate`)
	cmd.Flags().StringVarP(&order, "order", "o", "", `fund estimate order, (asc|desc|none)`)
	cmd.Flags().BoolVarP(&holdings, "holding-stock", "H", false, `show fund holding stocks`)
	cmd.Flags().StringVarP(&season, "season", "s", "", `specific fund holding stock season, use together with "-H, --holding-stock"`)

	return cmd
}
