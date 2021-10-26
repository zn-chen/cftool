package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var buildData = ""

const version = "v1.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cftool",
	Long:  `All software has versions. This is cftool 's`,
	Run: func(cmd *cobra.Command, args []string) {
		if buildData == "" {
			buildData = time.Now().String()
		}

		fmt.Printf("cftool_%s, Built on %s\n", version, buildData)
	},
}
