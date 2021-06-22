package cftool

import (
	"gopkg.in/ini.v1"
	"strings"
)

type IniEdit map[string]interface{}

func NewIniEdit(data []byte) (*IniEdit, error) {
	var c IniEdit = make(map[string]interface{})

	cfg := ini.Empty()
	err := cfg.Append(data)
	if err != nil {
		return nil, err
	}

	sections := cfg.Sections()
	for i := 0; i < len(sections); i++ {
		section := sections[i]
		keys := section.Keys()
		for j := 0; j < len(keys); j++ {
			key := keys[j]
			value := cfg.Section(section.Name()).Key(key.Name()).String()
			c[section.Name()+"."+key.Name()] = value
		}
	}

	return &c, nil
}

// getValue 从json内容中读取一个key值
func (j IniEdit) GetValue(path []string) string {
	key := strings.Join(path, ".")
	if v, ok := j[key]; ok {
		return v.(string)
	}
	return ""
}

// setValue 设置json内中指定json中的值
func (j IniEdit) SetValue(path []string, d interface{}) string {
	key := strings.Join(path, ".")

	v, ok := j[key]
	j[key] = d
	if ok {
		return v.(string)
	}
	return ""
}
