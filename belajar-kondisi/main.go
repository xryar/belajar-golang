package main

import "fmt"

func main() {

	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda %d\n", point)
	}

	//variabel temporary pada if statement
	var nilai = 8840.0

	if percent := nilai / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent > 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	//switch case
	//bisa banyak kondisi
	var pointSwitch = 6

	switch pointSwitch {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
		//kaya break lah
		fallthrough
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	var bigPoint = 10

	if bigPoint > 7 {
		switch bigPoint {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if bigPoint == 5 {
			fmt.Println("not bad")
		} else if bigPoint == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
