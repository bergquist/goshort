package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/redis.v3"
)

func TestMissingShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/asdf", nil)

	w := httptest.NewRecorder()
	resolveHandler := ResolveShortUrlHandlerstruct{
		client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0}),
	}
	resolveHandler.Execute(w, req)

	if w.Code != 404 {
		t.Error("Expected handler to return 404")
	}
}

func TestExistnigShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/apa", nil)

	w := httptest.NewRecorder()
	resolveHandler := ResolveShortUrlHandlerstruct{
		client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0}),
	}
	resolveHandler.Execute(w, req)

	if w.Code != 302 {
		t.Error("Expected handler to return redirect to full url")
	}
}
