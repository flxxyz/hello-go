package main

import "fmt"

func main() {
	var n int = 5

	culc := func() func() {
		return func() {
			fmt.Printf("new param{n} is %d", n)
			fmt.Println()
		}
	}()

	culc()

	n += 5

	culc()
}
