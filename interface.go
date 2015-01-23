package main

import (
	"fmt"
)

/*
	interface 大致和OC中的protocol类似，声明一些接口。然后其他struct可以实现。
	所以配合前面的struct和function可以实现interface。

	本节的重点当然是interface的使用。还有一个知识点就是类型判断。usb.()
	还有interface类型转换。interface之间的类型转换只有范围大的可以转换为范围小的。
	下面code中USB中就包含Connect，所以USB类型的struct可以强制转换成Connect的struct
*/

type USB interface {
	Name() string
	Connect // 嵌套interface
}

type Connect interface {
	Connect()
}

//------------------------------
type connecter struct {
	name string
}

// =========================

func (c connecter) Name() string {
	return c.name
}

func (c connecter) Connect() {
	fmt.Println("connected", c.name)
}

type A struct {
}

func main() {
	pc := connecter{
		name: "pc",
	}
	pc.Connect()
	disconnect(pc)

	phone := connecter{
		name: "iPhone",
	}
	phone.Connect()
	disconnect(phone)

	// 以下两行如果添加是编译不通过的，因为A类型没有实现USB的接口
	// a := A{}
	// disconnect(a)

	// 类型转换
	// pc 转换为e,其实e还是一个pc，只不过是pc的拷贝
	e := Connect(pc)
	e.Connect()
	pc.name = "new pc"
	e.Connect() // 打印的还是pc原来的名字，证明是pc的拷贝
}

func disconnect(usb USB) {
	fmt.Println("disconnected", usb.Name())
	// 由于参数usb是一个interface 类型，所以不知道具体是什么struct类型，通过下面代码判断。
	// 代码就是interface.(struct)
	if u, ok := usb.(connecter); ok {
		// usb是connecter类型变量
		fmt.Println("I'm connecter struct", u.name)
	}

}
