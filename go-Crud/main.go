package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)             //parameter meneriman var dari URL
	for index, item := range movies { //cari id dari kumpulan slice
		if item.ID == params["id"] { //jenis parameter ID maka params menerima value ID
			movies = append(movies[:index], movies[index+1:]...) //mengambil id dari kumpulan id dan menghapusnya
			break                                                //jika id 3 maka hapus 3 dari dempetan sesudah 2 dan sebelum 4
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func findMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	for index, item := range movies {
		if item.ID == params["id"] { //mencari id
			updatedMovie.ID = params["id"] // update berdasarkan id
			movies[index] = updatedMovie
			json.NewEncoder(w).Encode(updatedMovie) //print data baru ke body
			return
		}
	}

	// Jika ID tidak ditemukan, berikan respons kosong
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "200", Title: "Movie One", Director: &Director{Firstname: "Jon", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "222", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "London"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", findMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
