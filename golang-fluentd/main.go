package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// GET requests
	http.HandleFunc("/me", MeHandler)
	http.HandleFunc("/login", LoginHandler)

	log.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		log.Println("failed to start server", err)
	}
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello to me!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logging in...")
}
