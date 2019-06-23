package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var people []Person

// Get all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// Get required person
func getRequiredPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func main() {
	people = append(people, Person{"1", "Ghalib", "Ahmed"})
	people = append(people, Person{"2", "Ali", "Iftikhar"})
	router := mux.NewRouter()

	router.HandleFunc("/", getPeople).Methods("GET")
	router.HandleFunc("/{id}", getRequiredPerson).Methods("GET")

	log.Println("Starting server at 3000")
	output := http.ListenAndServe(":3000", router)
	log.Fatal(output)
}
