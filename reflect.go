package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var a float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(a))

	// valueOf类型有许多方法，type kind 取类型。int float取值。
	v := reflect.ValueOf(a)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	// 修改值，传递地址可以修改本身
	var x float64 = 3.4
	va := reflect.ValueOf(&x)
	fmt.Println("va settablility of va:", va.CanSet())
	vae := va.Elem()
	fmt.Println("settability of vae:", vae.CanSet())
	vae.SetFloat(6.8)
	fmt.Println(vae.Interface())
	fmt.Println(x)

	// 结构的反射，反射出结构内的变量
	t := T{203, "m123"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i) // 通过索引获取对应成员
		fmt.Println(i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
