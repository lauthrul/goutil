package cmd

import (
	"fund/config"
	"fund/model"
	"github.com/spf13/cobra"
)

func FundCmd() *cobra.Command {
	var (
		configFile  string
		addFav      bool
		removeFav   bool
		addGroup    string
		removeGroup string
		setRemark   string
	)

	cmd := &cobra.Command{
		Use:   "fund",
		Short: "Fund management",
		Long:  "Fund management",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)
			isSetRemark := cmd.Flags().Changed("set-remark")
			if !addFav && !removeFav && addGroup == "" && removeGroup == "" && !isSetRemark {
				_ = cmd.Usage()
				return
			}
			if addFav {
				_ = model.SetFundFav(true, args...)
			}
			if removeFav {
				_ = model.SetFundFav(false, args...)
			}
			if addGroup != "" {
				_ = model.AddFundGroup(addGroup, args...)
			}
			if removeGroup != "" {
				_ = model.RemoveFundGroup(removeGroup, args...)
			}
			if isSetRemark {
				_ = model.SetFundRemark(setRemark, args...)
			}
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	cmd.Flags().BoolVarP(&addFav, "add-fav", "f", false, "add fund(s) to favorites")
	cmd.Flags().BoolVarP(&removeFav, "remove-fav", "F", false, "remove fund(s) from favorites")
	cmd.Flags().StringVarP(&addGroup, "add-group", "g", "", "add fund(s) to group")
	cmd.Flags().StringVarP(&removeGroup, "remove-group", "G", "", "remove fund(s) from group")
	cmd.Flags().StringVarP(&setRemark, "set-remark", "m", "", "set fund(s) remark")

	return cmd
}
