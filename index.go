package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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
	db, err := ParseDB("index")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Printf("localhost:7777/: %v", err)
		return
	}
	templ, err := template.ParseFiles("ui/index.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Printf("localhost:7777/: %v", err)
		return
	}
	err = templ.Execute(w, db)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Printf("localhost:7777/: %v", err)
		return
	}
}
