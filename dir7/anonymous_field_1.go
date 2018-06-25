package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "computer science"}

	// 直接取出Human的成员，变相的把Human的成员引入Student中
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His speciality is ", mark.speciality)

	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)

}
