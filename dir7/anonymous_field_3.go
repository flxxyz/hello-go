package main

import "fmt"

type Human struct {
	name string
	age  int
	phone string
}

type Employee struct {
	Human
	speciality string
	phone string
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-xxxx"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 还是通过对象链访问成员
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
