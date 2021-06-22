package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "r",
	Short: "write from cftool file",
	Long: `Example:
cftool r -f config.ini session.key`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("read called")
	},
}
