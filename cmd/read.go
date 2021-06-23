package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/airbike233/cftool/cftool"
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
	Run: readFile,
}

func readFile(_ *cobra.Command, args []string) {
	if 1 > len(args) {
		Exit("args error")
	}
	keys := strings.Split(args[0], ".")
	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		Exit(err)
	}

	switch FileType(cfgFile) {
	case "ini":
		i, err := cftool.NewJsonEdit(data)
		if err == nil {
			_, _ = fmt.Fprint(os.Stdout, i.GetValue(keys))
		} else {
			Exit(err)
		}
	case "json":
		j, err := cftool.NewJsonEdit(data)
		if err == nil {
			_, _ = fmt.Fprint(os.Stdout, j.GetValue(keys))
		} else {
			Exit(err)
		}
	}
	return
}

// FileType 返回文件后缀
func FileType(path string) string {
	s := strings.Split(path, ".")
	if len(s) >= 2 {
		return s[len(s)-1]
	}
	return ""
}

func Exitf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func Exit(des interface{}) {
	_, _ = fmt.Fprint(os.Stderr, des)
	os.Exit(1)
}
