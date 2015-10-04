package main

import (
	"fmt"
	"net/http"
)

type HomeHandlerstruct struct {
}

func (this HomeHandlerstruct) Execute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
