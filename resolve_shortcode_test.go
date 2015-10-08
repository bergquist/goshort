package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMissingShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/asdf", nil)

	w := httptest.NewRecorder()
	resolveHandler := ResolveShortUrlHandlerstruct{
		client: NewFakeDatabase(),
	}
	resolveHandler.Execute(w, req)

	if w.Code != 404 {
		t.Error("Expected handler to return 404")
	}
}

func TestExistnigShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/apa", nil)

	w := httptest.NewRecorder()
	db := NewFakeDatabase()
	db.Set("apa", []byte("http://www.grafana.org"))
	resolveHandler := ResolveShortUrlHandlerstruct{
		client: db,
	}
	resolveHandler.Execute(w, req)

	if w.Code != 302 {
		t.Error("Expected handler to return redirect to full url")
	}
}
