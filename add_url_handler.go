package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/bergquist/goshort/models"
)

type res_create_short_code struct {
	ShortCode string
}

type AddUrlHandlerstruct struct {
	client Database
}

func (this AddUrlHandlerstruct) Execute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t models.Req_create_short_code_post
	err := decoder.Decode(&t)
	if err != nil || t.Url == "" {
		http.Error(w, "Invalid format", http.StatusBadRequest)
		return
	}

	//have this url allready been shortened
	res, checkErr := this.client.Get(t.Url)
	if checkErr == nil {
		code := res_create_short_code{string(res)}
		resjson, _ := json.Marshal(code)
		fmt.Fprint(w, string(resjson))
		return
	}

	code, _ := this.client.Incr("counter") //get uniqe id
	shortCode := res_create_short_code{string(code)}

	this.client.Set(shortCode.ShortCode, []byte(t.Url)) //set shortcode to url map
	this.client.Set(t.Url, []byte(shortCode.ShortCode)) //set url to shortcode map

	j, _ := json.Marshal(shortCode)
	fmt.Fprint(w, string(j))
}
