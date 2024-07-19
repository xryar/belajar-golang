package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// inisiasi nilai awal array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah semua element \t\t", len(fruits))
	fmt.Println("Isi semua element\t", fruits)
}
