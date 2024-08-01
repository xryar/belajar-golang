package main

import (
	f "fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	f.Println("time elapse in seconds:", duration.Seconds())
	f.Println("time elapse in minutes:", duration.Minutes())
	f.Println("time elapse in hours:", duration.Hours())

	// calculateDurationObject()
}

func calculateDurationObject() {
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration := t2.Sub(t1)

	f.Println("time elapse in Seconds:", duration.Seconds())
	f.Println("time elapse in Minutes:", duration.Minutes())
	f.Println("time elapse in Hours:", duration.Hours())
}
