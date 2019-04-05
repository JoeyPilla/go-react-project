package main

import (
	"log"
	"net/http"
	"os"

	"./api"
)

func main() {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	//h := homepage.NewHandlers(logger)
	//todo := todo.NewHandlers(logger)
	api := api.NewHandlers(logger)
	mux := http.NewServeMux()
	api.SetupRoutes(mux)
	fs := http.FileServer(http.Dir("../client/build"))
	mux.Handle("/", fs)
	logger.Println("server starting")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
