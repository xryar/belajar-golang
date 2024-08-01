package main

import (
	f "fmt"
	"os"
	"time"
)

func timer(timeOut int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeOut)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeOut int, ch <-chan bool) {
	<-ch
	f.Println("\ntime out! no answer more than", timeOut, "seconds")
	os.Exit(0)
}

func main() {
	var timeOut = 5
	var ch = make(chan bool)

	go timer(timeOut, ch)
	go watcher(timeOut, ch)

	var input string
	f.Print("what is 725/5 ? ")
	f.Scan(&input)

	if input == "29" {
		f.Println("the answer is right!")
	} else {
		f.Println("the answer is wrong!")
	}
}
