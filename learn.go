// import page main
package main

// import package
import (
	"fmt"
)

//const
const (
	a = 1
	b        //不赋值就等于上面的常量
	c = iota // 赋值iota意思是常量+1,如果前面没有常量为0，有常量则为常量+1。
	d        // 不赋值就等于上面的常量，因为上面的常量为iota，所以还是要加1
)

var (
	e = "good"
	f = "day"
)

type 中文 string

// func
func run(arg string) {
	// nothing
	fmt.Println(arg)
}

// main founction
func main() {

	// print const
	fmt.Println(a, b, c, d)

	// print varible
	fmt.Println(e + f)

	// print typealias
	var chinese 中文 = "中文"
	fmt.Print(chinese + "\n")

	// if 语句
	var g = 1
	// if 中可以初始化变量，如果和外面的变量名相同，则在判断语句中，外部变量就隐藏了，内部变量生效。但除了作用域，内部变量失效。
	if g = 10; g > 2 {
		fmt.Println("a > 2")
	}

	// for 循环
	// 第一种，无线循环
	fmt.Println("for 语句")
	fmt.Println("第一种，无线循环")
	var h = 1
	for {
		// 无限循环
		h++
		if h > 3 {
			break
		}
		fmt.Println(h)
	}
	// 第二种 有条件循环
	fmt.Println("第二种 有条件循环")
	for h < 5 {
		fmt.Println(h)
		h++
	}

	// 第三种，普通for
	fmt.Println("第三种，普通for")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// swith 语句不用break，自动break。如果不想break，需要显式输入fallthrough
	fmt.Println("switch 语句")
	j := 1
	switch j {
	case 0:
		fmt.Print("j = 0")
	case 1:
		fmt.Print("j = 1 \n")
		fallthrough
	default:
		fmt.Println("j ==== 1")
	}

	// 标签 LABEL
	// 两层循环，可以break任意一层
LABEL:
	for k := 0; k < 10; k++ {
		for {
			fmt.Println(k)
			break LABEL
		}
	}

	go run("go run here!")
}
