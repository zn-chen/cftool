package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/zn-chen/cftool/cftool"
)

func init() {
	rootCmd.AddCommand(writeCmd)
}

// writeCmd represents the read command
var writeCmd = &cobra.Command{
	Use:   "w",
	Short: "write to cftool file",
	Long: `Example:
cftool w -f config.ini session.key=value > config.ini`,
	Run: writeFile,
}

func writeFile(cmd *cobra.Command, args []string) {
	if 1 > len(args) {
		_ = cmd.Help()
		os.Exit(1)
	}
	kv := strings.Split(args[0], "=")
	if 2 > len(kv) {
		_ = cmd.Help()
		os.Exit(1)
	}
	value := kv[1]
	keys := strings.Split(kv[0], ".")
	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		Exit(err)
	}

	switch FileType(cfgFile) {
	case "ini":
		i, err := cftool.NewIniEdit(data)
		if err == nil {
			_ = i.SetValue(keys, value)
			_, _ = fmt.Fprint(os.Stdout, i)
		} else {
			Exit(err)
		}
	case "json":
		j, err := cftool.NewJsonEdit(data)
		if err == nil {
			_ = j.SetValue(keys, value)
			_, _ = fmt.Fprint(os.Stdout, j)
		} else {
			Exit(err)
		}
	}
	return
}
