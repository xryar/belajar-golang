package main

import (
	"fmt"
	"math"
)

func main() {

	var diameter float64 = 15
	var area, circumference = calculated(diameter)

	fmt.Printf("luas lingkaran \t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran \t\t: %.2f \n", circumference)

}

func calculated(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	//hitung keliling
	var circumference = math.Pi * d

	return area, circumference
}
