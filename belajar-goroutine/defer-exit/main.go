package main

import (
	f "fmt"
)

func main() {
	number := 3

	if number == 3 {
		f.Println("halo 1")
		//dibikin function kalau mau dimunculkan sesudah statement diatas
		func() {
			defer f.Println("halo 4")
		}()

		defer f.Println("halo 3")
	}

	f.Println("halo 2")

	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer f.Println("Terimakasih, silakan tunggu")

	if menu == "pizza" {
		f.Print("Pilihan tepat!", " ")
		f.Print("Pizza ditempat kami paling enak!", "\n")
		return
	} else if menu == "burger" {
		f.Print("Pilihan tepat!", " ")
		f.Print("Burger kami wareg cik!", "\n")
		return
	}

	f.Println("Pesanan anda:", menu)
}
