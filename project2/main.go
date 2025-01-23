package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"//convert math/rand id into string

)


type Movie struct{
ID string `json:"id"`
Isbn string `json:"isbn"`
Title string `json:"title`
Director *Director `json:"director"`
}

type Director struct{
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	for index, item:= range movies {
		if item.ID==params["id"]{
			//igm=nore that movie and append all other movies
			movies=append(movies[:index], movies[index+1:]...)
			break
		}
	}

}


func getMovie(w http.ResponseWriter , r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for _,item :=range movies {
		if(item.ID==params["id"]){
			json.NewEncoder(w).Encode(item)
			return;
		}
	}
}

func createMovie(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	// params=mux.Vars(r);
	var movie Movie;
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(1000000))
	movies=append(movies, movie);
}

func updateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)

	for index, item := range movies {
		if(item.ID==params["id"]){
			movies=append(movies[:index], movies[index+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies, movie)
		}

	} 
}



func main(){
	r:=mux.NewRouter()


	movies= append(movies, Movie{ID:"1", Isbn:"448743", Title:"Movie One", Director:&Director{Firstname:"John", Lastname:"Doe"}})
	movies= append(movies, Movie{ID:"2", Isbn:"448744", Title:"Movie Two", Director:&Director{Firstname:"Johnny", Lastname:"Done"}})

	// r.Handle("/", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w, "hello");
	// })
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/moies/{id}",getMovie).Methods("GET")

	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}