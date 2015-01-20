package main

import (
	"fmt"
)

/*
	slice简要说明：slice是array的另一种形式，本质上slice是指向array的指针。所以首先内存要有
	一整块连续的内存块，即一个array，slice规定了截取数组的开始和长度，这也就是slice英文意思“切”的
	本质。把array切开。但要提前规定slice需要截取的长度，和slice的容量capacity。如果长度在capacity
	之内，则slice的内存地址不变，如果slice的长度增加到超过了capacity，那么slice需要重新取内存地址。
	还要说明的是因为slice本质是指针，所以slice是引用类型。array是值类型。

	本示例代码的主要内容有：
	1. slice 声明方法：两种。第一种是从一个数组中直接截取一段。第二种用make函数。
	2. reslice
	3. append
	4. copy

	总结：slice相当于一个可变数组，从数组中截取一段，可以变长，可以拷贝到另一个数组中。改变slice元素，同时可以改变底层array元素。
*/

func main() {
	fmt.Println("Hello,world")

	//slice的声明。
	var noneSlice = []int{} // 空slice
	fmt.Println(noneSlice)

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)

	// 第一种声明方法，在数组中截取。前面的2说明从index=2的地方开始截取，6说明到index=6为止（不包括6）
	s1 := array[2:6]
	fmt.Println(s1)

	// 第二种声明方法，直接声明一个slice，使用make。中间参数表示长度，最后参数表示capacity。
	s2 := make([]int, 3, 6)
	fmt.Println(len(s2), cap(s2), s2)

	// reslice
	arrayByte := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	arrayByte2 := arrayByte
	s3 := arrayByte[2:6]
	s3 = append(s3, 'k', 'l') // append 向slice后面加元素。随之也修改了源数组。
	fmt.Printf("%v,%p", s3, s3)
	s4 := s3[0:8] // reslice 的序号是跟着被slice的，此例子序号应该是s3的。虽然s3长度没有8，但是容量大于等于8。
	fmt.Printf("%v,%p", s4, s4)
	s4[0] = 'z'
	fmt.Println(arrayByte)
	fmt.Println(arrayByte2) // 数组赋值是copy的，所以arrayByte2没有变化。

	s5 := []int{1, 2, 3, 4, 5}
	s6 := []int{6, 7}
	copy(s5, s6)    // 把s6中的元素拷贝到s5的对应地方去。
	fmt.Println(s5) // 输出[6 7 3 4 5]

}
