package main

import (
	"net/http"

	"gopkg.in/redis.v3"

	"github.com/gorilla/mux"
)

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func Router(db Database) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", new(HomeHandlerstruct).Execute)
	r.HandleFunc("/static/", StaticFileHandler)

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
