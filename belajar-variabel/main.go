package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	//deklarasi tanpa tipe data
	firstFood := "Kebab"

	//deklarasi multi variabel
	first, second, third := "satu", "dua", "tiga"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	//deklarasi dengan keyword new
	name := new(string)

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println(firstFood)
	fmt.Println(first, second, third)
	fmt.Println(say, one, isFriday, twoPointTwo)
	fmt.Println(name)
	fmt.Println(*name)

}
