package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
	req, _ := http.NewRequest("GET", "http://goshort.com/apa", nil)

	w := httptest.NewRecorder()
	ResolveShortUrlHandler(w, req)

	if w.Code != 302 {
		t.Error("Expected handler to return redirect to full url")
	}
}
