package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}


// fungsi dengan predufined return value
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2,2)
	circumference = math.Pi * d

	return
}

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t\t: %.2f \n", circumference)

	var area2, circumference2 = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area2)
	fmt.Printf("keliling lingkaran\t\t: %.2f \n", circumference2)
}
