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
	Director *Director `json:"director"` // Pointer to Director. So it will have the same values that define inside the Director
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// It is a slice of movie
var movies []Movie

// Passing a pointer of the request that you will send from your Postman to this function
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Setting content type as JSON
	w.Header().Set("Content-Type", "application/json")
	// Encoding as a JSON and sending as JSON format
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Getting ID from the params
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// movies[:index]: won't exist
			// movies[index+1:]...: all other data will just append
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// Sending the remaining movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			// Sending only 1 movie
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	// While creating a movie we'll send something in the body, an entire movie
	// and we'll send it in the body in Postman
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Converting the value generated of rand to string
	movie.ID = strconv.Itoa((rand.Intn(100000000)))
	// The new movie that has come out from the body is now inside the `movies`
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// params
	params := mux.Vars(r)

	// loop over the movies, range
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie with the id that you've sent
			movies = append(movies[:index], movies[index+1:]...)

			// add a new movie - the movie that we send in the body of Postman
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
