package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type Hand struct {
}

func (h Hand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	/*
		resp := func(w http.ResponseWriter, r *http.Request) {
			// 请求返回
			fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
		}
		// 监听根目录的请求，监听到后给予反馈
		http.HandleFunc("/", resp)
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/

	//以上就可以定义一个简单的httpServer了。如果想多自定义server，也是可以的。
	s := &http.Server{
		Addr:           ":8080",
		Handler:        Hand{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
