package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMissingShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/asdf", nil)

	w := httptest.NewRecorder()
	ResolveShortUrlHandler(w, req)

	if w.Code != 404 {
		t.Error("Expected handler to return 404")
	}

	fmt.Printf("%d - %s \n", w.Code, w.Body.String())
}

func TestExistnigShortUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://goshort.com/asdf", nil)

	w := httptest.NewRecorder()
	ResolveShortUrlHandler(w, req)

	if w.Code != 301 {
		t.Error("Expected handler to return redirect to full url")
	}
}

func TestAddUrlHandler(t *testing.T) {
	body := strings.NewReader("{\"url\": \"that\"}")
	req, _ := http.NewRequest("POST", "http://goshort.com/asdf", body)

	w := httptest.NewRecorder()
	AddUrlHandler(w, req)

	if w.Code != 200 {
		t.Error("Expected handler to return 200")
	}

	fmt.Printf("%d - %s \n", w.Code, w.Body.String())
}
