package main

import "fmt"

func main() {
	fmt.Println("hello,world!")

	var a = [2]int{1, 2}
	b := [20]int{2: 2, 19: 1} // index : value。
	c := [1]int{1}
	d := [...]int{1, 2, 3, 4, 5} // 可以自动推断个数
	e := [...]int{19: 1}         // 通过index知道数组元素个数至少有20个
	e[2] = 3
	x, y := 1, 2
	f := [...]*int{&x, &y} //指针数组
	var g *[20]int = &e    //数组指针
	fmt.Println(a, b, c, d, e, f, g)

	// 数组比较可以使用 == 或 !=，相等的条件是，数组元素个数相同，类型相同，值相同才可以。
	i := [2]int{1, 2}
	j := [2]int{1, 3}
	fmt.Println(i == j)

	// 同上的数组指针，new关键字可以创建数组指针
	k := new([10]int)
	fmt.Println(k)
	k[5] = 2 // 此方式可以修改数组某index的元素的值，也可以修改数组指针某index的值
	fmt.Println(k)

	// 多维数组
	l := [2][2]int{
		{1, 1},
		{2, 2}}

	fmt.Println(l)

	// 实例： 冒泡排序
	m := [8]int{4, 3, 7, 5, 1, 9, 6, 2}
	n := len(m)
	for o := 0; o < n; o++ {
		for p := o + 1; p < n; p++ {
			if m[o] < m[p] {
				var temp = m[o]
				m[o] = m[p]
				m[p] = temp
			}
		}
	}
	fmt.Println(m)
}
