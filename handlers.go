package main

import (
	"net/http"
	"strings"

	"gopkg.in/redis.v3"
)

var client *redis.Client

type resolve_short_code_get struct {
	ShortCode string
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
