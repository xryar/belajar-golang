package main

import (
	f "fmt"
	"runtime"
)

// channel sebagai param
func printMessage(what chan string) {
	f.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = f.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var messages1 = <-messages
	f.Println(messages1)

	var messages2 = <-messages
	f.Println(messages2)

	var messages3 = <-messages
	f.Println(messages3)
	f.Println(" ")

	// penerapan channel sebagai param
	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data1 = f.Sprintf("hello %s", who)
			messages <- data1
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
