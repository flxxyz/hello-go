package test

import (
	"fmt"
	"reflect"
)

var a, b, c = 1, 2, 3.5

var House = "大房子"

func Run() {
	num()
	fmt.Println("外部可访问的公用方法")
}

func run() {
	num()
	fmt.Println("只允许内部访问的私有方法")
}

func num() {
	// 在函数体内可以使用 := 来给多个变量赋值
	d, e, f := 3, 2, 1.2

	fmt.Println(d, e, f)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
}
