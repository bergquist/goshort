package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/redis.v3"
)

func TestAddUrlHandler(t *testing.T) {
	//create request
	body := strings.NewReader("{\"url\": \"http://www.grafana.com\"}")
	req, _ := http.NewRequest("POST", "http://goshort.com/create", body)

	w := httptest.NewRecorder()
	AddUrlHandler(w, req)

	if w.Code != 200 {
		t.Error("Expected handler to return 200")
	}

	//decode result
	decoder := json.NewDecoder(w.Body)

	var res res_create_short_code
	err := decoder.Decode(&res)

	if err != nil {
		t.Error("Cannot decode json result")
	}

	//verify that it exists in db
	client2 := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	fullUrl, err := client2.Get(res.ShortCode).Result()

	if fullUrl != "http://www.grafana.com" {
		t.Errorf("full url is not correct. Found %s", fullUrl)
	}
}
