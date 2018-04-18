package main

import (
	"fmt"
)

// 跳值
const (
	a1 = iota
	a2 = iota
	_ = iota
	a3 = iota
	_ = iota
	a4 = iota
)

// 插值
const (
	b1 = iota
	b2 = iota
	b3 = 3.1415926
	b4 = iota
)

const (
	c1 = iota
	c2 = iota
	_
	c3 = iota
	c4 = iota
)

const (
	d1 = iota + 7
	d2
	d3 = iota * 2
	d4
)

const (
	e1, e2 = iota, iota + 5
	e3, e4
	e5 = iota
)

func main() {
	fmt.Println(a1, a2, a3, a4)

	fmt.Println(b1, b2, b3, b4)

	fmt.Println(c1, c2, c3, c4)

	fmt.Println(d1, d2, d3, d4)

	fmt.Println(e1, e2, e3, e4, e5)
}
