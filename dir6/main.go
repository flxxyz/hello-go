package main

import (
	"fmt"
	"time"
)

func block1() {
	goto One
	fmt.Println("code block")

	One:
		fmt.Println("one pices")
		time.Sleep(time.Second)
}

func block2() {
	for i := 0; i < 10; i++  {
		for j := 0; j < i; j++  {
			if j == 0 {
				continue
			}

			if (i % j) == 2 {
				fmt.Println(i, ",", j)
				goto Two
			}
		}
	}

	Two:
		fmt.Println("试试能跳出不")
}

func main() {
	block1()

	block2()
}
