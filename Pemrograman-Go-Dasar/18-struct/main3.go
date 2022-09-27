package main

import "fmt"

type person struct {
	name string
	age int
}

// Embedded struct dengan nama property yang sama
type student struct {
	person
	age int
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21 // age of student
	s1.person.age = 22 // age of person

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)

	// anonymous struct
	var s2 = struct {
		person
		grade int
	}{}
	s2.person = person{"wick", 21}
	s2.grade = 2

	fmt.Println("name :", s2.person.name)
	fmt.Println("age :", s2.person.age)
	fmt.Println("grade :", s2.grade)
}
