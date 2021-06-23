package cftool

import (
	"encoding/json"
)

type JsonEdit map[string]interface{}

func NewJsonEdit(data []byte) (JsonEdit, error) {
	var j JsonEdit
	if err := json.Unmarshal(data, &j); err != nil {
		return nil, err
	}
	return j, nil
}

// searchMap 递归搜索，返回递归后的值
// 当path用尽或map层数用尽时返回nil
// 如果搜索到对应值可通过edit进行操作
func searchMap(source map[string]interface{}, path []string, edit func(map[string]interface{}, string)) interface{} {
	if len(path) == 0 {
		return source
	}

	next, ok := source[path[0]]
	if ok {
		if len(path) == 1 {
			edit(source, path[0])
			return next
		}

		switch next.(type) {
		case map[string]interface{}:
			return searchMap(next.(map[string]interface{}), path[1:], edit)
		default:
			return nil
		}
	}
	return nil
}

// setMap 对json进行编辑, 如搜素路径不存在则创建
func setMap(source map[string]interface{}, path []string, v interface{}) interface{} {
	next, ok := source[path[0]]
	if ok {
		if len(path) == 1 {
			source[path[0]] = v
			return next
		}

		switch next.(type) {
		case map[string]interface{}:
			return setMap(next.(map[string]interface{}), path[1:], v)
		default:
			newValue := make(map[string]interface{})
			source[path[0]] = newValue
			return setMap(newValue, path[1:], v)
		}
	}
	if len(path) == 1 {
		source[path[0]] = v
		return nil
	}
	newValue := make(map[string]interface{})
	source[path[0]] = newValue
	return setMap(newValue, path[1:], v)
}

// getValue 从json内容中读取一个key值
func (j *JsonEdit) GetValue(path []string) string {
	v := searchMap(*j, path, func(m map[string]interface{}, s string) {})
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return ""
	}
}

// setValue 设置json内中指定json中的值
func (j *JsonEdit) SetValue(path []string, d interface{}) string {
	//v := searchMap(*j, path, func(m map[string]interface{}, s string) {
	//	m[s] = d
	//})
	v := setMap(*j, path, d)
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return ""
	}
}

func (j JsonEdit) String() string {
	b, _ := json.MarshalIndent(j, "", "  ")
	return string(b)
}
