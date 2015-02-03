package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	// "time"
)

type Hand struct {
}

func (h Hand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 根据不同的请求地址不同的返回
	if r.URL.String() == "/hello" {
		// 输出
		io.WriteString(w, "hello,world!")
		return
	} else if r.URL.String() == "/bye" {
		io.WriteString(w, "bye,world!")
		return
	}

	//另一种输出形式
	fmt.Fprintf(w, "URL, %q", html.EscapeString(r.URL.Path))
}

func main() {
	/*
		// 最高层封装
		resp := func(w http.ResponseWriter, r *http.Request) {
			// 请求返回
			fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
		}
		// 监听根目录的请求，监听到后给予反馈
		http.HandleFunc("/", resp)
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/

	// 低一层封装
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	serverMux := http.NewServeMux()
	serverMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd)))) //fileServer返回一个handler，不用自己handle
	log.Fatal(http.ListenAndServe(":8080", serverMux))

	/*
		// 最底层封装
		//以上就可以定义一个简单的httpServer了。如果想多自定义server，也是可以的。
		s := &http.Server{
			Addr:           ":8080",
			Handler:        &Hand{},
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(s.ListenAndServe())
	*/
}
