package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		bisa juga ditulis kaya gini untuk ngisi paramaternya
		var numbers = []int{2, 3, 4, 5, 4, 3, 3, 5, 5, 3}
		var avg = calculate(numbers...)
	*/
	var avg = calculate(2, 3, 4, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)

	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg

}

// bisa juga dikombinasikan dengan param biasa dan variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, My name is: %s\n", name)
	fmt.Printf("My Hobbies are: %s\n", hobbiesAsString)
}
