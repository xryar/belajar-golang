package main

import (
	f "fmt"
	"time"
)

func main() {
	f.Println("start")
	time.Sleep(time.Second * 4)
	f.Println("after 4 seconds")
	f.Println(" ")

	newTimer()
	f.Println(" ")
	afterFunc()
	f.Println(" ")
	timeAfter()
	f.Println(" ")
	timeTicker()
}

func newTimer() {
	var timer = time.NewTimer(4 * time.Second)
	f.Println("New Timer")

	f.Println("start")
	<-timer.C
	f.Println("finish")
}

func afterFunc() {
	var ch = make(chan bool)

	f.Println("After func")
	time.AfterFunc(4*time.Second, func() {
		f.Println("expired")
		ch <- true
	})

	f.Println("start")
	<-ch
	f.Println("finish")
}

func timeAfter() {
	f.Println("Time After")
	<-time.After(4 * time.Second)
	f.Println("expired")
}

func timeTicker() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			f.Println("hello!!", t)
		}
	}
}
