package main

import (
	"log"
	"../greatjson"
)

func main() {
	data := []byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 8000000000
		}
	}`)

	js, err := greatjson.New(data)

	if err != nil {
		log.Println(err)
	}

	log.Println(js.Get("test").Get("bignum"))

	js, err = greatjson.New([]byte{})

	if err != nil {
		log.Println(err)
	}

	js.Set("a", 1)

	data, err = js.Marshal()

	if err != nil {
		log.Println(err)
	}

	log.Println(string(data))
}
