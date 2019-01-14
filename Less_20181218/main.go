package main

import (
	"fmt"
	"math"
	"runtime"
)


const Uint1024 = 1024

func main()  {
	//fmt.Printf("%v", uintConver())

	// 1.4142135623730951
	//fmt.Printf("%v\n", sqrt(2))
	//fmt.Printf("%v\n", math.Sqrt(2))


	//queryOS()

	//fmt.Println(time.Now())


	//fmt.Println("counting")
	//deferInTurn()
	//fmt.Println("done")

	slice()
}

func uintConver() int {
	return int(math.Pow(Uint1024, 3))
}

func sqrt(x float64) float64 {
	num := x / 2.0
	z := 0.0
	count := 1

	for math.Abs(num - z) > 0.00000001  {
		count += 1
		z = num
		num = ((1.0 / 2.0) * num) + (x * 1.0) / (2.0 * z)
		//fmt.Println(z)
	}

	//x = z - ((z * z) - x) / (2 * x)
	//
	fmt.Println("count=", count)

	return z
}

func queryOS()  {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	case "freebsd":
		fmt.Println("OpenBSD.")
	case "openbsd":
		fmt.Println("FreeBSD.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Printf("OS=%s.\n", runtime.GOOS)
}

func deferInTurn() {
	fmt.Println("deferInTurn() start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("deferInTurn() end")
}

func slice()  {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:1]
	fmt.Println(s)

	s = s[0:]
	fmt.Println(s)
	fmt.Println(s[:])
}
