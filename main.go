package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var people []Person

// Get all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func getRequiredPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
}
func updatePerson(w http.ResponseWriter, r *http.Request) {
}
func deletePerson(w http.ResponseWriter, r *http.Request) {
}

func main() {
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
