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
		list       bool
	)
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Group management",
		Long:  "Group management",
		Run: func(cmd *cobra.Command, args []string) {
			conf := config.Load(configFile)
			Init(conf)
			if !add && !remove && !list {
				_ = cmd.Usage()
				return
			}
			if add {
				_ = model.AddGroup(args...)
			}
			if remove {
				_ = model.RemoveGroup(args...)
			}
			if list {
				data, err := model.ListGroup()
				if err != nil {
					return
				}
				fmt.Println(data)
			}
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.json", "config file")
	cmd.Flags().BoolVarP(&add, "add", "a", false, "add group")
	cmd.Flags().BoolVarP(&remove, "remove", "r", false, "remove group")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "list all groups")

	return cmd
}
