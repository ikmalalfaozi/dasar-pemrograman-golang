package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(2,4,3,5,4,3,3,5,5,3)
	var msg = fmt.Sprintf("Rata-rata: %.2f", avg)
	fmt.Println(msg)

	// pengisian parameter fungsi variadic menggunakan data slice
	var numbers = []int{2,4,3,5,4,3,3,5,5,3}
	var avg2 = calculate(numbers...)
	var msg2 = fmt.Sprintf("Rata-rata: %.2f", avg2)
	fmt.Println(msg2)

	// fungsi dengan parameter biasa dan variadic
	yourHobbies("wick", "sleeping", "eating")

	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}

// fungsi dengan parameter variadic
func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// fungsi dengan parameter biasa dan variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
