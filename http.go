package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://mobile.cfda.gov.cn/datasearch/QueryList?pageIndex=1&pageSize=10&tableId=27&searchF=Quick Search&searchK=CoCr"
	r, e := http.Get(url)
	if e != nil {
		fmt.Println("error:", e.Error())
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println(string(b))
}
