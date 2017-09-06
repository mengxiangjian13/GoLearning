package main

import "fmt"

/*
	struct 简介：
	struct就是结构体，由于Go和C语言类似，所以struct很重要。
	Go没有Class，struct也没有继承，但不能说Go完全不是面向对象的语言，他有很多靠近面向对象的用法。
	struct基本可以代替Class

	代码内容：
	1. struct的声明。值类型。
	2. struct靠近Class
	3. struct的引用类型
	4. 匿名struct
	5. 模拟继承
*/

type human struct {
	gender int
}

type person struct {
	Name  string
	Age   int
	human // 匿名参数，默认名字同struct名
}

type car struct {
	human
	Name string
}

func main() {
	// 声明一个person类型变量，设置属性。
	me := person{}
	me.Name = "mengxiangjian"
	me.Age = 26
	me.gender = 1
	fmt.Println(me)

	// 可以在声明struct类型变量的时候直接设置字面量属性。
	she := person{
		Name:   "guozhenyan",
		Age:    26,
		gender: 2,
	}
	fmt.Println(she)

	// sturct是值类型赋值或者传递函数参数后，对拷贝操作不会影响自身。
	// 但Go做了补充，让struct可以更靠近OC的class，变成引用类型。
	// 在初始化的时候在类型前面加&取内存地址，就变成了引用类型。赋值属性的时候不用*bmw，更像Class了。
	bmw := &car{
		Name: "bmw",
	}
	fmt.Println(bmw.Name)
	changeCar(bmw)
	fmt.Println(bmw.Name)
	bmw.changeCarName("mini")

	// 匿名struct
	apple := &struct {
		Name string
	}{
		Name: "Apple",
	}
	fmt.Println(apple)

	santana := &car{
		Name: "santana",
	}
	santana.gender = 1 // santana.human.gender赋值也是对的。
	fmt.Println(santana)
}

func (c *car) changeCarName(name string) {
	c.Name = name
}

func changeCar(a *car) {
	// 不需要*a.Name = "benz"
	a.Name = "benz"
}
