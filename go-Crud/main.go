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
	ID       string    `json:"id"` //ID akan dikenal sebagai id dalam var
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json:"director"` //point ke type director dan dipanggil dengan &Director
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	//send request dan response =
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies { //dalam js for each
		if item.ID == params["id"] { //id untuk delete
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //setelah delete view allMovie
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
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) //setelah input send data movie baru ke body lalu decode
	movie.ID = strconv.Itoa(rand.Intn(100))    //id movie akan di random dan convert string
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "200", Title: "Movie One", Director: &Director{Firstname: "Jon", lastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "222", Title: "Movie Two", Director: &Director{Firstname: "Steve", lastName: "London"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", findMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("movies", updateMovie).Methods("PUT")
	r.HandleFunc("movies", deleteMovie).Methods("DELETE")

	fmt.Printf("starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
