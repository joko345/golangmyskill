package main

import (
	"fmt"
	"strings"
)

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

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	var chicken map[string]int
	//variable tipe map key bersifat string value bersifat int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"])
	//merujuk ke variable chicked dengan key kata kunci januari yang berisi value 50

	type student struct {
		name  string
		grade int
	}

	var s1 student
	s1.name = "my own"
	s1.grade = 2
	fmt.Println("nama", s1.name)
	fmt.Println("grade", s1.grade)

	var makeFruits = make([]string, 2)
	makeFruits[0] = "apple"
	makeFruits[1] = "manggo"
	fmt.Println(makeFruits)

	var names = []string{"muy", "new"} //var slice jumlah bertambah sesuai kebuuhan
	names[0] = "berubah"
	names = append(names, "ketiga")
	printMsg("halo", names)
	personName, personAge := getPersonInfo()
	fmt.Println(personAge, personName)
}

func printMsg(message string, arr []string) {
	var nameString = strings.Join(arr, "") //fungsi join names
	fmt.Println(message, nameString)
}

func getPersonInfo() (string, int) { //return dua nilai
	name2 := "Alice"
	age2 := 25
	return name2, age2
}
