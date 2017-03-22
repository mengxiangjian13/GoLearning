package main

/*
	概述：
	Go语言的网络库支持很多种协议的网络，并且有更好的代码封装，使得非常好用。
	net/http 对http请求做了封装。

	重点：1. get请求。
		 2. response body的取出
*/

import (
	"fmt"
	"io/ioutil"
	"net/http" // 网络库
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage:%s <url>\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1] // 取得url
	// get 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	defer resp.Body.Close() //调用http get者有责任close body

	b, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		fmt.Println("error:", err.Error())
	}

	fmt.Println(string(b))
}
