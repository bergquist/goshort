package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type req_create_short_code_post struct {
	Url string
}

type res_create_short_code struct {
	ShortCode string
}

type AddUrlHandlerstruct struct {
	client Database
}

func (this AddUrlHandlerstruct) Execute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t req_create_short_code_post
	err := decoder.Decode(&t)

	if err != nil {
		panic(err) //this seems dramatic
	}

	shortCode := res_create_short_code{"apa"}

	this.client.Set(shortCode.ShortCode, []byte(t.Url))

	j, _ := json.Marshal(shortCode)
	fmt.Fprint(w, string(j))
}
