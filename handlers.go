package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/redis.v3"
)

var client *redis.Client

type req_create_short_code_post struct {
	Url string
}

type res_create_short_code struct {
	ShortCode string
}

type resolve_short_code_get struct {
	ShortCode string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func ResolveShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortcode := strings.TrimLeft(r.URL.Path, "/")
	client = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})

	fullUrl, err := client.Get(shortcode).Result()

	if err != nil || fullUrl == "" {
		http.Error(w, "url not found", http.StatusNotFound)
	} else {
		http.Redirect(w, r, fullUrl, http.StatusFound)
	}
}

func AddUrlHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t req_create_short_code_post
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	shortCode := "apa"
	client = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	client.Set(shortCode, t.Url, 0)

	fmt.Fprintf(w, "{ \"shortcode\": \"%s\"} ", shortCode)
}
