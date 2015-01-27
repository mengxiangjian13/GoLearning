package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		fmt.Print("channel communication")
		c <- 4
		close(c) //关闭routines
	}()

	//以这种方式取channel中的值，取完毕后它不会自动关闭channel，进行下一次的等待。
	//所以利用这样的方式迭代channel，必须在适合的时候，close channel。
	for v := range c {
		fmt.Println("\n", v)
	}

	// =====================================
	runtime.GOMAXPROCS(runtime.NumCPU()) // 使用多核并发

	sumc := make(chan int, 10) // 设置缓存是10
	for i := 0; i < 10; i++ {
		go sum(sumc, i)
	}

	for i := 0; i < 10; i++ {
		<-sumc // 读取10次，完成channel
	}

	// =========================================
	// select: 和for range 差不多，也是为了取值的。
	// select{}里面为空，阻塞。
	o := make(chan bool)
	sc := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-sc:
				if !ok {
					fmt.Println("not ok")
				}
				fmt.Println("find sc:", v)
				o <- true
				break
			}
		}

	}()

	sc <- true
	<-o

	// ===========================================
	// select 设置超时
	tc := make(chan bool)

	select {
	case v := <-tc:
		fmt.Println(v)
	case v := <-time.After(2 * time.Second):
		fmt.Println("timeout", v)
	}

	// ==========================================
	// 一个goroutine互相发送信息若干次
	inc := make(chan string)
	go send(inc)
	a := <-inc
	fmt.Println(a)
	inc <- "hi"
	a = <-inc
	fmt.Println(a)
	inc <- "i am guozhenyan"
	a = <-inc
	fmt.Println(a)
}

func send(c chan string) {
	c <- "hello"
	a := <-c
	fmt.Println(a)
	c <- "i am mengxiangjian"
	a = <-c
	fmt.Println(a)
	c <- "ok"
}

func sum(a chan int, index int) {
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	fmt.Println(index, sum)
	a <- sum
}
