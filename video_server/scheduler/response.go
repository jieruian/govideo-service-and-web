package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, responseStr string) {
	w.WriteHeader(sc)
	io.WriteString(w, responseStr)
}

func sendNormalResponse(w http.ResponseWriter, sc int, responseStr string) {
	w.WriteHeader(sc)
	io.WriteString(w, responseStr)
}
