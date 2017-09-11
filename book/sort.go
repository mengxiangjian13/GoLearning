package main

import (
	"fmt"
	"sort"
	"strings"
)

type foldedString []string

func main() {
	a := []string{"mengxiangjian", "Guozhenyan", "yangxiangming", "liuyuqing", "Xinghongbin", "Yangchong", "xuxiaolong"}
	// sort
	sort.Strings(a)
	fmt.Print("区分大小写\n")
	fmt.Print(a)
	fmt.Print("\n")
	sort.Sort(foldedString(a))
	fmt.Print("不区分大小写\n")
	fmt.Print(a)

	// search 自定义查找要使用已经排序的slice进行查找
	target := "MENGXIANGJIAN"
	index := sort.Search(len(a), func(i int) bool {
		return strings.ToLower(a[i]) >= strings.ToLower(target)
	})

	if index < len(a) && strings.EqualFold(target, a[index]) {
		fmt.Print("\n不区分大小写\n查找目标为 " + target + "\n查找结果为 " + a[index])
	}

}

func (f foldedString) Len() int {
	return len(f)
}

func (f foldedString) Less(i, j int) bool {
	return (strings.ToLower(f[i]) < strings.ToLower(f[j]))
}

func (f foldedString) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
