package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
