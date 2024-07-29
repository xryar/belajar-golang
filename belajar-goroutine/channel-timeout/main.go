package main

import (
	f "fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go sendData(messages)
	retriveData(messages)
}

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func retriveData(ch chan int) {
loop:
	for {
		select {
		case data := <-ch:
			f.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			f.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}
