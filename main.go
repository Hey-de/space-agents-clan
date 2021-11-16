package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "favicon.ico")
	})
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./images"))))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/users/", userProcessor)
	http.HandleFunc("/singup", singupHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
