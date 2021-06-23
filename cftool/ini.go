package cftool

import (
	"bytes"
	"gopkg.in/ini.v1"
)

type IniEdit struct {
	*ini.File
}

func NewIniEdit(data []byte) (*IniEdit, error) {
	cfg := ini.Empty()
	err := cfg.Append(data)
	if err != nil {
		return nil, err
	}

	return &IniEdit{cfg}, nil
}

// getValue 从json内容中读取一个key值
func (j IniEdit) GetValue(path []string) string {
	var session, key string
	if 1 == len(path) {
		session = ""
		key = path[0]
	} else if 2 == len(path) {
		session = path[0]
		key = path[1]
	} else {
		return ""
	}

	return j.Section(session).Key(key).Value()
}

// setValue 设置json内中指定json中的值
func (j IniEdit) SetValue(path []string, d interface{}) string {
	var session, key string
	if 1 == len(path) {
		session = ""
		key = path[0]
	} else if 2 == len(path) {
		session = path[0]
		key = path[1]
	} else {
		return ""
	}

	v := j.Section(session).Key(key).Value()
	j.Section(session).Key(key).SetValue(d.(string))
	return v
}

func (j IniEdit) String() string {
	buf := bytes.NewBuffer([]byte{})
	_, _ = j.WriteToIndent(buf, "")
	return buf.String()
}
