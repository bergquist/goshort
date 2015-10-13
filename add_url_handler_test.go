package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func parse_result(r io.Reader) (res_create_short_code, error) {
	decoder := json.NewDecoder(r)

	var res res_create_short_code
	err := decoder.Decode(&res)

	return res, err
}

func TestUrlsShouldGetSameShortCode(t *testing.T) {
	fakedb := NewFakeDatabase()

	//create request
	body := strings.NewReader("{\"url\": \"http://www.grafana.com\"}")
	req, _ := http.NewRequest("POST", "http://goshort.com/create", body)

	body2 := strings.NewReader("{\"url\": \"http://www.grafana.com\"}")
	req2, _ := http.NewRequest("POST", "http://goshort.com/create", body2)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	router := Router(fakedb)
	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	res1, _ := parse_result(w.Body)
	res2, _ := parse_result(w2.Body)

	if res1.ShortCode != res2.ShortCode {
		t.Errorf("Same url should have the same shortcode. %s != %s", res1.ShortCode, res2.ShortCode)
	}
}

func TestUrlsShouldNotGetSameShortCode(t *testing.T) {
	fakedb := NewFakeDatabase()

	//create request
	body := strings.NewReader("{\"url\": \"http://www.grafana.com\"}")
	req, _ := http.NewRequest("POST", "http://goshort.com/create", body)

	body2 := strings.NewReader("{\"url\": \"http://www.grafana.org\"}")
	req2, _ := http.NewRequest("POST", "http://goshort.com/create", body2)

	router := Router(fakedb)
	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	res1, _ := parse_result(w.Body)
	res2, _ := parse_result(w2.Body)

	if res1.ShortCode == res2.ShortCode {
		t.Errorf("different urls should not return same shortcode. %s != %s", res1.ShortCode, res2.ShortCode)
	}
}

func TestBrokenPostSouldNotPanic(t *testing.T) {
	fakedb := NewFakeDatabase()

	//create request
	body := strings.NewReader("{\"u\": \"http://www.brkgrafana.com\"}")
	req, _ := http.NewRequest("POST", "http://goshort.com/create", body)

	w := httptest.NewRecorder()
	Router(fakedb).ServeHTTP(w, req)

	if w.Code != 400 {
		t.Error("invalid requests should fail")
	}
}

func TestAddUrlHandler(t *testing.T) {
	fakedb := NewFakeDatabase()

	//create request
	body := strings.NewReader("{\"url\": \"http://www.grafana.com\"}")
	req, _ := http.NewRequest("POST", "http://goshort.com/create", body)

	w := httptest.NewRecorder()

	Router(fakedb).ServeHTTP(w, req)

	if w.Code != 200 {
		t.Error("Expected handler to return 200")
	}

	//decode result
	res, err := parse_result(w.Body)
	if err != nil {
		t.Error("Cannot decode json result")
	}

	//verify that it exists in db
	fullUrl, err := fakedb.Get(res.ShortCode)

	if fullUrl != "http://www.grafana.com" {
		t.Errorf("full url is not correct. Found %s", fullUrl)
	}
}
