package main

import (
	"fmt"
)

/*
	概述：函数是语言必须掌握的内容。
	Go语言的函数，不支持嵌套，重载，默认函数。
	支持不定长度变参，多返回值，匿名函数，闭包。

	代码内容：
	1. 函数的声明，返回值以及各种函数特性（多返回值，重载，闭包）
	2. defer函数。函数在完成之后调用，也可以叫析构函数。
	3. 异常捕捉机制。panic recover。panic可以在任何时候触发，但recover只能在defer里触发。
		panic可以停止程序，需要用recover来恢复。
*/

func main() {
	a := sum(1, 2)
	fmt.Println(a)

	multiParam("OK", 1, 2, 3)

	b := 1
	pointerParam(&b)
	fmt.Println(b)

	c, d, e := multiReturn(8)
	fmt.Println(c, d, e)

	// 匿名函数
	f := func() {
		fmt.Print("unnamed function\n")
	}
	f()

	g := 100
	// 闭包。我的理解闭包就是带有包外变量的匿名函数。
	closure := func(value int) {
		fmt.Println(g + value)
		g = 200 // 明确一点闭包里面的 包外变量 是引用型的。所以闭包执行完了，g就变成200了。
	}

	closure(50)
	fmt.Println(g)

	deferTest()

	stop()
}

// 基本格式 func <函数名> (参数 参数类型)(返回值 返回值类型)
/*
	如下参数a、b都是int型，就可以只写一个int。
*/
func sum(a, b int) int {
	return a + b
}

// 不定长参数，意思是参数个数不一定。用...表示。不定长参数只能作为最后一个参数。
// 不定长参数用一个参数表示，这个参数在函数作用域内作为slice。
func multiParam(a string, b ...int) {
	fmt.Println(a, b, len(b))
}

// 参数传递指针。
// 函数的参数传递都是拷贝，普通值拷贝不会影响外部变量，而指针的拷贝是会印象外部棉量的。
func pointerParam(a *int) {
	*a++
}

// 多返回值
func multiReturn(a int) (int, int, int) {
	b := a / 3
	return b, b, b
}

//
func deferTest() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // defer在注册的时候，i是一个拷贝
	}

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) // defer 由于在函数结束时候调用。defer调用闭包，闭包的变量是引用型的。所以i是循环最后的值。
		}()
	}
}

func stop() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("app recover")
		}
	}()
	panic("app stop")
}
