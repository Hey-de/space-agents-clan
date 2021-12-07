package main

import (
	"log"
	"net/http"
	"sync"
	"os"
)

var mu sync.Mutex

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "favicon.ico")
	})
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./images"))))
	http.Handle("/userdata/", http.StripPrefix("/userdata/", http.FileServer(http.Dir("./userdata"))))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/users/", userProcessor)
	http.HandleFunc("/singup", singupHandler)
	http.HandleFunc("/note", newPost)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/check", testChecker)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}