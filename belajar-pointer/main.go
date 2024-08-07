package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println(" ")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println(" ")

	var numberC = 4
	fmt.Println("before : ", numberC)

	change(&numberC, 10)
	fmt.Println("after : ", numberC)

}

// bisa juga buat diparameter, contohnya
func change(original *int, value int) {
	*original = value
}
