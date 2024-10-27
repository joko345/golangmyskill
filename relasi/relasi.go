package main

import "fmt"

func main() {
	var point = 8
	if point == 10 {
		fmt.Println("lulus sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("kurang")
	} else {
		fmt.Println("tidak lulus")
	}
	var equal = (2 == 2)
	fmt.Print("nilai \n", equal)
}
