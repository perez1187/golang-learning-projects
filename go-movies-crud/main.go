// https://www.youtube.com/watch?v=TkbhQQS3m_o&list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&index=2

// 28:15

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

// we dont use db, so we use struct
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // * is a pointer, every movie has one director
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie //slice

// function for routes
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// now we want to encode w, and append to movies slice
	json.NewEncoder(w).Encode(movies) // this is something that we return ?
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// we dont use index, so we use _
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

	// we sent in postman entrire body, so we want to decode  and put to movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // save json to movie

	// we create a random id number, and convert to string
	movie.ID = strconv.Itoa(rand.Intn(100000000))

	// and add to db/struct
	movies = append(movies, movie)

	// we return movie
	json.NewEncoder(w).Encode(movie)
}

// in update movie, we first delete old movie, and create new one
// it is not something what we do when we work with db
func updateMovie(w http.ResponseWriter, r *http.Request) {

	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// we get access to params
	params := mux.Vars(r)
	// we loop over the movies
	// delete the movie with the id that you have sent
	// add a new movie
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			// we sent in postman entrire body, so we want to decode  and put to movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // save json to movie

			// we create a random id number, and convert to string
			movie.ID = strconv.Itoa(rand.Intn(100000000))

			// and add to db/struct
			movies = append(movies, movie)

			// we return movie
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	// manaully add a couple of movies
	// & address
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438227", Title: "Movie two", Director: &Director{Firstname: "Mark", Lastname: "Asdf"}})

	// we define routes
	r.HandleFunc("/movies", getMovies).Methods("GET") //getMovies will be a funtion
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	// create log if server doesnt start
	log.Fatal(http.ListenAndServe(":8000", r)) // :8000 because localhost
}
