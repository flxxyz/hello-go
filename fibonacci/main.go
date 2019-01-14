package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 50; i++ {
		fmt.Printf("fibonacci(%d) = %d \n", i, fibonacci(i))
	}

	//time.Sleep(time.Second * 10)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return f(n, 2) + f(n, 1)
}

func f(n int, other int) int {
	return fibonacci(n - other)
}
