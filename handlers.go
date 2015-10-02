package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type create_short_code_post struct {
	Url string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func ResolveShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello")
	http.Error(w, "url not found", http.StatusNotFound)
}

func AddUrlHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t create_short_code_post
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s\r\n", t.Url)
}
