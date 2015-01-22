package main

import (
	"fmt"
)

/*
	method 简介：
	method其实应该是Class中的方法，Go没有Class，但Go巧妙的做到了。

	代码内容有：见下方代码前注释

	额外说明：属性，方法，函数的大小写区分可见性的概念，是基于包的范围的。
	包外只能看到大写字母开头的属性，方法，函数等。不管大小写对于包内都是可见的。
*/

type person struct {
	Name string
	Age  int
}

// Go使用function来实现method，只是多了一个receiver
func (a person) ChangeName() {
	a.Name = "guozhenyan" // 修改的只是拷贝
}

// 和函数一样，receiver也可以传递指针。
func (a *person) PointerChangeName() {
	a.Name = "guozhenyan"
}

// 类型alias
type customInt int

func (a *customInt) increase() {
	*a = *a + 100
}

func main() {
	me := person{}
	me.Name = "mengxiangjian"
	me.Age = 26
	me.ChangeName()
	fmt.Println(me.Name)
	me.PointerChangeName()
	fmt.Println(me.Name)
	// method 特殊调用方式，类似类方法。method expression
	// 格式：receiver类型.method(receiver变量)
	person.ChangeName(me)
	(*person).PointerChangeName(&me)

	// 类型alias
	var a customInt
	a.increase()
	(*customInt).increase(&a)
	fmt.Println(a)
}
