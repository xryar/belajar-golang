package main

import (
	f "fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)

	go func() {
		for {
			i := <-messages
			f.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		f.Println("send data", i)
		messages <- i
	}

	time.Sleep(1 * time.Millisecond)
}
