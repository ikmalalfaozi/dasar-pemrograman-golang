package main

import "fmt"

func main() {
	// tipe data numerik non-desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1234523644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// tipe data numerik desimal
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// tipe data boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// tipe data string
	var message string = "halo"
	fmt.Printf("message: %s \n", message)
}
