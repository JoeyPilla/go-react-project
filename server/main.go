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
		fmt.Fprintf(w, "Hello from server!")
	})
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
