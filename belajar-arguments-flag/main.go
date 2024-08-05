package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	arg()
	flg()
	passingReference()
}

func arg() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
}

func flg() {
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}

func passingReference() {
	var data1 = flag.String("name", "anonymous", "type your name")
	fmt.Println(*data1)

	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	fmt.Println(data2)
}
