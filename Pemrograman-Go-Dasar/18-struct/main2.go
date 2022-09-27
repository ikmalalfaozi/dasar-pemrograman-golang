package main

import "fmt"

type person struct {
	name string
	age int
}

// embedded struct
type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("age :", s1.person.age)
	fmt.Println("grade:", s1.grade)

	// Pengisian sub-struct
	var p2 = person{name: "wick", age: 21}
	var s2 = student{person: p2, grade: 2}
	// var s2 = student{person: person{"wick", 2}, grade: 2}

	fmt.Println("name :", s2.name)
	fmt.Println("age :", s2.age)
	fmt.Println("grade :", s2.grade)
}
