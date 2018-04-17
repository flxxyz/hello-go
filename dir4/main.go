package main

import "fmt"

const PARAM1 int = 1
const PARAM2 = 2
const (
	PARAM3 string = "大佬鼠"
	PARAM4 = "小脑斧"
)


func main() {
	fmt.Println(PARAM1)
	fmt.Println(PARAM2)

	fmt.Println(PARAM3)
	fmt.Println(PARAM4)

	// len()对string类型有效
	fmt.Println(len(PARAM4))
}
