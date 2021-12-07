package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-session/session"
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
	storage, err := session.Start(context.Background(), w, r)
	if err != nil {
		Error500(err, w, r)
		return
	}
	notes, err := ParseDB("data")
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
	notes["Usercount"] = db["Usercount"]
	login, ok := storage.Get("login")
	if ok {
		notes["Login"] = login
		fmt.Println(login)
	} else {
		notes["Login"] = ""
	}
	templ, err := template.ParseFiles("ui/index.html")
	if err != nil {
		Error500(err, w, r)
	}
	err = templ.Execute(w, notes)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		log.Printf("localhost:7777/: %v", err)
		return
	}
}