package main

import (
	"log"
	"net/http"
	"os"

	"./api"
	"./homepage"
)

func main() {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	//define mux
	mux := http.NewServeMux()
	// grab routes
	api.AddAPIHandler(mux)
	homepage.AddHomepageHandler(mux)

	// start server
	logger.Println("server starting")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
