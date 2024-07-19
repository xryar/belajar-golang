package main

import "fmt"

func main() {

	//operasi matematika
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	//operasi logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("Left && Right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("Left || Right \t(%t) \n", leftOrRight)

	var leftReserve = !left
	fmt.Printf("Left \t\t(%t) \n", leftReserve)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}
