package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succes")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}
func haloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/halo" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //deklarasi fungsi untuk router
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/halo", haloHandler)

	fmt.Println("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { //err deklarasi server dan listen jika ada error munculkan error
		log.Fatal(err)
	}
}
