package main

import "fmt"

func main() {
	// Operator aritmatika +, -, *, /, %
	var value = (((2 + 6) % 3) * 4 - 2) / 3

	fmt.Println(value)

	// Operator perbandingan ==, !=, <, <=, >, >=
	var isEqual = (value == 2)

	fmt.Println(isEqual)

	// operator logika &&, ||, !
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

}
