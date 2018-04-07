package greatjson

import (
	"encoding/json"
	"errors"
)

// JSON 定义
type JSON struct {
	data interface{}
}

// New 解析 data 并生成 JSON 结构体指针
func New(data []byte) (*JSON, error) {
	j := new(JSON)

	if len(data) == 0 {
		data = []byte(`{}`)
	}

	err := j.Unmarshal(data)

	if err != nil {
		return nil, err
	}

	return j, nil
}

// Marshal 生成 data
func (j *JSON) Marshal() ([]byte, error) {
	return json.Marshal(&j.data)
}

// Unmarshal 解析 data
func (j *JSON) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &j.data)
}

// Get 获取 key 对应的 value
func (j *JSON) Get(key string) *JSON {
	m, err := j.Map()

	if err == nil {
		if value, ok := m[key]; ok {
			return &JSON{value}
		}
	}

	return &JSON{nil}
}

// Set 设置 key 对应的 value
func (j *JSON) Set(key string, value interface{}) {
	m, err := j.Map()

	if err != nil {
		return
	}

	m[key] = value
}

// Del 删除 key
func (j *JSON) Del(key string) {
	m, err := j.Map()

	if err != nil {
		return
	}

	delete(m, key)
}

// Map 访问 JSON
func (j *JSON) Map() (map[string]interface{}, error) {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m, nil
	}

	return nil, errors.New("类型断言错误")
}

// Data 获得 JSON 的 data 段
func (j *JSON) Data() interface{} {
	return j.data
}
