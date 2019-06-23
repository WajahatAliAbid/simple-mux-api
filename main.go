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
	r.HandleFunc("/", createPerson).Methods("POST")
	r.HandleFunc("/{id}", updatePerson).Methods("PUT")
	r.HandleFunc("/{id}", deletePerson).Methods("DELETE")
	log.Println("Hello World")

	http.ListenAndServe(":3000", router)
}
