package main

import (
	f "fmt"
	"runtime"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		time.Sleep(100 * time.Millisecond)
		f.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "hallo")
	print(5, "apa kabar")

	var input string
	f.Scanln(&input)
}
