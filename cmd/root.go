package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "cftool",
	Short: "config edit help tools",
	Long:  `With this tool, you can easily manipulate many types of configuration files, such as reading and writing`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "config file being operated on")
}
