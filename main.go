package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/create", AddUrlHandler).
		Methods("POST")

	r.HandleFunc("/{id}", ResolveShortUrlHandler).
		Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
