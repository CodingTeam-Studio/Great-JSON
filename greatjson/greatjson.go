package greatjson

import (
	"encoding/json"
	"errors"
	"reflect"
)

// JSON 定义
type JSON struct {
	data interface{}
}

// New 生成空 JSON
func New() (*JSON, error) {
	j := new(JSON)

	err := j.Unmarshal([]byte(`{}`))

	if err != nil {
		return nil, err
	}

	return j, nil
}

// NewFrom 解析 data 并生成 JSON 结构体指针
func NewFrom(data []byte) (*JSON, error) {
	j := new(JSON)

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

// CheckGet 检查 key 存在性
func (j *JSON) CheckGet(key string) (*JSON, bool) {
	m, err:= j.Map()

	if err == nil {
		if value, ok := m[key]; ok {
			return &JSON{value}, true
		}
	}

	return nil, false
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

// GetIndex 获取指定 index 的 value
func (j *JSON) GetIndex(index int) *JSON {
	arr, err := j.Array()

	if err == nil {
		if len(arr) > index {
			return &JSON{arr[index]}
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

	return nil, errors.New("类型断言失败")
}

// Array 转义为数组
func (j *JSON) Array() ([]interface{}, error) {
	if arr, ok := (j.data).([]interface{}); ok {
		return arr, nil
	}

	return nil, errors.New("数组类型断言失败")
}

// Data 获得 JSON 的 data 段
func (j *JSON) Data() interface{} {
	return j.data
}

// Bool 转义为 bool
func (j *JSON) Bool() (bool, error) {
	if b, ok := (j.data).(bool); ok {
		return b, nil
	}

	return false, errors.New("布尔类型断言失败")
}

func (j *JSON) String() (string, error) {
	if str, ok := (j.data).(string); ok {
		return str, nil
	}

	return "", errors.New("字符串类型断言失败")
}

// Int 转义为 int
func (j *JSON) Int() (int, error) {
	switch j.data.(type) {
	case json.Number:
		i, err := j.data.(json.Number).Int64()
		return int(i), err
	case int, int8, int16, int32, int64:
		i := int(reflect.ValueOf(j.data).Int())
		return i, nil
	case uint, uint8, uint16, uint32, uint64:
		i := int(reflect.ValueOf(j.data).Uint())
		return i, nil
	case float32, float64:
		i := int(reflect.ValueOf(j.data).Float())
		return i, nil
	}

	return 0, errors.New("数值无法转换为 int")
}

// Float32 转义为 float32
func (j *JSON) Float32() (float32, error) {
	switch j.data.(type) {
	case json.Number:
		f, err := j.data.(json.Number).Float64()
		return float32(f), err
	case int, int8, int16, int32, int64:
		f := float32(reflect.ValueOf(j.data).Int())
		return f, nil
	case uint, uint8, uint16, uint32, uint64:
		f := float32(reflect.ValueOf(j.data).Uint())
		return f, nil
	case float32, float64:
		f := float32(reflect.ValueOf(j.data).Float())
		return f, nil
	}

	return 0, errors.New("数值无法转换为 float32")
}

// Float64 转义为 float64
func (j *JSON) Float64() (float64, error) {
	switch j.data.(type) {
	case json.Number:
		f, err := j.data.(json.Number).Float64()
		return f, err
	case int, int8, int16, int32, int64:
		f := float64(reflect.ValueOf(j.data).Int())
		return f, nil
	case uint, uint8, uint16, uint32, uint64:
		f := float64(reflect.ValueOf(j.data).Uint())
		return f, nil
	case float32, float64:
		f := float64(reflect.ValueOf(j.data).Float())
		return f, nil
	}

	return 0, errors.New("数值无法转换为 float64")
}
