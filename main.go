package main

import (
	"net/http"

	"gopkg.in/redis.v3"

	"github.com/gorilla/mux"
)

func Router(db Database) *mux.Router {
	r := mux.NewRouter()

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	r.HandleFunc("/", new(HomeHandlerstruct).Execute)

	addurlHandler := AddUrlHandlerstruct{client: db}

	r.HandleFunc("/create", addurlHandler.Execute).
		Methods("POST")

	resolveHandler := ResolveShortUrlHandlerstruct{client: db}

	r.HandleFunc("/{id}", resolveHandler.Execute).
		Methods("GET")

	return r
}

func main() {
	red := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	db := &RedisDatabase{
		client: red,
	}

	http.Handle("/", Router(db))
	panic(http.ListenAndServe(":8080", nil))
}
