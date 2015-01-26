package main

import (
	"fmt"
)

type A struct {
	count int
}

func (a *A) Count(ch chan int) {
	a.count++
	fmt.Println("Counting", a.count)
	ch <- a.count
}

func main() {
	chs := make([]chan int, 10)

	a := A{}
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go a.Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
