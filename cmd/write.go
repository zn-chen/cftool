package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(writeCmd)
}

// writeCmd represents the read command
var writeCmd = &cobra.Command{
	Use:   "w",
	Short: "write to cftool file",
	Long: `Example:
cftool r -f config.ini session.key`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("write called")
	},
}
