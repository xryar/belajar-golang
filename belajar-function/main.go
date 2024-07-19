package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var names = []string{"John", "Wick"}
	var randomValue int

	printMessage("halo", names)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}
