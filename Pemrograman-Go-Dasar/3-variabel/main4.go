package main

import "fmt"

func main() {
	// deklarasi variabel menggunakan keyword new
	name := new(string)
	*name = "john wick"

	fmt.Println(name)
	fmt.Println(*name)
}
