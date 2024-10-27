package main

/* */

import "fmt"

func main() {
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
}
