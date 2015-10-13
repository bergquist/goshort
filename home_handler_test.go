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
	if w.Code != 200 {
		t.Errorf("Home handler should return 200. %s", body)
	}
}
