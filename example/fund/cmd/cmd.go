package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	rootCmd := RootCmd()
	rootCmd.AddCommand(
		InitCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
