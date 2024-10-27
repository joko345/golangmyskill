package main

import "fmt"

func main() {
	var point = 8
	if point == 10 && point < 20 {
		fmt.Println("lulus sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("kurang")
	} else {
		fmt.Println("tidak lulus")
	}

	var tukar = 1
	switch tukar {
	case 8:
		fmt.Println("lulus sempurna")
	case 5:
		fmt.Println("tidak lulus")
	default:
		fmt.Println("kurang")
	}
	var equal = (2 == 2)
	fmt.Print("nilai \n", equal)
}
