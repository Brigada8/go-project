package main

import (
	// Import the gorilla/mux library we just installed
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")

	r.HandleFunc("/login", login).Methods("POST")

	r.HandleFunc("/register", register).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!")
}

func register(w http.ResponseWriter, r *http.Request) {
}

func login(w http.ResponseWriter, r *http.Request) {
}
