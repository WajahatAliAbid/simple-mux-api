package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        int    `json:"id"`
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
	json.NewEncoder(w).Encode(people)
}

// Create a new person
func createPerson(w http.ResponseWriter, r *http.Request) {
}

// Update an existing person
func updatePerson(w http.ResponseWriter, r *http.Request) {
}

// Delete a person
func deletePerson(w http.ResponseWriter, r *http.Request) {
}

func main() {
	people = append(people, Person{1, "Wajahat", "Ali"})
	people = append(people, Person{2, "Ali", "Iftikhar"})
	router := mux.NewRouter()

	router.HandleFunc("/", getPeople).Methods("GET")
	router.HandleFunc("/{id}", getRequiredPerson).Methods("GET")
	router.HandleFunc("/", createPerson).Methods("POST")
	router.HandleFunc("/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/{id}", deletePerson).Methods("DELETE")

	log.Println("Starting server at 3000")
	output := http.ListenAndServe(":3000", router)
	log.Fatal(output)
}
