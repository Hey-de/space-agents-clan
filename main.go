package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

func main() {
	var addr = flag.String("addr", "localhost:7777", "Address of the website")
	var useTls = flag.Bool("tls", false, "Use TLS")
	var tlsCertFile = flag.String("tcf", "./cert/server.crt", "Certification file for TLS")
	var tlsKeyFile = flag.String("tkf", "./cert/server.key", "Key file for TLS")
	flag.Parse()
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
	if *useTls {
		log.Fatal(http.ListenAndServeTLS(*addr, *tlsCertFile, *tlsKeyFile, nil))
	} else {
		log.Fatal(http.ListenAndServe(*addr, nil))
	}
}
