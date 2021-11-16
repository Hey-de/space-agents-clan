package main

import (
	"html/template"
	"log"
	"net/http"
)

func singupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/singup" {
		http.NotFound(w, r) // 404
		log.Printf("localhost:7777%s: Error 404", r.URL.Path)
		return
	}
	err := UpdateCounter()
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Printf("localhost:7777/: %v", err)
		return
	}
	templ, err := template.ParseFiles("ui/singup.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Printf("localhost:7777/: %v", err)
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Printf("localhost:7777/: %v", err)
		return
	}
}
