package main

import (
	"fmt"
)

func useIf() {
	a := 1
	if 2 > a {
		fmt.Println(2)
	}else {
		if true {
			fmt.Println(true)
		}else {
			fmt.Println(false)
		}
	}
}

func useSwitch()  {
	// 简单的学习interface使用

	var i interface{}

	i = '阿'
	switch i.(type) {
	case int:
		fmt.Println("整形")
	case string:
		fmt.Println("字符串")
	case bool:
		fmt.Println("布尔")
	case rune:
		fmt.Println("字面量")
	default:
		fmt.Println(false)
	}
}

func useFor()  {
	//time.Sleep(5 * time.Second);

	sum := 0
	for a := 1; a <= 100; a++  {
		sum += a
	}
	fmt.Println(sum)

	a := []string{"1", "2", "3"}
	for k, v := range a {
		fmt.Println("key=", k , ", value=", v)
	}
}

func main() {
	useIf()

	useSwitch()

	useFor()
}
