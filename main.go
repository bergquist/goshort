package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HomeHandlerstruct struct {
}

func (this HomeHandlerstruct) Execute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", new(HomeHandlerstruct).Execute)

	r.HandleFunc("/create", AddUrlHandler).
		Methods("POST")

	r.HandleFunc("/{id}", ResolveShortUrlHandler).
		Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
