package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/redis.v3"
)

type req_create_short_code_post struct {
	Url string
}

type res_create_short_code struct {
	ShortCode string
}

type AddUrlHandlerstruct struct {
	client *redis.Client
}

func (this AddUrlHandlerstruct) Execute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t req_create_short_code_post
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	shortCode := "apa"

	this.client.Set(shortCode, t.Url, 0)

	fmt.Fprintf(w, "{ \"shortcode\": \"%s\"} ", shortCode)
}
