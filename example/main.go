package main

import (
	"fmt"
	"../greatjson"
)

func main() {
	data := []byte(`{
		"test_data": {
			"int": 1,
			"float": 1.023,
			"string": "Hello World",
			"bool": true,
			"array": [1, 2.333, "Hello World", true, null],
			"array_with_object": [
				{"key_one": 1},
				{"key_one": 1, "key_two": 2}
			]
		}
	}`)

	js, err := greatjson.NewFrom(data)

	if err != nil {
		fmt.Println(err)
	}

	// 获取转换后的 json 实例
	fmt.Println("获取转换后的 json 实例")
	fmt.Println(js.Map())
	fmt.Println()

	// 检查指定 key 的 value 存在性
	fmt.Println("检查指定 key 的 value 存在性")
	value, ok := js.CheckGet("test_existence")
	fmt.Println(value, ok)
	fmt.Println()

	// 获取 int
	fmt.Println("获取 int")
	fmt.Println(js.Get("test_data").Get("int").Int())
	fmt.Println()

	// 获取 float32
	fmt.Println("获取 float32")
	fmt.Println(js.Get("test_data").Get("float").Float32())
	fmt.Println()

	// 获取 float64
	fmt.Println("获取 float64")
	fmt.Println(js.Get("test_data").Get("float").Float64())
	fmt.Println()

	// 获取 string
	fmt.Println("获取 string")
	fmt.Println(js.Get("test_data").Get("string").String())
	fmt.Println()

	// 获取 bool
	fmt.Println("获取 bool")
	fmt.Println(js.Get("test_data").Get("bool").Bool())
	fmt.Println()

	// 获取 array
	fmt.Println("获取 array")
	fmt.Println(js.Get("test_data").Get("array").Array())
	fmt.Println()

	// 获取 array 里的指定 index 的值
	fmt.Println("获取 array 里的指定 index 的值")
	fmt.Println(js.Get("test_data").Get("array").GetIndex(0).Int())
	fmt.Println(js.Get("test_data").Get("array").GetIndex(1).Float32())
	fmt.Println(js.Get("test_data").Get("array").GetIndex(2).String())
	fmt.Println(js.Get("test_data").Get("array").GetIndex(3).Bool())
	fmt.Println(js.Get("test_data").Get("array").GetIndex(4).Data())
	fmt.Println(js.Get("test_data").Get("array_with_object").GetIndex(1).Get("key_one").Int())
	fmt.Println()

	js, err = greatjson.New()

	if err != nil {
		fmt.Println(err)
	}

	js.Set("a", 1)

	data, err = js.Marshal()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
