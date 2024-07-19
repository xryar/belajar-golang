package main

import "fmt"

func main() {

	//looping standar
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//looping tanpa argument
	var i = 0

	for {
		fmt.Println("Angka", i)
		i++

		if i == 5 {
			break
		}
	}

	loopingRange()

}

func loopingRange() {
	var xs = "123" //string
	for i, v := range xs {
		fmt.Println("Index = ", i, "value", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} //array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] //slice
	for _, v := range zs {
		fmt.Println("value = ", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} //map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}
}
