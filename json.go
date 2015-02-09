package main

import (
	"encoding/json"
	"fmt"
)

/*
	go 可以将一个struct实例变成一个json。
	并且可以将此json再精确的解析成一个struct实例。
*/

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func main() {
	b := Book{
		"Go语言编程",
		[]string{"mengxiangjian", "guozhenyan"},
		"ituring.com.cn",
		true,
		9.99,
	}

	bjson, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	//解析时候如果声明解析成Book类型，就会顺利解析为Book实例
	var newbook Book
	json.Unmarshal(bjson, &newbook)
	fmt.Println(newbook)

	// 如果声明的是一般类型，则会解析成一个map[string]interface{}类型
	var newb interface{}

	json.Unmarshal(bjson, &newb)
	a, ok := newb.(map[string]interface{})
	if ok {
		fmt.Println(a)
	}
}
