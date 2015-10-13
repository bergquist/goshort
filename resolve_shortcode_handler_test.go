package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMissingShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/asdf", nil)

	w := httptest.NewRecorder()
	fakedb := NewFakeDatabase()

	Router(fakedb).ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Expected handler to return 404 not %d", w.Code)
	}
}

func TestExistnigShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/apa", nil)

	w := httptest.NewRecorder()
	fakedb := NewFakeDatabase()
	fakedb.Set("apa", []byte("http://www.grafana.org/"))

	Router(fakedb).ServeHTTP(w, req)

	if w.Code != 302 {
		t.Error("Expected handler to return redirect to full url")
	}
}
