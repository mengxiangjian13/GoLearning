package main

/*
	概述：
	Go语言的网络库支持很多种协议的网络，并且有更好的代码封装，使得非常好用。
	net/http 对http请求做了封装。

	重点：1. get请求。
		 2. response body的取出
*/

import (
	"bytes"
	"fmt"
	"io"
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

	// 请求到数据并取出
	result := bytes.NewBuffer(nil) // 建立一个缓冲，将请求返回写入
	var buf [512]byte              // 读repsonse body到缓冲的中间介质（比特流），[]byte的slice

	for {
		n, err := resp.Body.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
	}

	fmt.Println(string(result.Bytes()))
}
