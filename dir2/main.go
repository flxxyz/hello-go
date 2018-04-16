package main

import (
	p "fmt"
	s "github.com/flxxyz/hello/dir2/show"
	h "github.com/flxxyz/hello/dir2/hide"
	"unsafe"
)

var n uint = 1

func main() {
	s.Run()
	h.Run()
	p.Println(unsafe.Sizeof(n))
}



