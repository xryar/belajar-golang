package main

import "fmt"

func main() {

	//tipe data numerik non decimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	//tipe data numerik decimal
	var decimalNumber = 2.62

	//tipe data boolean
	var exist bool = true

	// tipe data constanta
	const pi = "22/7"

	//multi level constanta
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	fmt.Printf("bilangan Positif: %d\n", positiveNumber)
	fmt.Printf("bilangan Negatif: %d\n", negativeNumber)
	fmt.Printf("bilangan Desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan Desimal: %.3f\n", decimalNumber)
	fmt.Printf("Exist? %t \n", exist)
	fmt.Println("Pi adalah", pi)
	fmt.Println(square, isToday, numeric, floatNum)
}
