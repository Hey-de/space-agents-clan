package main

import (
	"log"
	"net/http"
)

func Error500(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	log.Printf("localhost:7777%s: %v", r.URL.Path, err)
}
func ErrorBadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad request", http.StatusBadRequest)
	log.Printf("localhost:7777%s: Bad request", r.URL.Path)
}
func Error404(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r) // 404
	log.Printf("localhost:7777%s: Error 404", r.URL.Path)
}
