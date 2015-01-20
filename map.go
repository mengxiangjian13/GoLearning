package main

import (
	"fmt"
	"sort"
)

/*
	概述：
	map 类似OC中的字典，key-value型。与OC中字典的key必须是遵从NSCoping协议的类似，
	Go中的map的key必须是可以用==比较的，比如int，string。不能是slice，map，函数。
	而value可以是任何类型。
	map的查找速率比slice，array以index查找要慢，但总比线性查找要快。

	本代码内容：
	map的初始化。和array slice类似，可以通过类型加{}初始化，也可以make
	给map的增、删、改、查操作
	多返回值
	* for range （enum || for each）
	排序 sort
*/

func main() {
	a := make(map[int]string) // 初始化map
	a[0] = "OK"               // 给某个key赋值value
	b := a[0]                 // 取某个key的value。如果没有这样的key，value为空
	fmt.Println(a, b)
	delete(a, 0) // 删除某个键值对
	fmt.Println(a)

	// 复杂一些的map，比如map的value还是个map
	c := make(map[int]map[int]string) // 第一级初始化
	c[1] = make(map[int]string)       // 内层第二级初始化
	c[1][1] = "GOOD"
	fmt.Println(c)

	//其实c[]这种取某个key的value的方式是有两个返回值的。第一个是value，第二个是是否有值的bool类型变量。
	d, ok := c[0][1]
	fmt.Println(d, ok) // 输出 _ false 即没有值

	// for range
	// slice map
	e := []int{1, 2, 3, 4, 5, 6}
	for i, v := range e {
		// i 为 index，v 为 value。且value是拷贝。如果修改某个v，其实slice中的元素不会受影响
		// 如果更改e[i]，那就会影响到e了
		// 如果for range 一个map，那么就没有index，换成了key。
		fmt.Println(i, v)
	}

	// sort
	f := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	g := make([]int, len(f))
	i := 0
	for k := range f {
		g[i] = k // 取出所有key放进slice中
		i++
	}
	fmt.Println(g) // slice是中的map key是无序的
	sort.Ints(g)   // 给g中的元素排序。间接证明slice是引用类型
	fmt.Println(g) // 排序过的
}
