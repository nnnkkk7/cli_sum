package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RoootCmd = &cobra.Command{
	Use:   "culc",
	Short: "command line calculator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}
