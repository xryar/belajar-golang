package library

import "fmt"

//intinya kalau huruf awalnya kapital berarti dia public dan bisa dipanggil di main

//diubah jadi AnonStruct dulu kalo mau pakai function init
var Student = struct {
	Name  string
	Grade int
}{}

func SayHello(name string) {
	fmt.Println("Hello")
	introduce(name)
}

//contohnya ini, function introduce tidak bisa di panggil di main karena huruf awalnya kecil
func introduce(name string) {
	fmt.Println("nama saya", name)
}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}
