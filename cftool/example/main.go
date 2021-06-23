package main

import (
	"fmt"
	"github.com/airbike233/cftool/cftool"
)

var testDataJson = `{
  "Base": {
    "host": "192.168.140.189",
    "port": 1098
  },
  "Test": {
    "key": "value"
  }
}`

var testDataINI = `name = oscar
[Base]
host = 192.168.140.202
port = 123
`

func jsonTest() {
	jsonEdit, _ := cftool.NewJsonEdit([]byte(testDataJson))

	fmt.Println(jsonEdit)
	fmt.Println(jsonEdit.GetValue([]string{"Base", "host"}))
	jsonEdit.SetValue([]string{"Base", "host"}, "dbapptestfile")
	jsonEdit.SetValue([]string{"Base", "host", "fuck"}, "dbapptestfile")
	fmt.Println(jsonEdit)
	jsonEdit.SetValue([]string{"Base", "host"}, "dbapptestfile")
	fmt.Println(jsonEdit)
}

func iniTest() {
	iniEdit, _ := cftool.NewIniEdit([]byte(testDataINI))

	fmt.Println(iniEdit)

	fmt.Println(iniEdit.GetValue([]string{"name"}))
	fmt.Println(iniEdit.GetValue([]string{"Base", "host"}))
	fmt.Println(iniEdit.GetValue([]string{"Base", "test"}))

	iniEdit.SetValue([]string{"Base", "host"}, "127.0.0.1")
	fmt.Println(iniEdit)

	iniEdit.SetValue([]string{"Base", "host1"}, "127.0.0.1")
	fmt.Println(iniEdit)
}

func main() {
	jsonTest()
	iniTest()
}
