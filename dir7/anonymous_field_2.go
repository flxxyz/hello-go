package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
	phone  string
}

type Student struct {
	Human
	Skills
	int
	speciality string
	phone      string
}

func main() {
	jane := Student{Human: Human{"Jane", 35, 100, "10086"}, speciality: "Biology", phone: "10000"}
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	fmt.Println("Her phone is ", jane.phone)
	fmt.Println("Her Human.phone is ", jane.Human.phone)

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("she acquired two new ones")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	jane.int = 3
	fmt.Println("Her perferred number is ", jane.int)

}
