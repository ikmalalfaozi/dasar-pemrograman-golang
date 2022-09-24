package main

import "fmt"

func main() {
	// deklarasi multi variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eight, ninth)
}
