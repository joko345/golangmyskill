package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //fprint jika akan ada argumen %
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path)) //%q menampilkan path setelah hello
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
