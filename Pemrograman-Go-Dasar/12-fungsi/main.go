package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"john", "wick"}
	printMessage("halo", names)

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

}

// fungsi tanpa kembalian
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// fungsi dengan kembalian
func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min + 1) + min
	return value
}
