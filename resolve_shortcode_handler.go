package main

import (
	"html/template"
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
		w.WriteHeader(404)
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, nil)
	} else {
		http.Redirect(w, r, fullUrl, http.StatusFound)
	}
}
