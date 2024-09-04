package main

import (
	"belajar-generics/dao"
	"fmt"
)

func main() {
	var m1 dao.UserModel[int]
	m1.Name = "Noval"
	m1.Scores = []int{1, 2, 3}
	fmt.Println("scores:", m1.Scores)

	var m2 dao.UserModel[float64]
	m2.Name = "Noval"
	m2.SetScoresB([]float64{10, 11})
	fmt.Println("scores:", m2.Scores)

	// total1 := Sum([]int{1, 2, 3, 4, 5})
	// fmt.Println("total:", total1)

	// total2 := Sum([]float32{2.5, 7.2})
	// fmt.Println("total:", total2)

	// total3 := Sum([]float64{1.23, 6.33, 12.6})
	// fmt.Println("total:", total3)

	// ints := map[string]int64{"first": 34, "second": 12}
	// floats := map[string]float64{"first": 35.98, "second": 26.99}

	// fmt.Printf("Generics Sums with Constraint : %v and %v\n", dao.SumNumbers3(ints), dao.SumNumbers3(floats))

}
