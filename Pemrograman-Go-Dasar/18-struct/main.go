package main

import "fmt"

type student struct {
	name string
	grade int
}

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	// Inisialisasi object struct
	var s2 = student{}
	s2.name = "wick"
	s2.grade = 2

	var s3 = student{"ethan", 2}

	var s4 = student{name: "jason"}

	fmt.Println("studnet 2 :", s2.name)
	fmt.Println("studnet 3 :", s3.name)
	fmt.Println("studnet 4 :", s4.name)

	// Variabel objek pointer
	var s5 = student{name : "wick", grade: 2}

	var s6 *student = &s5
	fmt.Println("student 5, name :", s5.name)
	fmt.Println("student 6, name :", s6.name)

	s5.name = "ethan"
	fmt.Println("student 5, name :", s5.name)
	fmt.Println("student 6, name :", s6.name)
}
