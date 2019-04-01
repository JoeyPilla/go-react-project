package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../client/build"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website! %s", r.URL.Path)
	})
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
