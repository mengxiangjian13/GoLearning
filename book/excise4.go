package main

import (
	"fmt"
)

func main() {
	sliceA := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	sliceB := uniqueInts(sliceA)
	fmt.Print(sliceB)
	fmt.Print("\n")
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	flattenSlice := flatten(irregularMatrix)
	fmt.Print(flattenSlice)
	fmt.Print("\n")
	dSlice := make2D(flattenSlice, 6)
	fmt.Print(dSlice)
}

func uniqueInts(s []int) []int {
	r := []int{}
	for _, v := range s {
		hasInt := false
		for _, vv := range r {
			if v == vv {
				hasInt = true
				break
			}
		}
		if hasInt == false {
			r = append(r, v)
		}
	}
	return r
}

func flatten(doubleSlice [][]int) []int {
	r := []int{}
	for _, v := range doubleSlice {
		r = append(r, v[0:len(v)]...)
	}
	return r
}

func make2D(s []int, column int) [][]int {
	count := len(s) / column
	remain := len(s) % column
	if remain != 0 {
		count = count + 1
	}
	r := make([][]int, count)
	index := 0
	for i := 0; i < count; i++ {
		r[i] = make([]int, column)
		for j := 0; j < column; j++ {
			if index < len(s) {
				r[i][j] = s[index]
			}
			index += 1
		}
	}
	return r
}
