package main

import (
	"fmt"
	"github.com/airbike233/cftool/cftool"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("test.json")
	jsonEdit, _ := cftool.NewJsonEdit(data)

	fmt.Println(jsonEdit)
	fmt.Println(jsonEdit.GetValue([]string{"Base", "host"}))
	jsonEdit.SetValue([]string{"Base", "host"}, "dbapptestfile")
	jsonEdit.SetValue([]string{"Base", "host", "fuck"}, "dbapptestfile")
	fmt.Println(jsonEdit)
}
