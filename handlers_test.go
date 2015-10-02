package main

import (
	"fmt"
	"net/http"
)

type ShortCode struct {
	fullurl   string
	shortcode string
}

type repo interface {
	Lookup(w string)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func ResolveShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func AddUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Added!")
	/*
		url := r.PostForm["url"]

		shorturl := "sss"

		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		err := client.Set(shorturl, url, 0).Err()

		if err == nil {
			fmt.Fprintf(w, "short url! %s", shorturl)
		}

		fmt.Fprintf(w, "Fail train!")
	*/
}
