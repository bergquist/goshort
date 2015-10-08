package main

import (
	"log"
	"net/http"

	"gopkg.in/redis.v3"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", new(HomeHandlerstruct).Execute)

	red := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	db := &RedisDatabase{
		client: red,
	}

	addurlHandler := AddUrlHandlerstruct{client: db}

	r.HandleFunc("/create", addurlHandler.Execute).
		Methods("POST")

	resolveHandler := ResolveShortUrlHandlerstruct{client: db}

	r.HandleFunc("/{id}", resolveHandler.Execute).
		Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
