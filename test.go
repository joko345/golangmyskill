package main

/* */

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	defer recov() //defer tetap berjalan meski ada panic
	// var last string = "new"
	// const valid = "main"
	var positif int = 89
	var negatif = -123
	var decima = 2.2
	var test *int = &positif //&positif dengan amperstand & akan menampilkan alamat memori
	*test = 2                //dengan asterisk maka var positif valuenya akan ikut berubah

	// fmt.Print("halo 2 ", last)
	fmt.Println("positif = ", positif, decima, *test) //agar muncul value pakai asterix *
	fmt.Print("negatif = ", negatif)
	defer fmt.Println("halo") //defer dijalankan diakhir blok program
	fmt.Println("selamat datang")
	// os.Exit(1) exit paksa

	fmt.Print("type number : ")
	var input string
	value, err := testError(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("converted =", value)

	}
}

var testError = func(input string) (int, error) { //func dalam var. func closure
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input) //input menjadi int
	if err == nil {
		fmt.Println(number, "is number")
		return number, nil
	} else {
		if strings.TrimSpace(input) == "" { //mengecek inputan yang awalnya "string"
			return 0, errors.New("cant be empty")
		} else {
			fmt.Printf(" %q is not number\n", input)                                   //%q menampilkan input dalam petik pada printF
			panic(err.Error())                                                         //panic membuat kode setelahnya tidak dijalankan
			return 0, fmt.Errorf("Konversi gagal: '%s' bukan angka yang valid", input) // %s menampilkan input dalam petik pada errorF
			// return 0, errors.New(fmt.Sprintf("Konversi gagal: '%s' bukan angka yang valid", input))

		}
	}
}

func recov() {
	if r := recover(); r != nil {
		fmt.Println("panic muncul:", r)
	} else {
		fmt.Println("app running smooth")
	}
}
