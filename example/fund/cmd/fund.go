package cmd

import (
	"fund/config"
	"fund/model"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

func FundCmd() *cobra.Command {
	var (
		configFile string
		addFav     bool
		removeFav  bool
		setRemark  string
		list       bool
		withFav    bool
		code       *[]string
		name       string
		manager    bool
	)

	cmd := &cobra.Command{
		Use:   "fund",
		Short: "Fund management",
		Long:  "Fund management",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)
			isSetRemark := cmd.Flags().Changed("set-remark")
			if !addFav && !removeFav && !isSetRemark && !list && !manager {
				_ = cmd.Usage()
				return
			}
			if addFav {
				_ = model.SetFundFav(true, args...)
			}
			if removeFav {
				_ = model.SetFundFav(false, args...)
			}
			if isSetRemark {
				_ = model.SetFundRemark(setRemark, args...)
			}
			if list {
				data, err := model.ListFund(model.ListFundArg{
					IsFav: withFav,
					Code:  *code,
					Name:  name,
				})
				if err != nil {
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(model.FundBasic{}.Titles())
				table.SetAlignment(tablewriter.ALIGN_RIGHT)
				for _, v := range data {
					table.Append(v.Values())
				}
				table.Render()
			}
			if manager {
				data, err := model.ListFundManager(args...)
				if err != nil {
					return
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(model.FundManager{}.Titles())
				table.SetAlignment(tablewriter.ALIGN_RIGHT)
				for _, v := range data {
					table.Append(v.Values())
				}
				table.Render()
			}
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	cmd.Flags().BoolVarP(&addFav, "add-fav", "f", false, "add fund(s) to favorites")
	cmd.Flags().BoolVarP(&removeFav, "remove-fav", "F", false, "remove fund(s) from favorites")
	cmd.Flags().StringVarP(&setRemark, "set-remark", "m", "", "set fund(s) remark")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "list all funds")
	cmd.Flags().BoolVarP(&withFav, "with-fav", "", false, `list funds with fav, use together with "-l,--list"`)
	code = cmd.Flags().StringSliceP("code", "", nil, `list fund with code, use together with "-l,--list"`)
	cmd.Flags().StringVarP(&name, "name", "", "", `list fund with name(fuzzy search), use together with "-l,--list"`)
	cmd.Flags().BoolVarP(&manager, "manager", "M", false, `show fund manager info`)

	return cmd
}
