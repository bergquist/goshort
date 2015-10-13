package main

import (
	"html/template"
	"net/http"
)

type HomeHandlerstruct struct {
}

func (this HomeHandlerstruct) Execute(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, nil)
}
