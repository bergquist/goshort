package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUrlHandler(t *testing.T) {
	fakedb := NewFakeDatabase()

	//create request
	body := strings.NewReader("{\"url\": \"http://www.grafana.com\"}")
	req, _ := http.NewRequest("POST", "http://goshort.com/create", body)

	w := httptest.NewRecorder()

	addurlHandler := AddUrlHandlerstruct{client: fakedb}

	addurlHandler.Execute(w, req)

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
	fullUrl, err := fakedb.Get(res.ShortCode)

	if fullUrl != "http://www.grafana.com" {
		t.Errorf("full url is not correct. Found %s", fullUrl)
	}
}
