package main

import "fmt"

func main() {
	// perulangan for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka",i)
	}
	fmt.Println()

	// perulangan  for dengan argumen hanya kondisi
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}
	fmt.Println()

	// perulangan for tanpa argumen
	i = 0
	for {
		fmt.Println("Angka", i)

		i++;
		if i == 5 {
			break
		}
	}
	fmt.Println()

	// perulangan menggunakan keyword break and continue
	for i := 0; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}
		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
	fmt.Println()

	// pemanfaatan label dalam perulangan
	outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Println("matriks [",i,"][",j,"]")
		}
	}
}
