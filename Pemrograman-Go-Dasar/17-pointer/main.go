package main

import "fmt"

func main() {
	// variabel bertipe pointer ditandai dengan adanya tanda asterisk
	var numberA int
	numberA = 4
	var numberB *int
	numberB =  &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberA (value) :", *numberB)
	fmt.Println("numberA (address) :", numberB)

	// Ketika salah satu variabel pointer diubah nilanya, sedang ada variabel lain
	// yang memiliki referensi memori yang sama, maka variabel lain tersebut juga akan berubah

	numberA = 5
	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberA (value) :", *numberB)
	fmt.Println("numberA (address) :", numberB)

	// Parameter pointer
	var number = 4
	fmt.Println("before : ", number) // 4

	change(&number, 10)
	fmt.Println("after :", number) // 10

}

// Parameter pointer
func change(original *int, value int) {
	*original = value
}
