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

	shortCode := res_create_short_code{"apa"}

	this.client.Set(shortCode.ShortCode, t.Url, 0)

	j, _ := json.Marshal(shortCode)
	fmt.Fprint(w, string(j))
}
