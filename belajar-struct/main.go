package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

/* nested struct kaya gini

	type student struct {
		person struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
}
*/

type student struct {
	person
	age   int
	grade int
}

func main() {
	// ngisi nilai sub-struct
	// var p1 = person{name: "wick", age: 21}
	// var s1 = student{person: p1, grade: 2}

	var s1 = student{}
	s1.name = "Jhon wick"
	s1.grade = 2
	s1.age = 21
	s1.person.age = 22

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
	fmt.Println(s1.grade)
	fmt.Println(" ")

	anonStruct()
	combineWithSlice()
}

func anonStruct() {
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"wick", 2}
	s1.grade = 2

	fmt.Println("name  :", s1.person.name)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("grade :", s1.grade)
	fmt.Println(" ")
}

func combineWithSlice() {
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	fmt.Println(" ")
}
