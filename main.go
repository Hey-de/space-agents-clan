package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/icon", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "favicon.ico")
	})
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
