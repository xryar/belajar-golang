package main

import (
	"errors"
	f "fmt"
	"strconv"
	"strings"
)

func main() {
	checkNumber()
	checkName()
	checkData()
}

func checkNumber() {
	var input string
	f.Print("Type some number: ")
	f.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		f.Println(number, "is number")
	} else {
		f.Println(input, "is not number")
		f.Println(err.Error())
	}

	f.Println(" ")
}

func checkName() {
	defer catch()
	var name string
	f.Print("Type your name: ")
	f.Scanln(&name)

	if valid, err := validate(name); valid {
		f.Println("halo", name)
	} else {
		panic(err.Error())
	}

	f.Println(" ")
}

func checkData() {
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					f.Println("Panic occured on looping", each, "| message:", r)
				} else {
					f.Println("Application running perfectly")
				}
			}()

			panic("some error happen")

		}()
	}

}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		f.Println("error occured", r)
	} else {
		f.Println("Application running perfectly")
	}

	f.Println(" ")
}
