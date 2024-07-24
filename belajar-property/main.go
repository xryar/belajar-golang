package main

import (
	"belajar-property/library"
	f "fmt"
)

func main() {
	//pakai ini kalau ga pake function init
	// var s1 = library.Student{"ethan", 21}
	// f.Println("name  ", s1.Name)
	// f.Println("grade ", s1.Grade)

	f.Println("name  :", library.Student.Name)
	f.Println("grade :", library.Student.Grade)

	sayHello("wick")
}
