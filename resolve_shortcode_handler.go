package main

import (
	"net/http"
	"strings"
)

type ResolveShortUrlHandlerstruct struct {
	client Database
}

func (this *ResolveShortUrlHandlerstruct) Execute(w http.ResponseWriter, r *http.Request) {
	shortcode := strings.TrimLeft(r.URL.Path, "/")
	fullUrl, err := this.client.Get(shortcode)

	if err != nil || fullUrl == "" {
		http.Error(w, "url not found", http.StatusNotFound)
	} else {
		http.Redirect(w, r, fullUrl, http.StatusFound)
	}
}
