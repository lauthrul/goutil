package cmd

import (
	"fmt"
	"fund/config"
	"fund/model"
	"github.com/spf13/cobra"
)

func GroupCmd() *cobra.Command {
	var (
		configFile string
		add        bool
		remove     bool
		addFund    *[]string
		removeFund *[]string
		list       bool
		funds      bool
	)
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Group management",
		Long:  "Group management",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)
			if !add && !remove && len(*addFund) == 0 && len(*removeFund) == 0 && !list && !funds {
				_ = cmd.Usage()
				return
			}
			if add {
				_ = model.AddGroup(args...)
			}
			if remove {
				_ = model.RemoveGroup(args...)
			}
			if len(*addFund) != 0 {
				_ = model.AddGroupFund(*addFund, args...)
			}
			if len(*removeFund) != 0 {
				_ = model.RemoveGroupFund(*removeFund, args...)
			}
			if list {
				data, err := model.ListGroup()
				if err != nil {
					return
				}
				fmt.Println(data)
			}
			if funds {
				data, err := model.ListGroupFund(args...)
				if err != nil {
					return
				}
				groups := map[string][]model.FundGroup{}
				for _, d := range data {
					groups[d.Group] = append(groups[d.Group], d)
				}
				for g, funds := range groups {
					fmt.Printf("Funds in group: %s (%d)\n", g, len(funds))
					for _, f := range funds {
						fmt.Printf("\t%s(%s)\n", f.FundName, f.FundCode)
					}
				}
			}
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	cmd.Flags().BoolVarP(&add, "add", "a", false, "add group(s)")
	cmd.Flags().BoolVarP(&remove, "remove", "r", false, "remove group(s)")
	addFund = cmd.Flags().StringSliceP("add-fund", "f", nil, "add fund(s) to group")
	removeFund = cmd.Flags().StringSliceP("remove-fund", "F", nil, "remove fund(s) from group")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "list all groups")
	cmd.Flags().BoolVarP(&funds, "funds", "", false, "list funds by given group name")

	return cmd
}
