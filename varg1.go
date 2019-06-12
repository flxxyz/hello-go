package main

import "fmt"

func arg(argv ...interface{}) {
	for _, a := range argv {
		switch a.(type) {
		case int:
			fmt.Println("int, ", a)
		case string:
			fmt.Println("string, ", a)
		case int64:
			fmt.Println("int64, ", a)
		}
	}
}

func main() {
	arg(1, int64(666), 2333, "fuck")
}
