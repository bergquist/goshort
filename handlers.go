package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/redis.v3"
	"net/http"
	"strings"
)

type create_short_code_post struct {
	Url string
}

type resolve_short_code_get struct {
	ShortCode string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func ResolveShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortcode := strings.TrimLeft(r.URL.Path, "/")
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})

	fullUrl, err := client.Get(shortcode).Result()

	if err != nil || fullUrl == "" {
		http.Error(w, "url not found", http.StatusNotFound)
	} else {
		http.Redirect(w, r, fullUrl, http.StatusFound)
	}
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
