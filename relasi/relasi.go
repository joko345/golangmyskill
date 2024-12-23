package main

import (
	"fmt"
	"strings"
)

func main() {
	var getMinMax = func(n []int) (int, int) { //func dalam var. func closure
		var min, max int
		for i, value := range n {
			switch {
			case i == 0:
				max, min = value, value
			case value > max:
				max = value
			case value < min:
				min = value
			}
		}
		return min, max
	}

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

	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("rata-rata %.2f", avg)
	fmt.Println(msg)

	var numbers2 = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers2) //panggil variabe
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers2, min, max)

	var hasilFilter = filter([]string{"ini", "data"}, func(each string) bool {
		return true
	})
	fmt.Println(hasilFilter)
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

func calculate(numbers ...int) float64 { //int menjadi array int. titik tiga (func variadic)
	var total int = 0
	for _, number := range numbers {
		//	_, digunakan sebagai penerima kosong untuk mengabaikan nilai indeks yang dihasilkan dari range.
		//hanya membutuhkan value (each) dan tidak peduli dengan indeksnya.
		total += number //looping _, 0+array number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func filter(data []string, callback func(string) bool) []string { //func menerima string dan callback/mengembalikan boolean
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered { // jika lolos maka callback ke result yang bernilai array string
			result = append(result, each) //add each ke result
		}
	}
	return result
}
