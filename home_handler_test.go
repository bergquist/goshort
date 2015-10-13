package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	fakedb := NewFakeDatabase()

	req, _ := http.NewRequest("GET", "http://goshort.com/", nil)
	w := httptest.NewRecorder()

	Router(fakedb).ServeHTTP(w, req)

	body := w.Body.String()
	if body != "Hello" {
		t.Errorf("Home handler should say Hello. %s", body)
	}
}
